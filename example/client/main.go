package main

import (
	"discovery"
	"fmt"
	"log"
	"time"
)
func init() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
}
func main() {
	for {
		time.Sleep(5 * time.Second)
		cli, e := discovery.NewClientDis([]string{"localhost:2379"})
		if e!=nil {
			panic(e)
		}
		rs, e := cli.GetService("/messageID/1/")
		fmt.Println(rs)
	}
}
