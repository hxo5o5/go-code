//这是帮助理解指针   * 和 & 的例子
package main

import (
	"fmt"
	"reflect"
)

type Request struct {
}
type userReq struct {
	Request
}

func returnfunc(p interface{}) {

}
func main() {
	var req = new(userReq)
	returnfunc(&req) //这里的&req 是什么
	fmt.Println("req 的地址:", req)
	fmt.Println("req 的类型", reflect.TypeOf(req))
	fmt.Println("&req 的 地址:", &req)
	fmt.Println("&req 的类型:", reflect.TypeOf(&req))
	var a *int
	fmt.Println("a的地址:", a)
	fmt.Println("a的类型", reflect.TypeOf(a))

	var b *int
	c := 1
	b = &c
	fmt.Println("b的值:", *b)
}
