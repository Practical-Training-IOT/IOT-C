package router

import (
	"context"
	"github.com/Practical-Training-IOT/IOT-C/pkg"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

func MyMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		authHeader := string(ctx.Request.Header.Peek("Authorization"))
		if authHeader == "" {
			response := map[string]interface{}{
				"code": 401,
				"msg":  "缺少授权信息",
				"data": nil,
			}
			ctx.JSON(consts.StatusOK, response)
			ctx.Abort()
			return
		}

		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			response := map[string]interface{}{
				"code": 401,
				"msg":  "Token 格式错误",
				"data": nil,
			}
			ctx.JSON(consts.StatusOK, response)
			ctx.Abort()
			return
		}
		tokenStr := authHeader[7:]

		// 解析 Token
		token, s := pkg.GetToken(tokenStr)
		if s != "" || token == nil {
			response := map[string]interface{}{
				"code": 401,
				"msg":  "登录已过期或无效 Token",
				"data": nil,
			}
			ctx.JSON(consts.StatusOK, response)
			ctx.Abort()
			return
		}

		userIDStr, ok := token["user"].(string)
		if !ok {
			response := map[string]interface{}{
				"code": 401,
				"msg":  "无效的用户信息",
				"data": nil,
			}
			ctx.JSON(consts.StatusOK, response)
			ctx.Abort()
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			response := map[string]interface{}{
				"code": 401,
				"msg":  "无效的用户 ID",
				"data": nil,
			}
			ctx.JSON(consts.StatusOK, response)
			ctx.Abort()
			return
		}

		ctx.Set("userID", userID)

		ctx.Next(c)
	}
}
