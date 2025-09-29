package intermediate

import (
	"fmt"
	"strconv"
)


func main() {

	//func strconv.Atoi(s string) (int, error)
	numStr:="12345"
	num,err:=strconv.Atoi(numStr)
	if err!=nil {
		fmt.Println("Error parsing the value:",err)
	}
	fmt.Println("Parsed Integer:",num)
	fmt.Println("Parsed Integer:",num+1)

	// Go 语言中非常常用的字符串到整数转换函数。
	// s string - 要解析的字符串
	//base int - 进制基数((0, 2 to 36)) 10-十进制
	//bitSize int - 位宽限制 int32
	//func strconv.ParseInt(s string, base int, bitSize int) (i int64, err error)
	numiStr,err:=strconv.ParseInt(numStr,10,32)
		if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed Integer:", numiStr) //12345


	//func strconv.ParseFloat(s string, bitSize int) (float64, error)
	floatStr:="3.14"
	floatVal,err:=strconv.ParseFloat(floatStr,64)
	if err!=nil{
		fmt.Println("Error parsing value",err)
	}
	fmt.Printf("Parsed float:%.2f\n",floatVal)

	//func strconv.ParseInt(s string, base int, bitSize int) (i int64, err error)
	binaryStr := "1010" // 0 + 2 + 0 + 8 = 10
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error parsing binary value:", err)
		return
	}
	fmt.Println("Parsed binary to decimal:", decimal)//10

	hexStr := "FF"
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error parsing binary value:", err)
		return
	}
	fmt.Println("Parsed hex to decimal:", hex)//255

	//-------------------------------------------------------------------
	//错误示例1：无效字符 
	invalidNum := "456abc"
	invalidParse, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Println("Error parsing value:", err)
	}
	fmt.Println("Parsed invalid number:", invalidParse)

	  // 错误示例2：空字符串
    _, err2 := strconv.ParseInt("", 10, 64)
    fmt.Println("错误2:", err2) // strconv.ParseInt: parsing "": invalid syntax

    // 错误示例3：超出范围
    _, err3 := strconv.ParseInt("9999999999999999999", 10, 32)
    fmt.Println("错误3:", err3) // strconv.ParseInt: parsing "9999999999999999999": value out of range

    // 错误示例4：无效的进制
    _, err4 := strconv.ParseInt("123", 1, 64) // 进制必须是 0 或 2-36
    fmt.Println("错误4:", err4) // strconv.ParseInt: parsing "123": invalid base 1

    // 错误示例5：只有符号没有数字
    _, err5 := strconv.ParseInt("+", 10, 64)
    fmt.Println("错误5:", err5) // strconv.ParseInt: parsing "+": invalid syntax
}