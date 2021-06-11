package main

import (
	"fmt"
	"time"
)

func main() {
	//base.UnbufferedChannel()
	//base.BufferChannel1()
	//base.BufferChannel2()
	//base.BufferChannel3()

	//pattern.UnbufferedChannelWait()

	//pattern.One()

	c := make(chan int, 10)

	for i := 0; i < 10; i++ {
		go func(index int) {
			c <- index
		}(i)
		time.Sleep(time.Microsecond * 500)
	}
	go close(c)

	for item := range c {
		fmt.Println(item)
	}
}
