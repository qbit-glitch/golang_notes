package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	str := "Hello Go!"
	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	fmt.Println(str[0])
	fmt.Println(str[1:7])

	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(len(str3))

	// strings splitting
	fruits := "apple,banana,grapes"
	fruits1 := "apple-banana-grapes"
	part := strings.Split(fruits, ",")
	part1 := strings.Split(fruits1, "-")
	fmt.Println(part)
	fmt.Println(part1)


	strwspace := " Hello EveryOne! "
	fmt.Println(strwspace, "k")

	fmt.Println(strings.TrimSpace(strwspace), "k")
	
	fmt.Println(strings.Repeat("Foo", 3))

	fmt.Println(strings.Count("Hello", "l"))
	fmt.Println(strings.HasPrefix("Hello", "He"))
	fmt.Println(strings.HasPrefix("Hello", "he"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))
	fmt.Println(strings.HasSuffix("Hello", "la"))

	str5 := "Hello, 123 Go! 78"
    re := regexp.MustCompile(`\d+`)
    matches := re.FindAllString(str5, -1)
    fmt.Println(matches)

	stringBuilder()
}


// STRING BUILDER

func stringBuilder() {

	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(",")
	builder.WriteString("world!")
	
	result := builder.String()
	fmt.Println(result)

	// Using WriteRune to add a character sequentially
	builder.WriteRune(']')
	builder.WriteString("How are you!")

	result = builder.String()
	fmt.Println(result)

	// Reset the builder
	builder.Reset()

	builder.WriteString("Starting Fresh1")
	result = builder.String()
	fmt.Println(result)


}