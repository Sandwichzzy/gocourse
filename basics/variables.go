package basics

import "fmt"

// middleName := "Cane" XXXX
var middleName string ="Cane"

func main(){
	var age int
	var name string="john"
  age=1
	var name1 ="jane"

	count :=10
	name2 :="jack"
	//:= canbe infered by go compile 
	//:= can only be used in func

	//default values
	//Numeric Types:0
	//Boolean Types:false
	//String Types :" "
	//Pinters,slice,maps,functions,and structs:nil
	fmt.Println(name,name1,age,count, name2)

	// --scope
	middleName:="mayor"	
	fmt.Println(middleName)
}

func printName(){
	firstName :="Michael"
	fmt.Println(firstName)
}