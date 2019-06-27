// 87_chan值传递和引用传递
package main

import (
	"fmt"
	"time"
)

//Count代表计数器的类型
type Counter struct {
	count int
}

//var mapChan = make(chan map[string]Counter, 1)

var mapChan = make(chan map[string]*Counter, 1)

func main() {
	synChan := make(chan struct{}, 2)
	go func() { // 用于演示接收操作
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		synChan <- struct{}{}
	}()
	go func() { // 用于演示发送操作
		//countMap := map[string]Counter{
		//	"count": Counter{},
		//}
		countMap := map[string]*Counter{
			"count": &Counter{},
		}
		for i := 0; i < 5; i++ {

			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: . [sender]\n", countMap["count"])
		}
		close(mapChan) //无论怎样都不应该在接收端关闭通道
		synChan <- struct{}{}
	}()
	<-synChan
	<-synChan
	fmt.Println("########END#######")
}
