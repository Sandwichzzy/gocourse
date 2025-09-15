package main

import (
	"fmt"
	"maps"
)


func main() {
	//1. var makeVariable =make(map[keyType]valueType)

	//2. using a Map Literal
	// mapVariable=map[keyType]valueType{
	// 	key1:value1,
	// 	key2:value2
	// }

	myMap:=make(map[string]int)
	fmt.Println(myMap) //map[]

	myMap["key1"]=9
	myMap["Code"]=90
	fmt.Println(myMap) //map[Code:90 key1:9]
	fmt.Println(myMap["key1"]) //9
	fmt.Println(myMap["key"]) //0 for int/ "" for string

	myMap["Code"]=21
	fmt.Println(myMap["Code"]) //21


	// delete(myMap,"key1")
	// fmt.Println(myMap)
  // clear(myMap)
	// fmt.Println(myMap)

	_,ok :=myMap["key1"]
	if ok {
		fmt.Println("a value exists with key1")
	}else {
		fmt.Println("no such key")
	}
	// fmt.Println(value) 
	fmt.Println("is a value associated with key1",ok) //true

	myMap2:=map[string]int{"a":1,"b":2}
	fmt.Println(myMap2)

	myMap3:=map[string]int{"a":1,"b":2}

	if maps.Equal(myMap3,myMap2){
			fmt.Println("are equal")
	}

	for k,v :=range myMap3{
		fmt.Println(k,v)
	}

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value.")
	} else {
		fmt.Println("The map is not initialized to nil value.")
	}

	val:=myMap4["key"]
	fmt.Println(val) //string ""

	// myMap4["key"] = "Value"
	// fmt.Println(myMap4)

	myMap4=make(map[string]string)
	myMap4["key"]="value"
	fmt.Println(myMap4)

	fmt.Println("myMap length is", len(myMap))

	myMap5 := make(map[string]map[string]string)

	myMap5["map1"] = myMap4
	fmt.Println(myMap5)
}
