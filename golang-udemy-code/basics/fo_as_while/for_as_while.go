package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(10)+1

	var guess int
	fmt.Println("Target number is between 1 and 10")
	for {
		
		fmt.Println("Enter your guess:")
		fmt.Scanln(&guess)

		if guess == target{
			fmt.Println("Congratulations ...")
			break
		} else if guess > target {
			fmt.Println("Value above target .. Try to guess a lesser number")
		} else {
			fmt.Println("Sorry ... Try to guess a bigger number")
		}
		

	}

}