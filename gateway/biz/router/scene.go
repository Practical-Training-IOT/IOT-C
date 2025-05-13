package router

import (
	"github.com/Practical-Training-IOT/IOT-C/gateway/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func SetSceneRouter(r *server.Hertz) {
	r.GET("api/scene/sceneList", MyMiddleware(), handler.SceneList)
	r.POST("api/scene/sceneUpdateEnabled", MyMiddleware(), handler.SceneUpdateEnable)
	r.GET("api/scene/sceneDetail", MyMiddleware(), handler.Detail)
}
