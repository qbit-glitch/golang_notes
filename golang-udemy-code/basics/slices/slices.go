package main

import "fmt"

func main() {

	numbers := []int{1,2,3,4,5,6,6,7,7}
	fmt.Println(numbers)

	a := [5]int{1,2,3,4,5}
	slice1 := a[1:4]

	fmt.Println(slice1)

	slice1 = append(slice1, 8,9,10)
	fmt.Println("Slice1:", slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("sliceCopy:", sliceCopy)

	// var nilSlice []int

	// Iterators for slices using range based for loops
	for i,v := range slice1 {
		fmt.Println(i,v)
	}

	twoD := make([][]int, 3)
	for i:=0;i<3;i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j:=0; j<innerLen; j++ {
			twoD[i][j] = i+j
			fmt.Printf("twoD[%d][%d] = %d + %d = %d\n", i,j,i,j, twoD[i][j])
		}
	}
	fmt.Println(twoD)

	// slice[low:high]
	slice2 := slice1[2:4]
	fmt.Println(slice2)

	fmt.Println("The capacity of slice2 is", cap(slice2))
	fmt.Println("The length of slice2 is", len(slice2))
}