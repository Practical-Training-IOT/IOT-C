package main

import (
	_ "github.com/Practical-Training-IOT/IOT-C/common/basic/core"
	ai "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/ai/ai"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:50058")
	svr := ai.NewServer(new(AiImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
