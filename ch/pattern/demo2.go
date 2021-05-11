package pattern

import (
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

// range等到channel的close动作后退出
// 如果channel不close，range就永远阻塞无法退出，又是在主程，就会死锁
func One() {
	c := make(chan int, 10)

	// 关闭channel防止range死锁
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
