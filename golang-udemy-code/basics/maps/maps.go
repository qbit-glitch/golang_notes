package main

import (
	"fmt"
	"maps"
)

func main() {

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 90
	fmt.Println(myMap)

	delete(myMap, "code")
	fmt.Println(myMap)

	myMap2 := map[string]int{"a": 1, "b": 2}

	myMap3 := map[string]int{"a": 1, "b": 2}

	if maps.Equal(myMap2, myMap3) {
		fmt.Println("Maps myMay2 and myMap3 are equal")
	}

	for k, v := range myMap3 {
		fmt.Println(k, ":", v)
	}

	_, ok := myMap["key1"]
	if ok {
		fmt.Println("A value exists with key1")
	} else {
		fmt.Println("No value exists with key1")
	}

	var myMap4 map[string]string
	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value")
	} else {
		fmt.Println("The map is not initialized to nil value")
	}
	val := myMap4["key1"]
	fmt.Println(val)

	// if we initialize a map in the myMap4 way ... then we cannot insert key and values using the myMap4["Key"] = "Value"
	// Instead we have to use the `make` method

	myMap4 = make(map[string]string)
	myMap4["Key"] = "Value"
	fmt.Println(myMap4)
}
