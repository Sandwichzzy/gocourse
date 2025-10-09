package intermediate

import (
	"log"

	"go.uber.org/zap"
)

func main() {

	logger,err:=zap.NewProduction()
	if err != nil {
		log.Println("Error in intializing Zap logger")
	}

	defer logger.Sync() //函数结束时刷新缓冲buffer
	logger.Info("This is an info message") //{"level":"info","ts":1759977741.3744855,"caller":"gocourse/zap_log.go:17","msg":"This is an info message"}

	
	logger.Info("User logged in", zap.String("username", "John Doe"), zap.String("method", "GET")) 
	//{"level":"info","ts":1759977791.1141815,"caller":"gocourse/zap_log.go:20","msg":"User logged in","username":"John Doe","method":"GET"}
}
