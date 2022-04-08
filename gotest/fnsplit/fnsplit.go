package main

import (
	"fmt"
	"strings"
)

func main() {
	email := "abc@a.com"
	emailS := strings.Split(email, "@")
	fmt.Println(emailS) //[abc a.com]

	s := strings.Split("abc,abc", "")
	fmt.Println("empty separator ", s, len(s)) // [a b c , a b c] 7
	s = strings.Split("", "")
	fmt.Println("empty && empty ", s, len(s)) // [] 0

	s = strings.Split("", ",")
	fmt.Println("empty && not empty seperator ", s, len(s)) // [] 1   注意len是1，不是0

	s = strings.Split("abc,abc", ",") // [abc abc] 2
	fmt.Println(s, len(s))
	s = strings.Split("abc,abc", "|")
	fmt.Println("not contain separator ", s, len(s)) // not contain separator  [abc,abc] 1

	fmt.Println(len(""), len([]string{""})) // 0 1
	//str := ""
	//fmt.Println(str[0]) //panic: runtime error: index out of range [0] with length 0

	// 取某一位的值
	str := "abc"
	fmt.Println(str[0]) // 97

	str2 := "abc中午"
	fmt.Printf("%v,%v,%v,%c", str2[0], str2[3], []byte(str2), []rune(str2)[3]) // 97,228,[97 98 99 228 184 173 229 141 136],中

}
