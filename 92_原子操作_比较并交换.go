package main

import "sync/atomic"

var value int32

func addValue(delta int32) {
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
			break
		}
	}
}

func main() {
	addValue(2)

}