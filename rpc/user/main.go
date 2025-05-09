package main

import (
	user "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/user/user"
	"log"
)

func main() {
	svr := user.NewServer(new(UserImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
