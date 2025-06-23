package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("\n---- RANDOM NUMBERS ----\n")

	basics()
	diceGame()

}

func basics() {
	fmt.Println(rand.Intn(101))

	var seed1 = rand.New(rand.NewSource(20))
	fmt.Println("Random with seed as 20: ", seed1.Intn(1000)+1)

	var seed2 = rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println("Random with seed as time:", seed2.Intn(1000)+1)

	fmt.Println("Random floating point number between 0 and 1:", rand.Float64())

}

func diceGame() {

	var choice int
	for {
		fmt.Println("\nWelcome to the dice Game:")

		fmt.Println("1: Roll the dices")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)

		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid Choice..")
			continue
		}

		if choice == 2 {
			fmt.Println("Thanks for Playing! Goodbye...")
			break
		}

		dice1 := rand.Intn(6) + 1
		dice2 := rand.Intn(6) + 1

		fmt.Printf("Dices: %d, %d\n", dice1, dice2)
		fmt.Printf("Total: %d\n", dice1+dice2)

		fmt.Print("Do you want to play again (y/n): ")
		var isPlay string
		_, err = fmt.Scan(&isPlay)

		if err != nil || (isPlay != "y" && isPlay != "n") {
			fmt.Println("Invalid Input. Assuming no ...")
			break
		}

		if isPlay == "n" {
			fmt.Println("Thanks for Playing! Goodbye...")
			break
		}

	}
}
