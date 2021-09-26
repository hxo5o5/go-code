//练习：斐波那契闭包
//实现一个 fibonacci 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 (0, 1, 1, 2, 3, 5, ...)。
package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	//该数列从0，1开始，两两相加成为第三个数，所以初始值设置为-1和1相加刚好从0开始
	var n, sum int = -1, 1
	//注意：要求返回int的函数。func fibonacci() 后面的 func() int 是返回参数
	return func() int {
		//双赋值，很好用，go语言特性之一，不需要声明一个新的中间变量
		n, sum = sum, (n + sum)
		return sum
	}
}

func main() {
	f := fibonacci()
	//循环10次输出数列的前10个
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
