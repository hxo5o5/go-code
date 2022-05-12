//基本类型切片排序
package main

import (
	"fmt"
	"sort"
)

func main() {
	//[]int 排序
	slInt := []int{5, 5, 4, 9, 2, 3, 7, 1} //unsorted
	sort.Ints(slInt)
	fmt.Println(slInt) //输出:[1 2 3 4 5 5 7 9]

	//[]float64 排序
	slF64 := []float64{5.2, -1.3, 07, -3.8, 2.6} //unsorted
	sort.Float64s(slF64)
	fmt.Println(slF64) //输出:[-3.8 -1.3 2.6 5.2 7]

	//[]string 字典排序
	slStr := []string{"go", "Go", "GO", "Java", "C", "C++"} //unsorted
	sort.Strings(slStr)
	fmt.Println(slStr) //输出:[C C++ GO Go Java go]
}
