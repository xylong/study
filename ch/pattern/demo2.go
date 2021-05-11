package pattern

import (
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func One() {
	c := make(chan int, 10)

	go func() {
		defer close(c)
		wg.Wait()
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			c <- index + 1
			time.Sleep(time.Millisecond * 500)
		}(i)
	}

	// 等待channel直到关闭,即是阻塞
	for item := range c {
		println(item)
	}
}
