package intermediate

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
		// fmt.Println(err)
	}
}

// 1.Key Concepts
// os.Mkdir
// os.MkdirAll
// os.ReadDir
// os.Chdir
// os.Remove
// os.RemoveAll
func main() {

	//  err := os.Mkdir("subdir", 0755)
	//  checkError(err)
	//  checkError(os.Mkdir("subdir1", 0755))

	//  defer os.RemoveAll("subdir1")

	// func os.WriteFile(name string, data []byte, perm os.FileMode) error
	// os.WriteFile("subdir1/file",[]byte(""),0755)

	// func os.MkdirAll(path string, perm os.FileMode) error
  // MkdirAll creates a directory named path, along with any necessary parents, and returns nil, or else returns an error.
	checkError(os.MkdirAll("subdir/parent/child", 0755))
	checkError(os.MkdirAll("subdir/parent/child1", 0755))
	checkError(os.MkdirAll("subdir/parent/child2", 0755))
	checkError(os.MkdirAll("subdir/parent/child3", 0755))

	os.WriteFile("subdir/parent/file", []byte(""), 0755)
	os.WriteFile("subdir/parent/child/file", []byte(""), 0755)

	//func os.ReadDir(name string) ([]os.DirEntry, error)
	// ReadDir reads the named directory, returning all its directory entries sorted by filename.
	result, err := os.ReadDir("subdir/parent")
	checkError(err)

	for _,entry :=range result{
		// child true d---------
		// child1 true d---------
		// child2 true d---------
		// child3 true d---------
		// file false ----------
		fmt.Println(entry.Name(),entry.IsDir(),entry.Type()) 
	}

	//func os.Chdir(dir string) error
	// Chdir changes the current working directory to the named directory. 改变工作目录
	checkError(os.Chdir("subdir/parent/child"))

	result,err=os.ReadDir(".")
	checkError(err)
	fmt.Println("Reading subdir/parent/child")
	for _,entry :=range result{
		fmt.Println(entry.Name()) 
	}

	checkError(os.Chdir("../../.."))
	dir,err:=os.Getwd()
	checkError(err)
  fmt.Println(dir) //  /home/sandwichzzy/goStu/gocourse


	// filepath.Walk and filepath.WalkDir

	pathfile := "subdir"
	fmt.Println("Walking Directory")

	err=filepath.WalkDir(pathfile, func(path string,d fs.DirEntry, err error) error {
		if err!=nil{
			fmt.Println("Error:",err)
			return err
		}
		fmt.Println(path)
		return nil
	})

	checkError(err)

	checkError(os.RemoveAll("subdir"))
	checkError(os.Remove("subdir1"))


}
