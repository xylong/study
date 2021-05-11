package main

func main() {
	ch := make(chan int, 1)
	go func() {
		<-ch
	}()

	println("a")
}
