package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	 reader:=bufio.NewReader(strings.NewReader("Hello,bufio packageeeeeee!\n How are you doing?"))

	 //reading byte slice
	 data:=make([]byte,20)
	 n,err:=reader.Read(data)
	 if err!=nil{
		fmt.Println("error reading:",err)
		return
	 }
	 fmt.Printf("Read %d bytes:%s\n",n,data[:n]) //Read 20 bytes:Hello,bufio packagee
	 

	 line,err:=reader.ReadString('\n')
	 if err!=nil {
		fmt.Println("Error reading string:",err)
		return
	 }
	 fmt.Println("Read string:",line) //Read string: eeeee! 

	//------------------writter----------------------------
	writer :=bufio.NewWriter(os.Stdout)
	// Writing byte slice
	data1:=[]byte("Hello,bufio package!\n")
	m,err:=writer.Write(data1)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes \n",n)

}