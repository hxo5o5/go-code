/*练习：等价二叉查找树
1. 实现 Walk 函数。
2. 测试 Walk 函数。
函数 tree.New(k) 用于构造一个随机结构的已排序二叉查找树，它保存了值 k, 2k, 3k, ..., 10k。
创建一个新的信道 ch 并且对其进行步进：
go Walk(tree.New(1), ch)
然后从信道中读取并打印 10 个值。应当是数字 1, 2, 3, ..., 10。
3. 用 Walk 实现 Same 函数来检测 t1 和 t2 是否存储了相同的值。
4. 测试 Same 函数。
Same(tree.New(1), tree.New(1)) 应当返回 true，而 Same(tree.New(1), tree.New(2)) 应当返回 false。
*/
package main

import (
	"fmt"
	//网页上的包
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	sendValue(t, ch)
	close(ch)
}

func sendValue(t *tree.Tree, ch chan int) {
	if t != nil {
		sendValue(t.Left, ch)
		ch <- t.Value
		sendValue(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	var ch = make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
