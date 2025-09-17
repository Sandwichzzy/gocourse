package main

import "fmt"

type person struct{
	name string
	age int
}

type Employee struct {
	person //Embedded struct
	empId string
	salary float64
}

func (p person) introduce(){
	fmt.Printf("HI im %s and im %d years old.\n",p.name,p.age)
}

//overriding
func (e Employee) introduce(){
	fmt.Printf("HI im %s, employee ID:%s,and im earn %.2f.\n",e.name,e.empId,e.salary)
}

func main() {
	emp:=Employee{
		person:person{name:"John",age:30},
		empId: "E001",
		salary: 50000,
	}

	fmt.Println("Name:",emp.name) //accessing the embedded struct field emp.person.name
	fmt.Println("Name:",emp.age)  //same as above
	fmt.Println("Name:",emp.empId)
	fmt.Println("Name:",emp.salary)

	//person is in employee, can directly use employee instance access introduce
	emp.introduce()
}
