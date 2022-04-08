package main

import (
	"flag" //flag 包实现了命令行标志的解析
	"fmt"
)

//方式一：
//使用flag.String ,Bool(),Int()...来定义
//定义命令行参数对应的变量，这三个变量都是指针类型
var (
	cliName   = flag.String("name", "hx", "input your name") //类型为 *string （指针类型）
	cliAge    = flag.Int("age", 18, "input your age")        //类型为 *int
	cliGender = flag.String("gender", "male", "input your gender")
)

//方式二：
//先定义一个 值类型 的命令行参数变量，必须在Init（）函数中对其初始化
var (
	cliFlag    int //类型为 int （值类型）
	cliFlag2   bool
	cliFlag3   string
	gopherType string
)

func Init() {
	// 初始化 flag.IntVar(),StringVar(),BoolVar()...来定义命令行参数
	flag.IntVar(&cliFlag, "int", 1, "一个值类型参数变量")
	flag.StringVar(&cliFlag3, "这里写参数名称flagname", "这里写参数值String类型", "这里写提示aaa")
	flag.BoolVar(&cliFlag2, "bool", true, "真")
	//由此可见命令行参数的定义和初始化是可以分开的

	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)

	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+"(shorthand)")

}

func main() {
	Init()
	// 读取配置文件地址
	confPath := flag.String("conf_path", "./config/cfg.json", "conf file path")
	flag.Parse()
	fmt.Printf("args=%s,num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *cliGender)
	fmt.Println("int=", *&cliFlag)
	fmt.Println("bool=", *&cliFlag2)
	fmt.Println("string=", *&cliFlag3)
	fmt.Println("gopherName=", *&gopherType)
	fmt.Println(*confPath)
}
