package main

import (
	"time"
)

func main() {
	ch := make(chan int, 1)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- 1
		}
	}()

	go func() {
		for {
			println(<-ch)           //5次 后 报死锁
			time.Sleep(time.Second) //为了效果 我故意 休眠1秒
		}
	}()
	select {}
}
