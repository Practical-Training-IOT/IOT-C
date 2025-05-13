package handler

import (
	"context"
	alarms "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/alarm"
	alarm "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/alarm/alarm"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"log"
	"strconv"
)

func AlarmList(ctx context.Context, c *app.RequestContext) {
	cli, err := alarm.NewClient("iot.alarm", client.WithHostPorts("0.0.0.0:50053"))
	if err != nil {
		log.Fatal(err)
	}
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	req := alarms.AlarmListReq{
		Page: int32(page),
		Size: int32(size),
	}
	list, err := cli.AlarmList(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, list)
}

func AlarmDetail(ctx context.Context, c *app.RequestContext) {
	cli, err := alarm.NewClient("iot.alarm", client.WithHostPorts("0.0.0.0:50053"))
	if err != nil {
		log.Fatal(err)
	}
	id, _ := strconv.Atoi(c.Query("id"))
	req := alarms.AlarmDetailReq{Id: int32(id)}
	list, err := cli.AlarmDetail(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, list)
}

func AlarmSearch(ctx context.Context, c *app.RequestContext) {
	cli, err := alarm.NewClient("iot.alarm", client.WithHostPorts("0.0.0.0:50053"))
	if err != nil {
		log.Fatal(err)
	}
	req := alarms.AlarmSearchReq{Title: c.Query("title")}
	list, err := cli.AlarmSearch(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, list)
}

func AlarmUpdate(ctx context.Context, c *app.RequestContext) {
	cli, err := alarm.NewClient("iot.alarm", client.WithHostPorts("0.0.0.0:50053"))
	if err != nil {
		log.Fatal(err)
	}
	id, _ := strconv.Atoi(c.Query("id"))
	req := alarms.AlarmUpdateReq{Id: int64(id)}
	update, err := cli.AlarmUpdate(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, update)
}
