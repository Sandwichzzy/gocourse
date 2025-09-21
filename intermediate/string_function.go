package intermediate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)


func main() {
	str:= "Hello Go!"

	fmt.Println(len(str))

	str1:="Hello"
	str2:="World"

	result:=str1 +" "+str2
	fmt.Println(result) // Hello World

	fmt.Println(str[0]) //72 ASCII
	var r rune ='H' 
	fmt.Println(r) //72
	fmt.Println(str[1:5]) //ello [)

	// s := "Hello, 世界"
  // for _, r := range s {
  //     fmt.Printf("字符: %c, Unicode码点: %U, unicode:%v\n", r, r,r)
  // }

	//string conversion
	num:=18
	str3:=strconv.Itoa(num) //int to string
	fmt.Println(len(str3)) //2
	
	// strings splitting
	fruits := "apple, orange, banana"
	fruits1 := "apple-orange-banana"
	parts := strings.Split(fruits, ",")
	parts1 := strings.Split(fruits1, "-")
	fmt.Println(parts) //[apple orange banana]
	fmt.Println(parts1) // [apple orange banana]

	//strings join
	countries:=[]string{"Germany","France","Italy"}
	joined:=strings.Join(countries,",")
	fmt.Println(joined) //Germany,France,Italy

	//strings.Contains
	fmt.Println(strings.Contains(str,"Go")) //true

	//strings.Replace(s string ，old string，new string, n int 替换前几个)
	replaced:=strings.Replace(str,"Go","World",1) //Hello World!
	fmt.Println(replaced)

	//strings.TrimSpace
	strwspace := " Hello Everyone! "
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace)) //移除空格

	// //strings.ToLower strings.ToUpper'
	// fmt.Println(strings.ToLower(strwspace))
	// fmt.Println(strings.ToUpper(strwspace))

	//strings.Repeat
	fmt.Println(strings.Repeat("foo",3)) //foofoofoo

	//strings.Count
	fmt.Println(strings.Count("Hello", "l")) //2
	fmt.Println(strings.HasPrefix("Hello", "he")) //false
	fmt.Println(strings.HasSuffix("Hello", "lo")) //true
	fmt.Println(strings.HasSuffix("Hello", "la"))

	//regexp
	str5 :="Hello,123 Go 11!"
	re:=regexp.MustCompile(`\d+`) //\d+代表一个和多个数字应该一起出现，不包括空格
	//-1 代表寻找该正则表达式的所有匹配项
	matches:=re.FindAllString(str5,-1)
	fmt.Println(matches) //[123 11]

	//unicode/utf8
	str6 := "Hello, 世界"
	fmt.Println(utf8.RuneCountInString(str6)) //9

	//strings.Builder
	var builder strings.Builder

	//write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	//Convert builder to a string
	resultString := builder.String()
	fmt.Println(resultString) //Hello, world!

	//using writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you")

	resultString = builder.String()

	fmt.Println(resultString) //Hello, world! How are you

	//Reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh!")
	resultString=builder.String()
	fmt.Println(resultString) //Starting fresh!
}
