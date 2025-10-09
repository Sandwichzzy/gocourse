package intermediate

import (
	"fmt"
	"os"
	"strings"
)


func main() {

	user := os.Getenv("USER")
	home:=os.Getenv("HOME")

	fmt.Println("User env var:", user) //User env var: sandwichzzy
	fmt.Println("Home env var:",home) //Home env var: /home/sandwichzzy

	err:=os.Setenv("FRUIT","APPLE")
	if err!=nil {
		fmt.Println("Error setting environment variable:",err)
	}
	fmt.Println("FRUIT:",os.Getenv("FRUIT")) //FRUIT: APPLE

	// for _,e:=range os.Environ(){
	// 	kvpair:=strings.SplitN(e,"=",2)
	// 	fmt.Println(kvpair[0])
	// }

	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error unsetting environment variable:", err)
		return
	}
	fmt.Println("Unset env var done on key FRUIT")


	fmt.Println("FRUIT env var:", os.Getenv("FRUIT")) //FRUIT env var: 


	fmt.Println("---------------------------------------------------")
	str := "a=b=c=d=e"
	// n = 1 returns "a=b=c=d"
	// n = 2 returns "a" and "b=c=d"
	// n = 3 returns "a" and "b" and "c=d"
	// n = 4 returns "a" and "b" and "c" and "d"
	fmt.Println(strings.SplitN(str, "=", -1))
	fmt.Println(strings.SplitN(str, "=", 0))
	fmt.Println(strings.SplitN(str, "=", 1))
	fmt.Println(strings.SplitN(str, "=", 2))
	fmt.Println(strings.SplitN(str, "=", 3))
	fmt.Println(strings.SplitN(str, "=", 4))
	fmt.Println(strings.SplitN(str, "=", 5))
	// 	[a b c d e]
	// []
	// [a=b=c=d=e]
	// [a b=c=d=e]
	// [a b c=d=e]
	// [a b c d=e]
	// [a b c d e]
}
