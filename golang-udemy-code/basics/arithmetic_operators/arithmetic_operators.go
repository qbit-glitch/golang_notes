package main

import (
	"fmt"
	"math"
)

func main(){
	// Overflow with signed integers
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	// Overflow with un-signed integers
	var umaxInt uint64 = 18446744073709551615
	fmt.Println(umaxInt)

	umaxInt = umaxInt + 1
	fmt.Println(umaxInt)

	// Underflow
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)


}