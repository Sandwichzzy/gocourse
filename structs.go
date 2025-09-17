package main

import "fmt"

// 1. Structs are defined using the `type` and `struct` keywords followed by curly braces `{} containing a list of fields.
// 2. Fields are defined with a name and a type.
// 3. Anonymous Structs
// 4. Anonymous Fields
// 5. Methodsfunc
// -func (value/pointer receiver) methodName(arguments, if any...) <return type, if any> {
// // Method implementation
// }
// 6. Method Declaration
// Value receiver method
// func (t Type) methodName(
// // Method implementation
// }
// Pointer receiver method
// func (t *Type) methodName( {
// // Method implementation
// }
// Comparing Structs
type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell  //匿名字段 Anonymous Fields
}

type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	city    string
	country string
}

func main() {

	p1:=Person{
		firstName:"John",
		lastName: "Wick",
		age:30,
		address: Address{
			city: "London",
			country: "UK",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "465456454",
			cell: "45456464544",
		},
	}

	p2:=Person{
		firstName: "Jane",
		age: 25,
	}
	p3 := Person{
		firstName: "Jane",
		age:       25,
	}

	// p2.address.city="New York"
	// p2.address.country="USA"

	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	fmt.Println(p1.fullName())
	fmt.Println(p1.address)
	fmt.Println(p2.address.country)
	fmt.Println(p1.address.city)
	fmt.Println(p1.cell) //匿名字段可以提升子字段到外部结构
	fmt.Println("Are p2 and p3 equal",p2==p3)

	//Annoymous struct
	user:=struct{
		username string
		email string
	}{
		username: "user1",
		email: "zetwsta@163.com",
	}

	fmt.Println(user.username)
	fmt.Println("before increment",p1.age)
	p1.incrementAgeByOne()
	fmt.Println("after increment",p1.age)
}



func (p Person) fullName() string{
	return p.firstName + " " +p.lastName
}

//为了访问原始值 并修改它 需要使用指针
func (p *Person) incrementAgeByOne() {
	p.age++
}




