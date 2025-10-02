package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

// Key Components
// bufio.Reader
// func NewReader(rd io.Reader)*Reader
// func (r *Reader) Read(p lbyte)(n int, err error)
// func (r *Reader) ReadString(delim byte) (line string, err error)
// bufio.Writer
// func NewWriter(wr io.Writer)*Writer
// func (w *Writer) Write(p llbyte)(n int, err error)
// func (w *Writer) WriteString(s string)(n int, err error)

func main() {

	//  reader:=bufio.NewReader(strings.NewReader("Hello,bufio packageeeeeee!\n How are you doing?"))

	 //reading byte slice
	//  data:=make([]byte,20)
	//  n,err:=reader.Read(data)
	//  if err!=nil{
	// 	fmt.Println("error reading:",err)
	// 	return
	//  }
	//  fmt.Printf("Read %d bytes:%s\n",n,data[:n]) //Read 20 bytes:Hello,bufio packagee
	 

	//  line,err:=reader.ReadString('\n')
	//  if err!=nil {
	// 	fmt.Println("Error reading string:",err)
	// 	return
	//  }
	//  fmt.Println("Read string:",line) //Read string: eeeee! 

	//------------------writter----------------------------
	writer :=bufio.NewWriter(os.Stdout)
	// Writing byte slice
	data := []byte("Hello, bufio package!\n")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)
	
	// Flush the buffer to ensure all data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	// Writing string
	str := "This is a string.\n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes.\n", n)

	// Flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

}