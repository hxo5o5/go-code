//这是一个defer 关键字使用的一个例子
package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}
func untrace(s string) {
	fmt.Println("leaving:", s)
}
func a() {
	defer untrace(trace("a"))
	fmt.Println("in a")
}
func b() {
	defer untrace(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
	b()
}
