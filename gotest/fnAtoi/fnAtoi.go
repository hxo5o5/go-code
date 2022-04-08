//这是将字符串转换成int 的方法
package main

import (
	"fmt"
	"strconv"
)

var str = "123456789"

func main() {
	var err error
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误", err.Error())
	}
	fmt.Println("转换成功", num)
}
