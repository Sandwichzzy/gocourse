package intermediate

//embed 指令是一个在 Go 1.16 版本中引入的特性，它允许将外部文件和目录直接嵌入到 Go 程序中，编译到最终的可执行文件中。
// import "embed"
// //go:embed <文件或目录路径>
// var <变量名> <类型>

import (
	"embed" //blank import
	"fmt"
	"io/fs"
	"log"
)

//go:embed example.txt
var content string

//go:embed basics
var basicsFolder embed.FS

func main() {

		fmt.Println("Embedded content:", content)
	  content,err:=basicsFolder.ReadFile("basics/hello.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println("Embedded file content:", string(content))

		err =fs.WalkDir(basicsFolder,"basics",func(path string, d fs.DirEntry, err error) error {
			if err !=nil {
				fmt.Println(err)
				return err
			} 
			fmt.Println(path)
			return nil 
		})

		if err != nil {
			log.Fatal(err)
		}
}
