package main

import (
	_ "github.com/Practical-Training-IOT/IOT-C/common/basic/core"
	scene "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/scene/scene"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:50054")
	svr := scene.NewServer(new(SceneImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
