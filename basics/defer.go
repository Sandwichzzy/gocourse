package basics

import "fmt"

// 1. defer
// 作用：延迟执行，用于确保函数调用在所在函数返回之前执行。

// go
// func main() {
//     defer fmt.Println("这行最后执行")
//     fmt.Println("这行先执行")
//     // 输出：
//     // 这行先执行
//     // 这行最后执行
// }
// 特点：

// 多个 defer 按后进先出顺序执行

// 常用于资源清理（文件关闭、锁释放等）

// defer 语句中的参数会立即求值

func main() {

	process(10)
}
// Use cases
// 1.resource cleanup
// 2.unlocking mutexes
// 3.logging and tracing
func process(i int){
	defer fmt.Println("Deffered i value:", i) //Deffered i value: 10 参数取决于延迟语句
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed")
	fmt.Println("Normal execution statement")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i:", i) //Value of i: 11
}