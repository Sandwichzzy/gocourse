package intermediate

import (
	"log"
	"os"
)


func main() {

	log.Println("This is a log message.")

	log.SetPrefix("INFO: ")

	log.Println("This is an info message")

	//log flags

	log.SetFlags(log.Ldate | log.Ltime |log.Llongfile )
	log.Println("this is a log message with date, time,file.")

	infoLogger.Println("This is an info message")
	warnLogger.Println("This is a warning message.")
	errorLogger.Println("This is an wrror message.")
 
	file,err:=os.OpenFile("app.log",os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err !=nil {
		//fatalf 打印错误然后退出程序 
		log.Fatalf("Failed to open log file:%v",err)
	}
	defer file.Close()

	warnLogger1 := log.New(file, "WARN: ", log.Ldate|log.Ltime)
	infoLogger1 := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger1 := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger :=log.New(file,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger.Println("This is a debug message.") //INFO: 2025/10/09 02:16:39 logging.go:34: This is a debug message
	warnLogger1.Println("This is a warning message.") //WARN: 2025/10/09 02:19:02 This is a warning message.
	infoLogger1.Println("This is an info mesage.") //INFO: 2025/10/09 02:19:02 logging.go:40: This is an info mesage.
	errorLogger1.Println("This is an error.") //ERROR: 2025/10/09 02:19:02 logging.go:41: This is an error.

}

var (
	infoLogger = log.New(os.Stdout,"INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger =log.New(os.Stdout,"WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout,"ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
