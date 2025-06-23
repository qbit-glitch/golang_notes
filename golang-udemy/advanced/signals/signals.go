package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// signalsBasic()
	// signalsBasic2()
	// signalsBasic3()
	signalsBasic4()

}

func signalsBasic4(){
	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	// Notify channel on interrupt or terminate signals
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	go func(){
		sig := <- sigs
		fmt.Println("Received Signal:", sig)
		done <- true
	}()

	go func(){
		for {
			select {
			case <- done:
				fmt.Println("Stopping work due to signal.")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(time.Second)
			}
		}
	}()
	
	// Simulate some work
	for {
		time.Sleep(time.Second)
	}

}

func signalsBasic3() {

	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
	sigs := make(chan os.Signal, 1)

	// Notify channel interrupt or terminate signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGHUP, syscall.SIGUSR1)

	go func() {
		for sig := range sigs {
			switch sig {
			case syscall.SIGINT:
				fmt.Println("Received SIGINT (Interrupt)")
			// case syscall.SIGTERM:
			// 	fmt.Println("Received SIGTERM (Terminate)")
			case syscall.SIGHUP:
				fmt.Println("Received SIGHUP (Hangup)")

			case syscall.SIGUSR1:
				fmt.Println("Received SIGUSR1 (User Defined Signal 1)")
				fmt.Println("User defined functions is executed.")
				// continue
			}

			// fmt.Println("Graceful exit..")
			// os.Exit(0)
		}
	}()

	// Simulate some work
	for {
		fmt.Println("Working...")
		time.Sleep(time.Second)
	}
}

func signalsBasic2() {

	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
	sigs := make(chan os.Signal, 1)

	// Notify channel interrupt or terminate signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGUSR1)

	go func() {
		sig := <-sigs
		switch sig {
		case syscall.SIGINT:
			fmt.Println("Received SIGINT (Interrupt)")
		case syscall.SIGTERM:
			fmt.Println("Received SIGTERM (Terminate)")
		case syscall.SIGHUP:
			fmt.Println("Received SIGHUP (Hangup)")

		case syscall.SIGUSR1:
			fmt.Println("Received SIGUSR1 (User Defined Signal 1)")
			fmt.Println("User defined functions is executed.")
		}

		fmt.Println("Graceful exit..")
		os.Exit(0)
	}()

	// Simulate some work
	for {
		fmt.Println("Working...")
		time.Sleep(time.Second)
	}
}

func signalsBasic() {

	pid := os.Getpid()
	fmt.Println("Process ID:", pid)

	sigs := make(chan os.Signal, 1)

	// Notify channel interrupt or terminate signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println("Received Signal:", sig)
		fmt.Println("Graceful exit..")
		os.Exit(0)
	}()

	// Simulate some work
	for {
		fmt.Println("Working...")
		time.Sleep(time.Second)
	}
}

/*
tasklist - List of all processes on Windows
taskkil /F /PID <PID> : taskkill /F /PID 12345
Stop-Process -Id 12345 -Force
*/
