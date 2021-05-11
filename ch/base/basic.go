package base

import "sync"

// 无缓冲channel
func UnbufferedChannel() {
	ch1, ch2 := make(chan string), make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		println("step2")
		ch1 <- "foo"
		println("step3")
		println(<-ch2)
		println("step6")
	}()

	println("step1")
	println(<-ch1) // 阻塞
	println("step4")
	ch2 <- "bar"
	println("step5")
	wg.Wait()
}

// channel本身结构是一个先进先出的队列
func BufferChannel1() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	println(<-ch, <-ch)
}

// channel关闭了，就不能往信道传值了，否则会报错
func BufferChannel2() {
	ch := make(chan string, 1)
	ch <- "a"
	close(ch)
	ch <- "b"
}

// 如果取完了信道存储的信息再去取信息，会死锁
func BufferChannel3() {
	ch := make(chan string, 2)
	ch <- "apple"
	ch <- "banana"

	for item := range ch {
		println(item)
	}
}
