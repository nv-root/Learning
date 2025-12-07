package main

import (
	"log"
	"os"
)

var infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var warnLogger = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
var errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)

func NativeLogging() {
	log.Println("This is a log message")
	log.SetPrefix("INFO: ")
	log.Println("This is an info message")

	// flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log message with")

	infoLogger.Println("This is from the info logger")
	warnLogger.Println("This is from the warn logger")
	errorLogger.Println("This is from the error logger")

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	debugLogger := log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	defer logFile.Close()
	debugLogger.Println("This is a debug message")

}
