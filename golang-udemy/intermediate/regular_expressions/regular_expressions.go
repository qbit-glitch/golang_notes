package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("He said, \"I'm great\"")
	fmt.Println(`He said, "I'm great"`)

	// compile a regex pattern to match email address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}`)

	// Test strings
	email1 := "username@email.com"
	email2 := "invalid_email"

	// Match
	fmt.Println("Email1:", re.MatchString(email1))
	fmt.Println("Email2:", re.MatchString(email2))

	// Capturing Groups
	// Compile a regex pattern to capture date components
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	// Test String
	date := "2024-07-30"

	// Find all submatches
	submatches := re.FindStringSubmatch(date)

	fmt.Println(submatches)
	fmt.Println(len(submatches))

	for i, v := range submatches {
		fmt.Println(i, ":", v)
	}

	// Target string
	str := "Hello World"

	re = regexp.MustCompile(`[aeiou]`)

	result := re.ReplaceAllString(str, "-")
	fmt.Println(result)

	// OPTIONS AND FLAGS

	re = regexp.MustCompile(`(?i)go`)
	re1 := regexp.MustCompile(`go`)
	// Test string
	text := "Golang is great"

	// Match
	fmt.Println("Match:", re.MatchString(text))
	fmt.Println("Match:", re1.MatchString(text))

}
