package main

import "fmt"

// 1. A pointer is a variable that stores the memory address of another variable
// 2. Use Cases
// （1）Modify the value of a variable indirectly
// （2）Pass large data structures efficiently between functions
// （3）Manage memory directly for performance reasons
// 3. Pointer Declaration and Initialization
// 4. Pointer Operations: Limited to referencing (&`) and dereferencing (`*~)
// 5. Nil Pointers
// 6. Go does not support pointer arithmetic like C or C++
// 7. Passing Pointers to Functions
// 8. Pointers to Structs
// 9. Use pointers when a function needs to modify an argument's value
// 10. unsafe.Pointer(&x) converts the address of x to an unsafe.Pointer

//go 中指针操作只有referencing 和dereferencing 没有指针运算
func main() {

  var ptr *int 
	var a int=10
	ptr=&a // referencing ptr now points to a's memory address

	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr) // dereferencing a pointer
	// if ptr == nil {
	// 	fmt.Println("Pointer is nil")
	// }

	modifyValue(ptr)
	fmt.Println(a)
}

//使用指针变量，变量的实际内存地址会传递给函数，会修改原始变量的值
func modifyValue(ptr *int){
	*ptr ++
}
