package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Epochs in Go")

	// 00:00:00 UTC Jan 1, 1970

	now := time.Now()
	unixTime := now.Unix()
	nsec := now.UnixMicro()
	fmt.Println("Current Unix time:", unixTime)

	// Converting a Unix Time to human readable format
	t := time.Unix(unixTime, nsec)
	fmt.Println(t)

	fmt.Println("Time:", t.Format("2006-01-02"))
}