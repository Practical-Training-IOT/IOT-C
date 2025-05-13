package router

import (
	"github.com/Practical-Training-IOT/IOT-C/gateway/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func SetAIRouter(r *server.Hertz) {
	r.GET("/api/ai/history", MyMiddleware(), handler.History)
	r.POST("/api/ai/change", MyMiddleware(), handler.Change)
	r.POST("/api/ai/oneHistory", MyMiddleware(), handler.OneHistory)
	r.POST("/api/ai/chat", MyMiddleware(), handler.Chat)
}
