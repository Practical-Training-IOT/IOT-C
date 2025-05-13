package main

import (
	_ "github.com/Practical-Training-IOT/IOT-C/common/basic/core"
	user "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/user/user"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:50051")
	svr := user.NewServer(new(UserImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
