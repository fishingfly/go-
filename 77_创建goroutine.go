// 77_创建goroutine
package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a new task")
		time.Sleep(time.Second) //延时一秒
	}
}

func main() {
	go newTask() //创建一个协程
	for {
		fmt.Println("this is main roroutine")
		time.Sleep(time.Second) //延时1秒
	}
}
