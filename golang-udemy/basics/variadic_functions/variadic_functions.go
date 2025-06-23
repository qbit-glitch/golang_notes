package main

import "fmt"

func main() {

	sequence, total := sum(1, 10,20,30,40,50)
	fmt.Printf("Sequence: %d, Total: %d\n", sequence, total)


	sequence2, total2 := sum(2, 110,120,130,140,150)
	fmt.Printf("Sequence: %d, Total: %d\n", sequence2, total2)

	numbers := []int{1,2,4,5,6,9}
	sequence3, total3 := sum(3, numbers...)
	fmt.Printf("Sequence: %d, Total: %d\n", sequence3, total3)
}

func sum(sequence int, nums ...int) (int, int){
	total := 0
	for _,v := range nums{
		total += v
	}
	return sequence, total
}