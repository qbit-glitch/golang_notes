package main

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)


func main() {

	// processSpawning1()
	// processSpawning2()
	// processSpawning3()
	// processSpawning4()
	// processSpawning5()
	processSpawningPiping()
	processSpawning6()

}

func processSpawning2(){
	cmd := exec.Command("grep", "foo")
	
	// Set input for the command
	cmd.Stdin = strings.NewReader("food is good\ndrinks are awesome\ninside bar")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Output:", string(output))
}

func processSpawning1(){
	cmd := exec.Command("echo","Hello World")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Output:", string(output))
	
}

func processSpawning3(){
	cmd := exec.Command("sleep", "3")
	// Start the command
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error while starting command:", err)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error while waiting:", err)
	}

	fmt.Println("Process is Complete")
	
}


func processSpawning4(){
	cmd := exec.Command("sleep", "60")
	// Start the command
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error while starting command:", err)
	}
	time.Sleep(2 * time.Second)
	err = cmd.Process.Kill()
	if err != nil {
		fmt.Println("Error killing the process:", err)
	}
	fmt.Println("Process killed")
}

func processSpawning5(){
	cmd := exec.Command("printenv", "SHELL")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Output:", string(output))
	
}

func processSpawningPiping(){
	
	pr, pw := io.Pipe()
	
	cmd := exec.Command("grep", "foo")
	cmd.Stdin = pr

	go func(){
		defer pw.Close()
		pw.Write([]byte("food is good\ndrinks are awesome\nbar is ok"))
	}()	
	
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Output:", string(output))
	
}

func processSpawning6(){
	cmd := exec.Command("ls", "-la")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Output:", string(output))
}