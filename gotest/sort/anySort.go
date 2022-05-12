//排序任意数据结构
package main

import (
	"fmt"
	"sort"
)

type Student struct {
	string
	int
}
type SortByAgeName []Student

func (s SortByAgeName) Len() int {
	return len(s)
}

func (s SortByAgeName) Less(i, j int) bool {
	if s[i].int < s[j].int {
		return true
	}
	if s[i].int > s[j].int {
		return false
	}
	return s[i].string < s[j].string
}

func (s SortByAgeName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	s := []Student{
		{"hx", 24},
		{"zzy", 25},
		{"yy", 23},
		{"zjw", 23},
		{"hjb", 25},
		{"lmj", 23},
		{"yxc", 24},
		{"hx", 24},
	}
	fmt.Println("排序前:", s) //排序前: [{hx 24} {zzy 25} {yy 23} {zjw 23} {hjb 25} {lmj 23} {yxc 24} {hx 24}]
	sort.Sort(SortByAgeName(s))
	fmt.Println("排序后:", s) //排序后: [{lmj 23} {yy 23} {zjw 23} {hx 24} {hx 24} {yxc 24} {hjb 25} {zzy 25}]
}
