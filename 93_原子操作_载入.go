package main

import "sync/atomic"

func addValue1(delta int32) {
	for {
		v := atomic.LoadInt32(&value)
		if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
			break
		}
	}
}
