//实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。
//函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。
//你会发现 strings.Fields 很有帮助。
package main

import (
	"strings"
	//也是网页上的包
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int) //创建映射
	c := strings.Fields(s)    //返回单词的个数 strings.Fields(x) 以连续的空白字符为分隔符，将x分成多个子串，结果不包含空白符本身
	for _, v := range c {     //每出现相同的单词（字符串）
		m[v] += 1 //出现次数就 + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}