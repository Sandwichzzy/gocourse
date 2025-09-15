package main

import (
	"fmt"
	"math"
)
func main(){
	var a,b int =10,3
	var result int
	result= a+b

	fmt.Println(result)

	result=a-b
	fmt.Println("Substraction",result)

	result=a*b
	fmt.Println("Mul",result)

	result=a/b
	fmt.Println("div",result) //3

	const p float64 =22/7.0
	fmt.Println(p) //3.142857142857143

	result=a%b
	fmt.Println("Remainder",result)

	//overflow with signed integers
	var maxInt int64=9223372036854775807 //max value that int64 can hold
	fmt.Println(maxInt)

	maxInt=maxInt+1
	fmt.Println(maxInt) //-9223372036854775808 overflow

	//overflow with unsigned integer
	var uMaxInt uint64 = 18446744073709551615 //max value for uint64
	fmt.Println(uMaxInt)

	uMaxInt=uMaxInt+1 
	fmt.Println(uMaxInt) //0

	//underflow with floating point numbers
	var smallFloat float64 =1.0e-323
	fmt.Println(smallFloat)
	smallFloat=smallFloat/math.MaxFloat64
	fmt.Println(smallFloat) //0
}