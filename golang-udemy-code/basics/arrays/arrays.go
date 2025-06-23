package main

import "fmt"

func main() {

	fruits := [4]string{"Apple", "Mango", "Banana", "Grapes"}
	for index, value := range fruits {
		fmt.Printf("Index: %d | Value: %s\n", index, value)
	}

	fmt.Println("Fruits Array: ")
	for _, value := range fruits {
		fmt.Printf("Value: %s\n", value)
	}
	
	fmt.Println("The length of Fruits array is:", len(fruits))

	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	fmt.Println("Array1 equal to Array2:", array1 == array2)

	var matrix [3][3]int = [3][3]int {
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println("Matrix:", matrix)


	original := [3]int{1, 2, 3}
	var copied [3]int = original

	copied[2] = 100

	fmt.Println("Original Array:", original)
	fmt.Println("Copied Array:", copied)



	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int = &originalArray

	copiedArray[2] = 100

	fmt.Println("Original Array:", originalArray)
	fmt.Println("Copied Array:", *copiedArray)

}
