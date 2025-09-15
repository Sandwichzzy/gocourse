package main

import "fmt"

func main(){
	message:="Hello World"

	for _,v:= range message{
		fmt.Printf("Rune:%c\n",v)
	}
}
