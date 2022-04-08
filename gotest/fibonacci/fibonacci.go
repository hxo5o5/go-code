package main

import "fmt"

//用Fibonacci来实现生产者消费者模式
func fibonacci(write chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case write <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag=", flag)
			return
		}
	}
}
func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			sum := <-ch
			fmt.Println(sum)
		}
		quit <- true
	}()

	fibonacci(ch, quit)
}
