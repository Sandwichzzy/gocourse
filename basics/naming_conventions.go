package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName string
	Age int
}

func main() {
	//PascalCase
	//Eg. CalculateArea, UserInfo, NewHTTPRequest
	//Structs,interfaces,enums

	//snake_case
	//Eg. user_id, first_name,http_request

	//UPPERCASE
	//constants
	const MEXRETRIES=5
	//mixedCase
	//Eg. javaScript,htmlDocument,isValid (varible)

	var employeeID =1001


	fmt.Println(employeeID)
}
