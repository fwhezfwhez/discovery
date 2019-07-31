package main

import (
	"discovery"
	"os"
	"time"
)

func main() {
	ser, e := discovery.NewServiceReg([]string{"localhost:2379"}, 5)
	if e!=nil {
		panic(e)
	}
	if e := ser.PutService("/messageID/1/node/1/", "10.0.1.1:8081"); e != nil {
		panic(e)
	}
	// 假装服务崩溃
	time.Sleep(60 * time.Second)
	os.Exit(1)
}
