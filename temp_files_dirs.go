package main

import (
	"fmt"
	"os"
)

func checkError(err error){
	if err!=nil {
		panic(err)
	}
}

func main() {
	// tempFileName := "temporaryFile"

	// tempFile, err := os.CreateTemp("", tempFileName)
	// checkError(err)

	// fmt.Println("Temporary file created:", tempFile.Name())

	// defer os.Remove(tempFile.Name())
	// defer tempFile.Close()

	tempDir,err:=os.MkdirTemp("","GoCourseTempDir")

	checkError(err)

	defer os.RemoveAll(tempDir)

	fmt.Println("Temporary directory created:", tempDir)
}
