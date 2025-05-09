package router

import (
	"github.com/Practical-Training-IOT/IOT-C/gateway/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func SetAlarmRouter(r *server.Hertz) {
	r.GET("api/alarmList", handler.AlarmList)
	r.GET("api/alarmDetail", handler.AlarmDetail)
	r.GET("api/alarmSearch", handler.AlarmSearch)
}
