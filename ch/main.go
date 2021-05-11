package main

import "study/ch/base"

func main() {
	base.UnbufferedChannel()
	//base.BufferChannel1()
	//base.BufferChannel2()
	//base.BufferChannel3()

	//c := make(chan int, 10)
	//
	//for i := 0; i < 10; i++ {
	//	go func(index int) {
	//		c <- index + 1
	//	}(i)
	//}
	//
	//for item := range c {
	//	println(item)
	//}
	//
	//go close(c)

}
