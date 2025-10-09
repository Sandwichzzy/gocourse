package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 在结构体定义中，反引号（backticks）被用作标签（tag）。
// 结构体标签是附加在结构体字段后面的字符串，
// 用于提供该字段的元数据。这些标签可以通过反射（reflect）包来获取，
// 并在运行时被一些库使用，例如JSON编码/解码、XML处理、数据库ORM等
type Person struct {
	FirstName    string  `json:"name"`
	Age          int     `json:"age,omitempty"`
	EmailAddress string  `json:"email,omitempty"`  // 字段名为 "email"，空字符串时忽略
	Address      Address `json:"address"`
}

// type Example struct {
//     Field1 string `json:"field_name"`                    // 指定字段名
//     Field2 string `json:"field_name,omitempty"`          // 零值时忽略
//     Field3 string `json:"-"`                             // 完全忽略该字段
//     Field4 string `json:"field_name,omitempty,string"`   // 多个选项用逗号分隔
// }

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}


// JSON (JavaScript Object Notation)
// `json.Marshal`-convert Go data structures into JSON (encoding)
// `json.Unmarshal`-convert JSON into Go data structures (decoding)
func main() {
	person1 := Person{FirstName: "Jane", Age: 30, EmailAddress: "jane@fakemail.com", 
										Address: Address{City: "New York", State: "NY"}}

	jsondata1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsondata1))
	// {"name":"Jane","age":30,"email":"jane@fakemail.com","address":{"city":"New York","state":"NY"}}


	jsonData1 := `{"full_name": "Jenny Doe", "emp_id": "0009", "age": 30, "address": {"city": "San Jose", "state": "CA"}}`
	var employeeFromJson Employee
	err=json.Unmarshal([]byte(jsonData1),&employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println(employeeFromJson) //{Jenny Doe 0009 30 {San Jose CA}}
	fmt.Println("Jenny's Age increased by 5 years", employeeFromJson.Age+5)
	fmt.Println("Jenny's city:", employeeFromJson.Address.City) //Jenny's city: San Jose

	listOfCityState := []Address{
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
		{City: "Las Vegas", State: "NV"},
		{City: "Modesto", State: "CA"},
		{City: "Clearwater", State: "FL"},
	}
	//[{New York NY} {San Jose CA} {Las Vegas NV} {Modesto CA} {Clearwater FL}]
	fmt.Println(listOfCityState) 

	jsonList,err:=json.Marshal(listOfCityState)
	if err != nil {
		log.Fatalln("Error Marshalling to JSON:", err)
	} 
	//JSON List: [{"city":"New York","state":"NY"},{"city":"San Jose","state":"CA"},{"city":"Las Vegas","state":"NV"},{"city":"Modesto","state":"CA"},{"city":"Clearwater","state":"FL"}]
	fmt.Println("JSON List:", string(jsonList)) 
		
	
	// 3.Handling unknown JSON structures
	jsonData2 := `{"name": "John", "age": 30, "address": {"city": "New York", "state": "NY"}}`
	var data map[string]interface{}

	err = json.Unmarshal([]byte(jsonData2), &data)
	if err != nil {
		log.Fatalln("Error Unmarshalling JSON:", err)
	}

	fmt.Println("Decoded/Unmarshalled JSON:", data) //map[address:map[city:New York state:NY] age:30 name:John]
	fmt.Println("Decoded/Unmarshalled JSON:", data["address"]) // map[city:New York state:NY]
	fmt.Println("Decoded/Unmarshalled JSON:", data["name"])  //John

}
