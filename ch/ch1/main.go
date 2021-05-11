package main

import (
	"fmt"
	"runtime"
)

// 无缓冲通道同步等待
func main() {
	c := make(chan struct{})

	go func(i chan struct{}) {
		sum := 0

		for i := 0; i < 1000; i++ {
			sum += i
		}
		println(sum)
		c <- struct{}{}
	}(c)

	fmt.Println("goroutine数目：", runtime.NumGoroutine())
	// 读c，通过通道进行同步等待
	<-c
}
