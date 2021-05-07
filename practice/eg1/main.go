package main

import (
	"fmt"
	"sync"
	"time"
)

// 任务
func job(index int) int {
	time.Sleep(time.Millisecond * 500)
	return index
}

// 常规执行
func convention(max int) time.Duration {
	start := time.Now()
	for i := 0; i < max; i++ {
		fmt.Println(job(i))
	}
	end := time.Since(start)
	return end
}

// 并发执行
func concurrent(max int) time.Duration {
	start := time.Now()

	wg := sync.WaitGroup{}
	for i := 0; i < max; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(job(i))
		}(i)
	}
	wg.Wait()

	return time.Since(start)
}

// 并发中加入无缓冲channel收集结果
func withChannel(max int) time.Duration {
	start := time.Now()

	wg := sync.WaitGroup{}
	resultChan := make(chan int)

	// channel收集结果
	for i := 0; i < max; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resultChan <- job(i)
		}(i)
	}
	// 等待收集结果的协程执行完毕，关闭channel
	go func() {
		defer close(resultChan)
		wg.Wait()
	}()
	// 从channel中遍历
	for item := range resultChan {
		fmt.Println(item)
	}

	return time.Since(start)
}

func main() {
	num := 5
	//fmt.Println("耗时：",convention(num).String())	// 普通执行
	//fmt.Println("耗时：", concurrent(num).String()) // 协程执行
	fmt.Println("耗时：", withChannel(num).String()) // 协程执行+channel
}
