package intermediate

//rune是int32的别名，它能够表示一个Unicode字符的码点。
//Unicode码点可以表示世界上几乎所有书写系统中的字符，包括ASCII字符（因为ASCII是Unicode的子集）。
//ASCII字符集只包含128个字符（0-127），而Unicode字符集则包含超过100万个字符，涵盖了多种语言和符号。
//rune则允许我们按照Unicode字符来处理字符串，每个rune对应一个Unicode字符，无论这个字符由多少个字节表示。

import (
	"fmt"
	"unicode/utf8"
)

// --Runes and characters
// -Similarities
// Representing Characters
// Storage Size
// -Differences
// Unicode Support
// Type and Size
// Encoding and Handling
// --Practical Considerations
// Internationalization
// Portability
// Efficiency
func main() {
	message:="Hello,\nGo!"
	message2:="Hello,\tGo!" //Hello,  Go!
	message3:="Hello,\rGo!" //Go!lo
	rawMessage:=`Hello\nGo`

	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println(rawMessage)

	fmt.Println("Length of rawmessage variable is", len(rawMessage)) //Length of rawmessage variable is 9

	fmt.Println("The first character in message var is", message[0]) // ASCII is 72

	greeting:="hello "
	name:="Alice"
	fmt.Println(greeting+name) //hello Alice

	str1 := "Apple"  // A has an ASCII value of 65
	str := "apple"   // a has an ASCII value of 97
	str2 := "banana" // b has an ASCII value of 98
	str3 := "app"    // a has an ASCII value of 97
	fmt.Println(str1 < str2) //true
	fmt.Println(str3 < str1) //false
	fmt.Println(str > str1) //true
	fmt.Println(str > str3) //true


// 	Character at index 0 is H
// Character at index 1 is e
// Character at index 2 is l
// Character at index 3 is l
// Character at index 4 is o
// Character at index 5 is ,
// Character at index 6 is 

// Character at index 7 is G
// Character at index 8 is o
// Character at index 9 is !
	for i, char := range message {
		 fmt.Printf("Character at index %d is %c\n", i, char) 
		// fmt.Printf("%v\n", char)
		//fmt.Printf("%x\n", char)
	}

	fmt.Println("RuneCount",utf8.RuneCountInString(greeting)) //RuneCount 6
	// fmt.Println("Length",len(greeting))

	//string immutable
	greetingWithName:=greeting+name
	fmt.Println(greetingWithName )

	//Rune is alias for int32 represents(Unicode value) integer value
	var ch rune ='a'
	jch :='日'
	fmt.Println(ch) //97
	fmt.Println(jch) //26085

	fmt.Printf("%c\n", ch) //a
	fmt.Printf("%c\n", jch) //日

	cstr := string(ch)
	fmt.Println(cstr) //a

	fmt.Printf("Type of cstr is %T\n", cstr) //Type of cstr is string

	const NIHONGO = "日本語" // Japanese text
	fmt.Println(NIHONGO) //日本語

	jhello := "こんにちは" // Japanese "Hello"
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue)
	}

	r := '😊'
	fmt.Printf("%v\n", r) //128522
	fmt.Printf("%c\n", r) //😊
}
