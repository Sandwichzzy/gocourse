package intermediate

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//在Go语言中，io.Reader和io.Writer是两个非常基础的接口，它们定义了读写操作的基本契约。
// 很多标准库中的类型都实现了这两个接口，包括os、bufio、bytes、strings等包中的类型。

func readFromReader(r io.Reader){
	buf := make([]byte,1024)
	n,err :=r.Read(buf)
	if err!=nil {
		log.Fatalln("Error reading from reader", err)
	}
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string){
	_,err:=w.Write([]byte(data))
	if err!=nil {
		log.Fatalln("Error reading from reader.",err)
	}
}



func closeResource(c io.Closer){
	err:=c.Close()
	if err!=nil {
		log.Fatalln("Error reading from reader",err)
	}
}


// 特性	      bufio.Writer	         bytes.Buffer
// 设计目的	缓冲写入，减少系统调用	内存字节序列构建
// 底层存储	包装另一个 Writer	     自己的字节切片
// 性能特点	减少 I/O 操作次数	     快速内存操作
// 主要用途	文件、网络 I/O 优化	   字符串/字节拼接
// 自动刷新	需要手动或自动刷新	      立即生效
func bufferExample() {
	var buf bytes.Buffer //(值类型 bytes.Buffer) //stack
	buf.WriteString("Hello Buffer!")
	fmt.Println(buf.String())
}

func multiReaderExample(){
	r1:=strings.NewReader("Hello ")
	r2:=strings.NewReader("World!")
	mr:=io.MultiReader(r1,r2)
	//使用new函数分配内存，返回一个指向新分配的bytes.Buffer类型零值的指针，即*bytes.Buffer类型。
	//然后，这个指针被赋值给变量buf。所以buf是一个指针，指向一个空的bytes.Buffer。
	//实际上，bytes.Buffer的方法都是定义在指针接收器上的，所以即使我们使用值类型声明的变量，
	//在调用方法时，Go语言会自动取地址，所以两种方式在调用方法时没有区别：

	// 特性	    var buf bytes.Buffer	          buf := new(bytes.Buffer)
	// 类型	    bytes.Buffer (值类型)	          *bytes.Buffer (指针类型)
	// 内存分配	 栈上(通常)	                          堆上
	// 零值	    已初始化的空 Buffer	               指向零值 Buffer 的指针
	// 直接使用	 可以立即使用	                   可以立即使用(自动解引用)
	// nil 检查	 永远不会为 nil	                    可能为 nil
	buf:=new(bytes.Buffer) //*bytes.Buffer (指针类型)
	_,err:=buf.ReadFrom(mr)
	if err!=nil {
		log.Fatalln("Error reading from reader",err)
	}
	fmt.Println(buf.String())
}

func pipeExample(){
	//Pipe 创建了一个同步的内存管道。
	//它可用于连接需要 io.Reader 的代码与需要 io.Writer 的代码。
	//每次对 [PipeWriter] 的写入都会阻塞，直到它满足来自 [PipeReader] 的一次或多次读取，
	//且这些读取完全消耗了已写入的数据。
	//数据会直接从写入端复制到对应的读取端（或多次读取端）
	pr,pw:=io.Pipe()
	//使用 go routine 解决 阻塞
	go func() {
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()
	buf :=new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func writeToFile(filepath string,data string){
	file,err:= os.OpenFile(filepath,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err!=nil {
		log.Fatalln("Error opening/creating file:",err)
	}
	defer closeResource(file)

	_,err =file.Write([]byte(data))
	if err!=nil {
		log.Fatalln("Error opening/creating file:",err)
	}
	


	//os.File implements the same write func:	func (f *File) write(b []byte) (n int, err error)
	// writer:=io.Writer(file)
	// _,err=writer.Write([]byte(data))
	// if err!=nil {
	// 	log.Fatalln("Error opening/creating file:",err)
	// }
}

func main() {

	fmt.Println("=== Read from Reader ===")
	readFromReader(strings.NewReader("Hello Reader!"))

	fmt.Println("=== Write to Writer ===")
	var writer bytes.Buffer 
	writeToWriter(&writer,"Hello,writer")
	fmt.Println(writer.String())


	fmt.Println("=== Buffer Example ===")
	bufferExample()

	fmt.Println("=== Multi Reader Example ===")
	multiReaderExample()

	fmt.Println("=== Pipe Example ===")
	pipeExample()

	filepath := "io.txt"
	writeToFile(filepath,"Hello File!")

	resource := &MyResource{name: "TestResource"}
	closeResource(resource)
	
}

type MyResource struct {
	name string
}

func (m MyResource) Close() error{
	fmt.Println("Closing resource:",m.name)
	return nil
}