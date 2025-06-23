package main

import "github.com/sirupsen/logrus"


func main() {

	log := logrus.New()

	// Set the log level
	log.SetLevel(logrus.InfoLevel)

	// Set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	// Logging examples
	log.Info("This is an info message.")
	log.Warn("This is a warning message.")
	log.Error("This is an error message.")

	log.WithFields(logrus.Fields{
		"username": "John",
		"method" : "GET",
	}).Info("User logged in.")

}