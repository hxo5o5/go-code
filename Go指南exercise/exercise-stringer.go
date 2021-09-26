//练习：Stringer
//通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。
//例如，IPAddr{1, 2, 3, 4} 应当打印为 "1.2.3.4"。
package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法

func (i IPAddr) String() string {
	str := "" //初值化""
	//str="",所以不执行if
	//127为byte(uint8)转成int(uint8和int是完全不一样的类型)，
	//再转成string然后赋值给了str，然后str就不为空开始加点
	for _, v := range i {
		//如果str不为空就执行
		if str != "" {
			str += "."
		}
		fmt.Printf(strconv.Itoa(int(v)) + " ") //strconv.Itoa(int(v))是将int转成Sting
		str += strconv.Itoa(int(v))
	}
	return str
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
