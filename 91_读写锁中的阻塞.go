package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwm sync.RWMutex
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("try to lock for reading....[%d]\n", i)
			rwm.RLock()
			fmt.Printf("Locked for reading. [%d]\n", i)
			time.Sleep(time.Second * 2)
			fmt.Printf("Try to unlock for reading... [%d]\n", i)
			rwm.RUnlock()
			fmt.Printf("Unlocked for reading. [%d]\n", i)
		}(i)
	}
	time.Sleep(time.Millisecond*100)
	fmt.Println("Try to lock for writing....")
	rwm.Lock()
	fmt.Println("Locked for writing...")
}

//对于统一个读写锁来说，施加于其上的读锁定可以有多个，因此只有对互斥锁进行等量的读解锁，才能够让某一写锁定获得进行的机会，否则就会使欲进行写锁定的goroutine一直处于阻塞状态
