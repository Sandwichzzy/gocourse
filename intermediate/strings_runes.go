package intermediate

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
	message2:="Hello,\tGo!"
	message3:="Hello,\rGo!" //Go!lo
	rawMessage:=`Hello\nGo`

	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println(rawMessage)

	fmt.Println("Length of rawmessage variable is", len(rawMessage))

	fmt.Println("The first character in message var is", message[0]) // ASCII

	greeting:="hello "
	name:="Alice"
	fmt.Println(greeting+name)

	str1 := "Apple"  // A has an ASCII value of 65
	str := "apple"   // a has an ASCII value of 97
	str2 := "banana" // b has an ASCII value of 98
	str3 := "app"    // a has an ASCII value of 97
	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	for i, char := range message {
		 fmt.Printf("Character at index %d is %c\n", i, char)
		// fmt.Printf("%v\n", char)
		//fmt.Printf("%x\n", char)
	}

	fmt.Println("RuneCount",utf8.RuneCountInString(greeting))
	// fmt.Println("Length",len(greeting))

	//string immutable
	greetingWithName:=greeting+name
	fmt.Println(greetingWithName )

	//Rune is alias for int32 represents(Unicode value) integer value
	var ch rune ='a'
	jch :='日'
	fmt.Println(ch)
	fmt.Println(jch)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of cstr is %T\n", cstr)

	const NIHONGO = "日本語" // Japanese text
	fmt.Println(NIHONGO)

	jhello := "こんにちは" // Japanese "Hello"
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue)
	}

	r := '😊'
	fmt.Printf("%v\n", r)
	fmt.Printf("%c\n", r)
}
