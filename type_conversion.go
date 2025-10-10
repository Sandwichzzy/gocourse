package main

import "fmt"


func main() {

	var a int=32
	b:=int32(a)
	c:=float64(b)

	// d:=bool(true)
	e:=3.14
	f:=int(e) //3
	fmt.Println(f,c)

	g:="Hello @ こんにちは 🧑 привет"
	var h []byte =[]byte(g) 
	//[72 101 108 108 111 32 64 32 227 129 147 227 130 147 227 129 171 227 129 161 227 129 175 32 240 159 167 145 32 208 191 209 128 208 184 208 178 208 181 209 130]
	fmt.Println(h)

	i := []byte{255, 120, 72}
	j := string(i)
	fmt.Println(j)
}