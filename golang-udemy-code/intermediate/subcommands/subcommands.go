package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	stringFlag := flag.String("user", "Guest", "Name of the user")
	flag.Parse()
	fmt.Println(stringFlag)

	subcommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subcommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	// fmt.Println(subcommand1)
	// fmt.Println(subcommand2)

	firstFlag := subcommand1.Bool("processing", false, "Command Processing status")
	secondFlag := subcommand1.Int("bytes", 1024, "Command Processing status")

	flagsc2 := subcommand2.String("language", "Go", "Enter your language")

	// fmt.Println(firstFlag)
	// fmt.Println(secondFlag)
	// fmt.Println(flagsc2)

	if len(os.Args) < 2 {
		fmt.Println("This Program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1]{
	case "firstSub":
		subcommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1:")
		fmt.Println("processing:", *firstFlag)
		fmt.Println("bytes:", *secondFlag)

	case "secondSub":
		subcommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2:")
		fmt.Println("language:", *flagsc2)
	default :
		fmt.Println("no subcommand entered")
		os.Exit(1)
	}
	



}