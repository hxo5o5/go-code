//冒泡排序
package main

import (
	"fmt"
)

//冒泡排序 普通版
func bubbleSort(s []int) {
	fmt.Println("排序前:", s)
	//1.遍历切片
	//2.比较切片内两个相邻的元素,如果s[i]>s[i+1],则交换位子，否者不变.
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	fmt.Println("排序后:", s)
}

/*冒泡排序优化版
普通冒泡排序,每次遍历一遍切片,都会将该次遍历的最大值放在最右边,那么就是每次比较后最右边的元素都不变,但是还是将这些元素遍历了一遍
优化版 就是采用倒序遍历的方式,避免重复遍历最右边的固定元素
*/
func bubbleSort2(s []int) {
	fmt.Println("排序前:", s)
	n := len(s)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	fmt.Println("排序后:", s)
}

func main() {
	arr := []int{3, 5, 7, 1, 6, 9, 84, 5, 4, 8, 6, 2}
	bubbleSort(arr)
	arr2 := []int{3, 5, 7, 1, 6, 9, 84, 5, 4, 8, 6, 2}
	bubbleSort2(arr2)
}
