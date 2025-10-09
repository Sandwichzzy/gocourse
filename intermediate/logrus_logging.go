package intermediate

import "github.com/sirupsen/logrus"
func main() {
	log:=logrus.New()
	//set log level (等级筛选)
	log.SetLevel(logrus.InfoLevel)
	//set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	//Logging examples
	log.Info("This is an info message.")
	log.Warn("This is a warning message.")
	log.Error("This is an error message.")

	log.WithFields(logrus.Fields{
		"username": "John Doe",
		"method":   "GET",
	}).Info("User logged in.")

// 	{"level":"info","msg":"This is an info message.","time":"2025-10-09T02:36:25Z"}
// {"level":"warning","msg":"This is a warning message.","time":"2025-10-09T02:36:25Z"}
// {"level":"error","msg":"This is an error message.","time":"2025-10-09T02:36:25Z"}
// {"level":"info","method":"GET","msg":"User logged in.","time":"2025-10-09T02:36:25Z","username":"John Doe"}
}
