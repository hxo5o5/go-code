//元素为map类型的切片
package main

import (
	"fmt"
)

func main() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")

	//对切片中的map元素初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "hx"
	mapSlice[0]["age"] = "24"
	mapSlice[0]["sex"] = "male"
	mapSlice[1] = make(map[string]string, 10)
	mapSlice[1]["name"] = "zjw"
	mapSlice[1]["age"] = "23"
	mapSlice[1]["sex"] = "male"
	mapSlice[2] = make(map[string]string, 10)
	mapSlice[2]["name"] = "xy"
	mapSlice[2]["age"] = "27"
	mapSlice[2]["sex"] = "male"

	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}
