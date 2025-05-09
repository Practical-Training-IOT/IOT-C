package main

import (
	_ "github.com/Practical-Training-IOT/IOT-C/common/basic/core"
	alarm "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/alarm/alarm"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:50053")

	svr := alarm.NewServer(new(AlarmImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
