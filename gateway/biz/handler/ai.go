package handler

import (
	"context"
	"github.com/Practical-Training-IOT/IOT-C/gateway/biz/handler/request"
	"github.com/Practical-Training-IOT/IOT-C/gateway/biz/handler/response"
	ais "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/ai"
	"github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/ai/ai"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"log"
)

func Chat(ctx context.Context, c *app.RequestContext) {
	cli, err := ai.NewClient("iot.ai", client.WithHostPorts("0.0.0.0:50058"))
	if err != nil {
		log.Fatal(err)
	}
	var id int
	if userIDInterface, exists := c.Get("userID"); exists {
		id = userIDInterface.(int)
	} else {
		responses := response.Response{
			Code: 400,
			Msg:  "未找到用户信息，请重新登录",
			Data: nil,
		}
		c.JSON(400, responses)
		return
	}
	var json request.ChatRequest
	c.BindJSON(&json)
	info, err := cli.Chat(ctx, &ais.ChatRequest{
		Id:      int32(json.Id),
		Message: json.Message,
		UserId:  int32(id),
	})
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

func History(ctx context.Context, c *app.RequestContext) {
	cli, err := ai.NewClient("iot.ai", client.WithHostPorts("0.0.0.0:50058"))
	if err != nil {
		log.Fatal(err)
	}
	var id int
	if userIDInterface, exists := c.Get("userID"); exists {
		id = userIDInterface.(int)
	} else {
		responses := response.Response{
			Code: 400,
			Msg:  "未找到用户信息，请重新登录",
			Data: nil,
		}
		c.JSON(400, responses)
		return
	}
	info, err := cli.History(ctx, &ais.HistoryRequest{UserId: int32(id)})
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

func OneHistory(ctx context.Context, c *app.RequestContext) {
	cli, err := ai.NewClient("iot.ai", client.WithHostPorts("0.0.0.0:50058"))
	if err != nil {
		log.Fatal(err)
	}
	var id int
	if userIDInterface, exists := c.Get("userID"); exists {
		id = userIDInterface.(int)
	} else {
		responses := response.Response{
			Code: 400,
			Msg:  "未找到用户信息，请重新登录",
			Data: nil,
		}
		c.JSON(400, responses)
		return
	}
	var json request.OneHistoryRequest
	c.BindJSON(&json)
	info, err := cli.OneHistory(ctx, &ais.OneHistoryRequest{
		Id:     int32(json.Id),
		UserId: int32(id),
	})
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

func Change(ctx context.Context, c *app.RequestContext) {
	cli, err := ai.NewClient("iot.ai", client.WithHostPorts("0.0.0.0:50058"))
	if err != nil {
		log.Fatal(err)
	}
	var json request.ChangeRequest
	c.BindJSON(&json)
	info, err := cli.Change(ctx, &ais.ChangeRequest{Message: json.Message})
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
