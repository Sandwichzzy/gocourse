package basics

import "fmt"


func main() {
	// var arrayName [size]elementType
	// var numbers[5]int
	// fmt.Println(numbers)
	// numbers[4]=20
	// fmt.Println(numbers)
	// numbers=[5]int{1,2,3,4,5}

	// fruits:=[4]string{"Apple","Banana","Orange","Grapes"}
	// fmt.Println("Third element:", fruits[2])

	// originalArray :=[3]int{1,2,3}
	// copiedArray:=originalArray

	// copiedArray[0]=100

	// fmt.Println("Original array:", originalArray)
	// fmt.Println("Copied array:", copiedArray)

	// for i:=0; i<len(numbers);i++{
	// 	fmt.Println("Element at index",i,":",numbers[i])
	// }

	// for index,value :=range numbers {
	// 	fmt.Printf("index: %d, value: %d\n",index,value)
	// }

	// // underscore is blank identifier, used to store unused values
	// for _, v := range numbers {
	// 	fmt.Printf(" Value: %d\n", v)
	// }

	a,_:=someFunction()
	fmt.Println(a)

		// fmt.Println("The length of numbers array is", len(numbers))

	// Comparing Arrays
	// array1 := [3]int{1, 2, 3}
	// array2 := [3]int{10, 2, 3}

	// fmt.Println("Array1 is equal to Array2:", array1 == array2)

	// var matrix [3][3]int=[3][3]int{
	// 	{1,2,3},
	// 	{4,5,6},
	// 	{7,8,9},
	// }

	// fmt.Println(matrix)
	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100

	fmt.Println("Original array:", originalArray)
	fmt.Println("Copied array:", *copiedArray)
	 
}
func someFunction() (int,int){
	return 1,2
}
