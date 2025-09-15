package basics

import (
	"fmt"
	"slices"
)

//slice is a reference to an array, shares storage with its array and with other slices of the same array
//本身是不存储任何数据，提供了一个进入数组元素的窗口，可以动态的增加和缩小
//make 初始化slice 初始化一个具有指定长度和容量的切片，并分配底层数组
//append  是内置函数，用于向切片中添加元素。由于切片的底层数组是固定大小的，当添加元素导致长度超过容量时，
//append 会自动创建一个新的更大的底层数组，并将原数据复制过去（即 “扩容”）。
//len 切片中当前包含的元素数量。
//cap 切片底层数组的大小（必须 ≥ 长度）。
func main() {
 
	// var numbers []int
	// var number1 =[]int{1,2,3}
	// numbers2:=[]int{9,8,7}

	// slice:=make([]int,5)
	a:=[5]int{1,2,3,4,5}
	slice1:=a[1:4] 
	fmt.Println(slice1)
	fmt.Println("slice1 cap",cap(slice1)) //容量：4（底层数组从索引 1 开始到数组 a 末尾，共 5-1=4 个位置）
	fmt.Println("slice1 length",len(slice1)) //长度：3（元素为 2,3,4）

	slice1 =append(slice1, 6,7)
	fmt.Println("slice1",slice1)
	fmt.Println("slice1 cap",cap(slice1)) //slice cap 8
	fmt.Println("slice1 length",len(slice1)) //原slice1长度为 3，容量为 4，添加 2 个元素后总长度变为 5

	// var sliceCopy2 =[5]int{1,23,4,5,6}
	sliceCopy:=make([]int,len(slice1))
	fmt.Println(sliceCopy)
	copy(sliceCopy,slice1)
	fmt.Println("sliceCopy",sliceCopy)

	// var nilSlice []int 
	// fmt.Println(nilSlice)
	for i,v :=range slice1 {
		fmt.Println(i,v)
	}

	fmt.Println("Element at index 3 of slice1", slice1[3])

	// slice1[3] = 50
	// fmt.Println("Element at index 3 of slice1", slice1[3])

	if slices.Equal(slice1,sliceCopy) {
		fmt.Println("slice1 is equal to slicecopy")
	}

	// twoD:=make([][]int,3)
	// for i:=0;i<3;i++{
	// 	innerLen :=i+1
	// 	twoD[i]=make([]int,innerLen)
	// 	for j:=0;j<innerLen;j++{
	// 		twoD[i][j]=i+j
	// 		fmt.Printf("Adding value %d in outer slice at index %d, and in inner slice index of %d\n", i+j, i, j)
	// 	}
	// }

	// fmt.Println(twoD)

	// slice[low:high]
	slice2:=slice1[2:4]
	fmt.Println(slice2)

	//子切片的容量 = 原切片容量 - 起始索引
	fmt.Println("the capacity of slice2 is",cap(slice2))//6 it is capacity of the slice that it can hold 
	fmt.Println("The len of slice2 is",len(slice2))//2


		// s := make([]int, 2, 5)  // 长度 2，容量 5
		// fmt.Println(cap(s))     // 输出：5

		// s = append(s, 1, 2, 3)  // 此时长度变为 5，容量仍为 5
		// fmt.Println(cap(s))     // 输出：5

		// s = append(s, 4)        // 长度超过容量，触发扩容，容量可能变为 10
		// fmt.Println(cap(s))     // 输出：10（具体扩容策略由 Go 内部实现决定）
}