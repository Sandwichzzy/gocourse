package main

import (
	"flag"
	"fmt"
	"os"
)

//在Go语言中，os.Args是一个字符串切片（[]string），
//用于存储命令行参数。其中，第一个元素os.Args[0]是程序的名称（或路径），后续元素是程序运行时传入的参数。

func main() {

	fmt.Println("Command:",os.Args[0])

	for i,arg :=range os.Args{
		fmt.Println("Argument",i,":",arg)
	}

	// Define flags
	var name string
	var age int
	var male bool


	//flag 包是 Go 语言标准库中用于解析命令行参数的包，它提供了简单易用的方式来定义和解析命令行标志。
	//flag 包支持以下类型的命令行参数：

	// 字符串 (string)

	// 整数 (int, int64)

	// 无符号整数 (uint, uint64)

	// 浮点数 (float64)

	// 布尔值 (bool)

	// 时间间隔 (time.Duration)
	//func flag.StringVar(p *string, name string, value string, usage string)
	// p : 指向字符串变量的指针，解析后的值将存储在这里

	// name : 命令行标志的名称

	// value : 默认值

	// usage : 帮助信息
	flag.StringVar(&name, "name","John", "Name of the user")
	flag.IntVar(&age, "age", 18, "Age of the user")
	flag.BoolVar(&male, "male", true, "Gender of the user")

	flag.Parse()
	//go run command_line_args.go -name James Doe -age 50 
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Male:", male)
	// Name: James
	// Age: 18
	// Male: true
}
