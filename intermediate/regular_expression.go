package intermediate

import (
	"fmt"
	"regexp"
)

// 元字符
//.	匹配任意单个字符（默认不包含换行符 \n）	a.c 匹配 "abc", "a&c"
// *	匹配前面的字符或子表达式 0 次或多次	ab*c 匹配 "ac", "abc", "abbc"
// +	匹配前面的字符或子表达式 1 次或多次	ab+c 匹配 "abc", "abbc"（不匹配 "ac"）
// ?	匹配前面的字符或子表达式 0 次或 1 次	ab?c 匹配 "ac", "abc"
//{n,}	匹配 至少 n 次	ab{2,}c 匹配 "abbc", "abbbc"...
// ^	匹配字符串的开始（在多行模式下匹配行首）	^abc 匹配以 "abc" 开头的字符串
// $	匹配字符串的结束（在多行模式下匹配行尾）	xyz$ 匹配以 "xyz" 结尾的字符串
// \	转义字符，用于匹配元字符本身	\. 匹配字面量的点号 "."

// 字符类	说明	示例
// [abc]	匹配 a, b, 或 c 中的任意一个字符	[aeiou] 匹配任意一个元音字母
// [a-z]	匹配从 a 到 z 的任意小写字母	[0-9] 匹配任意一个数字
// [^abc]	否定字符类，匹配不在 a, b, c 中的任意字符	 [^0-9] 匹配任意非数字字符
// \d	匹配数字，等价于 [0-9]	\d+ 匹配一个或多个数字
// \D	匹配非数字，等价于 [^0-9]
// \s	匹配任意空白字符（空格、制表符 \t、换行符 \n 等）
// \S	匹配任意非空白字符
// \w	匹配单词字符（字母、数字、下划线），等价于 [a-zA-Z0-9_]	\w+ 匹配一个单词
// \W	匹配非单词字符，等价于 [^a-zA-Z0-9_]
func main() {

	fmt.Println("He said,\"I am great\"" )
	fmt.Println(`He said,"I am great"` )

	// complie a regex pattern to match email address
	re:=regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	//test strings
	email1:="user@email.com"
	email2:="invaslide_email"

	//Match
	fmt.Println("email:",re.MatchString(email1))
	fmt.Println("email:",re.MatchString(email2))

	//capturing groups
	//compile a regex pattern to capture data components

	//使用圆括号 () 可以创建捕获组。
	// 分组：将多个字符视为一个整体进行量化操作。例如，(ab)+ 匹配 "ab", "abab", "ababab"。
	// 捕获：被匹配到的子字符串会被保存下来，可以通过索引来获取。
	re=regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	date:="2024-07-30"
	submatches:=re.FindStringSubmatch(date)
	fmt.Println(submatches) //[2024-07-30 2024 07 30]
	fmt.Println(submatches[0]) //2024-07-30
	fmt.Println(submatches[1]) //2024
	fmt.Println(submatches[2]) // 07
	fmt.Println(submatches[3]) //30

	//target string
	str:="Hello world"
	re=regexp.MustCompile(`[aeiou]`)
	result :=re.ReplaceAllString(str,"*")
	fmt.Println(result) //H*ll* w*rld

	//  标志（Flags）
	// 	Go 的正则表达式在编译时通过给 Compile 函数传递特定的字符串来设置标志。

	// i：大小写不敏感。
	// m：多行模式（使 ^ 和 $ 匹配每一行的开始和结束，而不仅是整个字符串）。
	// s：让 . 匹配包括换行符 \n 在内的所有字符。

	//前面要加？
	// re=regexp.MustCompile(`(?i)go`)
	re=regexp.MustCompile(`go`)
	//test string
	text:="Golang is great"
	fmt.Println("Match:",re.MatchString(text))

}
