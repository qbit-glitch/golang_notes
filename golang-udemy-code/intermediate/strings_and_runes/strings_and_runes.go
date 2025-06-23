package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello\nGo!"
	message2 := "Hello,\tGo!"
	message3 := "Hello,\rGo!"

	rawMessage := `Hello,\nGo!`

	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println(rawMessage)

	fmt.Println("Length of message variable is:", len(message))
	fmt.Println("Length of message3 variable is:", len(message3))
	fmt.Println("Length of rawMessage variable is:", len(rawMessage))

	fmt.Println("The first character in message var is:", message[0])

	greeting := "Hello"
	name := "Alice"
	fmt.Println(greeting+name)

	str1 := "Apple"		// A has an ASCII value of 65
	str2 := "banana"	// b has an ASCII value of 98
	str3 := "app"		// a has an ascii value of 97

	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)

	for i, char := range message {
		fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%d: %x\n", i, char)	// hexadecimal value of character
		fmt.Printf("%d: %v\n",i,char)   // Default value as ascii value
	}

	fmt.Println("Rune Count:", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	var jch rune = '日'

	fmt.Println(ch, jch)
	fmt.Printf("%c, %c\n", ch, jch)

	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of cstr is %T\n", cstr)

	const NIHONGO = "会ったり"
	fmt.Println(NIHONGO)

	jhello := "おはよう"   // Japanese hello

	for _, runeValue := range jhello {
		fmt.Printf("%c : %v\n", runeValue, runeValue)
	}

	smiley := '😁'
	fmt.Printf("%c\n", smiley)
}
