// 82_通过channel实现同步
package main

import (
	"fmt"
	"time"
)

//全局变量，创建一个channel
var ch = make(chan int)

//Person1执行完后。才能到person2执行
func perosn1() {
	Printer("hello")
	ch <- 666 //给管道写数据，发送
}

func perosn2() {
	<-ch //从管道取数据，接收，如果通道没有数据他就会阻塞
	Printer("wolrd")
}

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func main() {

	// 新建两个协程去抢着打印
	go perosn1()
	go perosn2()

	for {
	}

}
