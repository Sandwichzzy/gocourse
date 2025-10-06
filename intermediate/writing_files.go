package intermediate

import (
	"fmt"
	"os"
)

// Key Components
// `os`Package
// Functions
// Create(name string)(*File, error)
// OpenFile(name string, flag int, perm FileMode)(*File, error)
// Write(b Пbyte)(n int, err error)
// WriteString(s string)(n int, err error)

func main() {

	file,err:=os.Create("output.txt")
	if err!=nil {
		fmt.Println("Error creating file.",file)
		return
	}
	defer file.Close()

	// write data to file
	data:=  []byte("Hello World!\n")
	_,err=file.Write(data)
	if err!=nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data has been written to file successfully.")

	file,err = os.Create("writeString.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()


	//func (f *os.File) WriteString(s string) (n int, err error)
  // WriteString is like Write, but writes the contents of string s rather than a slice of bytes.
	_, err = file.WriteString("Hello Go!\n\n\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Writing to writeString.txt complete.")

}
