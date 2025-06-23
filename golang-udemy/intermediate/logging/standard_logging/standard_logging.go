package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	log.Println("This a logging message.")

	log.SetPrefix("INFO: ")
	log.Println("This is an info message.")

	// Log Flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log message with date, time and filename.")

	infoLogger.Println("This is a info message.")
	warnLogger.Println("This a warning message.")
	errorLogger.Println("This is an error message.")

	fmt.Println("------------------------------------")

	file, err := os.OpenFile("app.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()

	infoLogger1 := log.New(file, "INFO: ", log.Ldate | log.Ltime | log.Lshortfile)
	errorLogger1 := log.New(file, "ERROR: ", log.Ldate | log.Ltime | log.Lshortfile)
	warnLogger1 := log.New(file, "WARNING: ", log.Ldate | log.Ltime | log.Lshortfile)
	debugLogger1 := log.New(file, "DEBUG: ", log.Ldate | log.Ltime | log.Lshortfile)
	
	infoLogger1.Println("This is an info message.")
	errorLogger1.Println("This is an error message.")
	warnLogger1.Println("This is a warning message")
	debugLogger1.Println("THis is a debug message.")
	
}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
