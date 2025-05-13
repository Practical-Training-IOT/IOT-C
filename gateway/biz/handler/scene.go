package handler

import (
	"context"
	"github.com/Practical-Training-IOT/IOT-C/gateway/biz/handler/request"
	"github.com/Practical-Training-IOT/IOT-C/gateway/biz/handler/response"
	scenes "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/scene"
	"github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/scene/scene"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"log"
	"strconv"
)

func SceneList(ctx context.Context, c *app.RequestContext) {
	cli, err := scene.NewClient("iot.scene", client.WithHostPorts("0.0.0.0:50054"))
	if err != nil {
		log.Fatal(err)
	}
	/*var id int
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
	}*/
	r, err := cli.List(ctx, &scenes.SceneListReq{})
	if err != nil {
		responses := response.Response{
			Code: 400,
			Msg:  "查询失败",
			Data: nil,
		}
		c.JSON(400, responses)
		return
	}
	responses := response.Response{
		Code: 200,
		Msg:  "查询成功",
		Data: r,
	}
	c.JSON(200, responses)
}

func SceneUpdateEnable(ctx context.Context, c *app.RequestContext) {
	cli, err := scene.NewClient("iot.scene", client.WithHostPorts("0.0.0.0:50054"))
	if err != nil {
		log.Fatal(err)
	}
	var json request.SceneUpdateEnable
	c.BindJSON(&json)
	r, err := cli.UpdateEnable(ctx, &scenes.SceneUpdateEnableReq{
		Id:     int32(json.ID),
		Enable: json.Enable,
	})
	if err != nil {
		responses := response.Response{
			Code: 400,
			Msg:  "修改失败",
			Data: nil,
		}
		c.JSON(400, responses)
		return
	}
	responses := response.Response{
		Code: 200,
		Msg:  "修改成功",
		Data: r,
	}
	c.JSON(200, responses)
}

func Detail(ctx context.Context, c *app.RequestContext) {
	cli, err := scene.NewClient("iot.scene", client.WithHostPorts("0.0.0.0:50054"))
	if err != nil {
		log.Fatal(err)
	}
	query := c.Query("id")
	id, _ := strconv.Atoi(query)
	detail, err := cli.Detail(ctx, &scenes.SceneDetailReq{Id: int32(id)})
	if err != nil {
		responses := response.Response{
			Code: 400,
			Msg:  "修改失败",
			Data: nil,
		}
		c.JSON(400, responses)
		return
	}
	responses := response.Response{
		Code: 200,
		Msg:  "修改成功",
		Data: detail,
	}
	c.JSON(200, responses)
}
