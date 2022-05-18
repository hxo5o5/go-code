//递归实现Fibonacci
package main

import "fmt"

func fibonacci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\n", fibonacci(i))
	}
}
