//用牛顿法实现开方函数
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	//1、设置猜测z值的初始值
	z := 1.0 //或者 z:=float64(1)
	//2、循环用公式计算z的值 公式：z -= (z*z - x) / (2*z) 重复计算10次
	/*
		for i := 0; i < 10; i++ {
			z -= (z*z - x) / (2 * z)
			//打印z值
			fmt.Println(z)
		}
	*/
	//3、修改循环条件，使得当值停止改变（或改变非常小）的时候退出循环
	//3.1 设置一个临时变量记录上次的z值
	temp := 0.0
	for {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		//math.Abs(X) 函数用于求绝对值
		if math.Abs(z-temp) < 0.000000000000000001 {
			break
		} else {
			//3.2 将上一次的值赋值给temp以便下一次相减
			temp = z
		}
	}
	return z

}

func main() {
	fmt.Println("牛顿法", Sqrt(4))
	fmt.Println("math.Sqrt(4)", math.Sqrt(4))
}
