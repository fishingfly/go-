// 88_单向通道的应用场景
package main

import (
	"fmt"
	"time"
)

var strChan1 = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go receive(strChan1, syncChan1, syncChan2)
	go send(strChan1, syncChan1, syncChan2)
	<-syncChan2
	<-syncChan2
}

func receive(strChan1 <-chan string, syncChan1 <-chan struct{}, syncChan2 chan<- struct{}) {
	<-syncChan1
	fmt.Println("Received a sync signal and wait a seconde ... [receiver]")
	time.Sleep(time.Second)
	for {
		if elem, ok := <-strChan1; ok {
			fmt.Println("Received: ", elem, "[receive]")
		} else {
			break
		}
	}
	fmt.Println("Stopped. [receiver]")
	syncChan2 <- struct{}{}
}

func send(strChan chan<- string, syncChan1 chan<- struct{},syncChan2 chan<- struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan<- elem
		fmt.Println("sent: ", elem, "[sender]")
		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("sent a sync signal.[sender]")
		}
	}
	fmt.Println("wait 2 seconds ... [sender]")
	time.Sleep(time.Second*2)
	close(strChan1)
	syncChan2 <- struct{}{}

}

