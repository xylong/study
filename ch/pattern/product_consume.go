// 生产消费模式
package pattern

import (
	"fmt"
	"time"
)

// producer1 生产者1
func producer1(c chan<- int) {
	defer close(c)

	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(time.Millisecond * 100)
	}
}

// consumer1 消费者1
func consumer1(c <-chan int) {
	for item := range c {
		fmt.Println(item)
	}
}

func run1() {
	c := make(chan int)
	go producer1(c)
	consumer1(c)
}

// producer2  生产者2
func producer2(c chan<- int) {
	defer close(c)

	for i := 0; i < 10; i++ {
		c <- i * 2
		time.Sleep(time.Millisecond * 100)
	}
}

func consumer2(c <-chan int, done chan<- struct{}) {
	for item := range c {
		fmt.Println(item)
	}
	done <- struct{}{}
}

func run2() {
	c := make(chan int)
	done := make(chan struct{})
	go producer2(c)
	go consumer2(c, done)
	<-done
}

func producer3(c chan<- int) {
	defer close(c)

	for i := 0; i < 10; i++ {
		c <- i * 3
		time.Sleep(time.Millisecond * 100)
	}
}

func consumer3(c <-chan int) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer func() {
			done <- struct{}{}
		}()

		for item := range c {
			fmt.Println(item)
		}
	}()

	return done
}

func run3() {
	c := make(chan int)
	go producer3(c)
	<-consumer3(c)
}
