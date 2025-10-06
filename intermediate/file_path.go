package intermediate

import (
	"fmt"
	"path/filepath"
	"strings"
)

// 1. Key Concepts
// Absolute Path
// Relative Path
// 2. "path/filepath` Package
// -Functions
// filepath.Join
// filepath.Split
// filepath.Clean
// filepath.Abs
// filepath.Base
// filepath.Dir

func main() {
	relativePath:="./data/file.txt"
	absolutePath:="/home/user/docs/file.txt"

	// Join paths using filepath.join
	joinedPath := filepath.Join("home", "Documents", "downloads", "file.zip")
	fmt.Println("Joined Path:", joinedPath) //Joined Path: home/Documents/downloads/file.zip

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized Path:", normalizedPath) //Normalized Path: data/file.txt
	

	dir,file:=filepath.Split("/home/user/docs/file.txt")
	fmt.Println("File:", file) //File: file.txt
	fmt.Println("Dir:", dir) //Dir: /home/user/docs/
	fmt.Println(filepath.Base("/home/user/docs/")) //docs

	fmt.Println("Is relativePath variable absolute:", filepath.IsAbs(relativePath)) //false
	fmt.Println("Is absolutePath variable absolute:", filepath.IsAbs(absolutePath)) //true

	//func filepath.Ext(path string) string
  //Ext returns the file name extension used by path. 
	fmt.Println(filepath.Ext(file)) //.txt
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file))) //file

	
	//func filepath.Rel(basepath string, targpath string) (string, error)
  //Rel returns a relative path that is lexically equivalent to targpath when joined to basepath with an intervening separator. 
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel) //t/file 

	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel) //../b/t/file


	// func filepath.Abs(path string) (string, error)
	// Abs returns an absolute representation of path. 
	// If the path is not absolute it will be joined with the current working directory to turn it into an absolute path. 
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Absolute Path:", absPath)  //Absolute Path: /home/sandwichzzy/goStu/gocourse/data/file.txt
	}
}
