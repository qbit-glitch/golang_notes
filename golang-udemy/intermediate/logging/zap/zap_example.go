package main

import (
	"fmt"

	"go.uber.org/zap"
)


func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("error in intializing zap logger:", err)
		return
	}
	defer logger.Sync()       // to flush out the buffer with zap logger

	logger.Info("This is an info message:")
	logger.Info("User logged in", zap.String("username", "John Doe"), zap.String("method", "get"))
}