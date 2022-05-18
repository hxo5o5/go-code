//快速排序
package main

import (
	"fmt"
)

func quickSort(arr []int, sortWay int) []int {
	var data = make([]int, len(arr))
	copy(data, arr)
	fmt.Println("排序前:", data)
	_quickSort(data, 0, len(data)-1, sortWay)
	fmt.Println("排序后:", data)
	switch sortWay {
	case 1:
		fmt.Println("排序方法:单路快排")
	case 2:
		fmt.Println("排序方法:双路快排")
	case 3:
		fmt.Println("排序方法:三路排序")
	}
	return data
}

func _quickSort(arr []int, left int, right int, sortWay int) []int {
	var (
		partitionIndex int
		lo             int //相等区左边界
		gt             int //相等区右边界
	)
	if left < right {
		switch sortWay {
		case 1:
			partitionIndex = partition1Way(arr, left, right)
			_quickSort(arr, left, partitionIndex-1, sortWay)
			_quickSort(arr, partitionIndex+1, right, sortWay)
		case 2:
			partitionIndex = partition2Way(arr, left, right)
			_quickSort(arr, left, partitionIndex-1, sortWay)
			_quickSort(arr, partitionIndex+1, right, sortWay)
		case 3:
			lo, gt = partition3Way(arr, left, right)
			_quickSort(arr, left, lo-1, sortWay)
			_quickSort(arr, gt, right, sortWay)
		default:
			println("请选择正确的排序方式")
		}
	}
	return arr
}

//单路快排:从左向右遍历
func partition1Way(arr []int, left int, right int) int {
	//先分区,最后把基准换到边界上
	privot := left
	index := privot + 1
	for i := index; i <= right; i++ {
		//当前值小于基准值就交换,大于的不用管
		if arr[i] < arr[privot] {
			arr[index], arr[i] = arr[i], arr[index]
			index++
		}
	}
	arr[privot], arr[index-1] = arr[index-1], arr[privot]
	return index - 1
}

//双路快排:双指针从首尾向中间移动
func partition2Way(arr []int, low int, high int) int {
	//基准
	tmp := arr[low]
	for low < high {
		//当队尾的元素大于等于基准数据时,向前挪动high指针
		for low < high && arr[high] >= tmp {
			high--
		}
		//如果队尾元素小于tmp了,需要将其赋值给low
		arr[low] = arr[high]
		//当队首元素小于等于tmp时,向前挪动low指针
		for low < high && arr[low] <= tmp {
			low++
		}
		//当队首元素大于tmp时.需要将其赋值给high
		arr[high] = arr[low]
	}
	//跳出循环时low和high相等,此时的low或high就是tmp的正确索引位置
	//由原理部分可以很清楚的知道low位置的值并不是tmp,所以需要将tmp赋值给arr[low]
	arr[low] = tmp
	return low
}

//三路排序:分成小于区、等于区、大于区，不对等于区进行递归操作
func partition3Way(arr []int, left, right int) (int, int) {
	key := arr[left]
	lo, gt, cur := left, right+1, left+1
	for cur < gt {
		//小于key,移到前面
		if arr[cur] < key {
			arr[cur], arr[lo+1] = arr[lo+1], arr[cur] //lo+1 保证最后arr[lo]小于key
			lo++                                      //左边界右移
			cur++                                     //能够确实换完之后该位置小于key
		} else if arr[cur] > key {
			arr[cur], arr[gt-1] = arr[gt-1], arr[cur]
			gt-- //从后面换到前面,不知道是否比key的大,还要再比一下,所以cur不移动
		} else {
			cur++
		}
	}
	arr[left], arr[lo] = arr[lo], arr[left] //最后移动基准,arr[lo]一定比key小
	return lo, gt
}

func main() {
	arr := []int{3, 5, 7, 1, 6, 9, 84, 5, 4, 8, 6, 2}
	quickSort(arr, 1)
	quickSort(arr, 2)
	quickSort(arr, 3)
}
