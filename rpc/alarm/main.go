package main

import (
	_ "github.com/Practical-Training-IOT/IOT-C/common/basic/core"
	alarm "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/alarm/alarm"
	"log"
)

func main() {
	svr := alarm.NewServer(new(AlarmImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
