package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	//1.初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)

	//2.赋值
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	//3.取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	//4.对切片进行排序
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
