//值为切片类型的map
package main

import (
	"fmt"
)

func main() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "姓名"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "hx", "zjw")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
