package handler

import (
	"context"
	"fmt"
	"github.com/Practical-Training-IOT/IOT-C/gateway/biz/handler/request"
	"github.com/Practical-Training-IOT/IOT-C/gateway/biz/handler/response"
	users "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/user"
	"github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/user/user"
	"github.com/Practical-Training-IOT/IOT-C/pkg"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"log"
	"strconv"
)

func Register(ctx context.Context, c *app.RequestContext) {
	cli, err := user.NewClient("iot.user", client.WithHostPorts("0.0.0.0:50051"))
	if err != nil {
		log.Fatal(err)
	}
	var register request.Register
	err = c.BindJSON(&register)
	if err != nil {
		responses := response.Response{
			Code: 400,
			Msg:  "注册失败",
			Data: nil,
		}
		c.JSON(400, responses)
		return
	}
	r, err := cli.Register(ctx, &users.RegisterReq{
		UserName: register.Username,
		Password: register.Password,
	})
	if err != nil {
		responses := response.Response{
			Code: 400,
			Msg:  "注册失败",
			Data: nil,
		}
		c.JSON(400, responses)
		return
	}
	responses := response.Response{
		Code: 200,
		Msg:  "注册成功",
		Data: r.Id,
	}
	c.JSON(200, responses)
}

func Login(ctx context.Context, c *app.RequestContext) {
	cli, err := user.NewClient("iot.user", client.WithHostPorts("0.0.0.0:50051"))
	if err != nil {
		log.Fatal(err)
	}
	var register request.Register
	err = c.BindJSON(&register)
	if err != nil {
		responses := response.Response{
			Code: 400,
			Msg:  "注册失败",
			Data: nil,
		}
		c.JSON(400, responses)
		return
	}
	r, err := cli.Login(ctx, &users.LoginReq{
		UserName: register.Username,
		Password: register.Password,
	})
	if err != nil {
		responses := response.Response{
			Code: 400,
			Msg:  "注册失败",
			Data: nil,
		}
		c.JSON(400, responses)
		return
	}
	fmt.Println(r.Id)
	id := strconv.Itoa(int(r.Id))
	handler, err := pkg.TokenHandler(id)
	if err != nil {
		responses := response.Response{
			Code: 400,
			Msg:  "注册失败",
			Data: nil,
		}
		c.JSON(400, responses)
		return
	}
	fmt.Println(handler)
	responses := response.Response{
		Code: 200,
		Msg:  "登录成功",
		Data: handler,
	}
	c.JSON(200, responses)
}

func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	cli, err := user.NewClient("iot.user", client.WithHostPorts("0.0.0.0:50051"))
	if err != nil {
		log.Fatal(err)
	}
	get := c.Request.Header.Get("Authorization")
	tokens := get[7:]
	fmt.Println(tokens)
	token, _ := pkg.GetToken(tokens)
	if token == nil {
		responses := response.Response{
			Code: 401,
			Msg:  "登录失效",
			Data: nil,
		}
		c.JSON(200, responses)
	}
	s := token["user"].(string)
	fmt.Println(s)
	fmt.Println(token)
	id, _ := strconv.Atoi(s)
	info, err := cli.UserInfo(ctx, &users.UserInfoReq{Id: int32(id)})
	if err != nil {
		responses := response.Response{
			Code: 400,
			Msg:  "获取失败",
			Data: nil,
		}
		c.JSON(400, responses)
		return
	}
	responses := response.Response{
		Code: 200,
		Msg:  "获取成功",
		Data: info,
	}
	c.JSON(200, responses)
}

func UpdateUserInfo(ctx context.Context, c *app.RequestContext) {
	cli, err := user.NewClient("iot.user", client.WithHostPorts("0.0.0.0:50051"))
	if err != nil {
		log.Fatal(err)
	}
	get := c.Request.Header.Get("Authorization")
	tokens := get[7:]
	fmt.Println(tokens)
	token, _ := pkg.GetToken(tokens)
	if token == nil {
		responses := response.Response{
			Code: 401,
			Msg:  "登录失效",
			Data: nil,
		}
		c.JSON(200, responses)
	}
	s := token["user"].(string)
	fmt.Println(s)
	fmt.Println(token)
	id, _ := strconv.Atoi(s)
	var userInfo request.UserInfoUpdate
	c.BindJSON(&userInfo)
	info, err := cli.UserInfoUpload(ctx, &users.UserInfoUploadReq{
		Department: userInfo.Department,
		Email:      userInfo.Email,
		Phone:      int64(userInfo.Phone),
		Username:   userInfo.Username,
		Id:         int32(id),
	})
	if err != nil {
		responses := response.Response{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(400, responses)
		return
	}
	responses := response.Response{
		Code: 200,
		Msg:  "修改成功",
		Data: info,
	}
	c.JSON(200, responses)
}
