// 86_chan传引用类型
package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)

func main() {
	synChan := make(chan struct{}, 2)
	go func() { // 用于演示接收操作
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"]++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		synChan <- struct{}{}
	}()

	go func() { // 用于演示发送操作
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		synChan <- struct{}{}
	}()
	<-synChan
	<-synChan
	fmt.Println("#####END####")
}
