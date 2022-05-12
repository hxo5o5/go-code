//自定义比较器
package main

import (
	"fmt"
	"sort"
)

func main() {
	slstudent := []struct {
		string
		int
	}{
		{"hx", 24},
		{"zzy", 25},
		{"yy", 23},
		{"zjw", 23},
		{"hjb", 25},
		{"lmj", 23},
		{"yxc", 24},
		{"hx", 24},
	}

	//用Slice方法,排序age
	sort.Slice(slstudent, func(i, j int) bool { return slstudent[i].int < slstudent[j].int })

	fmt.Println(slstudent) //输出:[{yy 23} {zjw 23} {lmj 23} {hx 24} {hx 24} {yxc 24} {hjb 25} {zzy 25}]

	slstudent2 := []struct {
		string
		int
	}{
		{"hx", 24},
		{"zzy", 25},
		{"yy", 23},
		{"zjw", 23},
		{"hjb", 25},
		{"lmj", 23},
		{"yxc", 24},
		{"hx", 24},
	}

	//用SliceStable方法,排序age 升序
	sort.SliceStable(slstudent2, func(i, j int) bool { return slstudent2[i].int < slstudent2[j].int })
	fmt.Println(slstudent2) //输出:[{yy 23} {zjw 23} {lmj 23} {hx 24} {yxc 24} {hx 24} {zzy 25} {hjb 25}]

	//用SliceStable方法,排序age 降序
	sort.SliceStable(slstudent2, func(i, j int) bool { return slstudent2[i].int > slstudent2[j].int })
	fmt.Println(slstudent2) //输出:[{zzy 25} {hjb 25} {hx 24} {yxc 24} {hx 24} {yy 23} {zjw 23} {lmj 23}]

	slstudent3 := []struct {
		string
		int
	}{
		{"hx", 24},
		{"zzy", 25},
		{"yy", 23},
		{"zjw", 23},
		{"hjb", 25},
		{"lmj", 23},
		{"yxc", 24},
		{"hx", 24},
	}

	//排序完age,之后再按姓名排序
	sort.SliceStable(slstudent3, func(i, j int) bool {
		if slstudent3[i].int < slstudent3[j].int {
			return true
		}
		if slstudent3[i].int > slstudent3[j].int {
			return false
		}
		//fmt.Println("看看参数比较的啥", slstudent3[i].string, slstudent3[j].string)
		return slstudent3[i].string < slstudent3[j].string
	})

	fmt.Println(slstudent3) //输出:[{lmj 23} {yy 23} {zjw 23} {hx 24} {hx 24} {yxc 24} {hjb 25} {zzy 25}]
}
