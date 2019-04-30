package main

import (
	"time"

	"github.com/yakaa/log4g"
)


func main() {

	log4g.Init(log4g.Config{
		Path: "logs",
	})
	for {
		log4g.InfoFormat("GOOD === >当前时间为%s", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(1 * time.Second)

	}

}
