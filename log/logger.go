package log

import (
	"os"
	"log"
)

var Logger *log.Logger
var logFile * os.File

func init() {
	var err error
	logFile, err = os.OpenFile("log.txt", os.O_WRONLY | os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	Logger = log.New(logFile, ":", log.Lmsgprefix | log.Lmicroseconds | log.Lshortfile)

	// logger.SetOutput(os.Stdout)
}
