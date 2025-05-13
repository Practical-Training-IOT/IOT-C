package router

import (
	"github.com/Practical-Training-IOT/IOT-C/gateway/biz/handler"
	"github.com/Practical-Training-IOT/IOT-C/pkg"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func SetUserRouter(r *server.Hertz) {
	r.POST("api/register", handler.Register)
	r.POST("api/login", handler.Login)
	r.GET("api/user/getUserInfo", MyMiddleware(), handler.GetUserInfo)
	r.POST("api/user/uploadAvatar", MyMiddleware(), pkg.UploadAvatarHandler)
	r.POST("api/user/updateUserInfo", MyMiddleware(), handler.UpdateUserInfo)
}
