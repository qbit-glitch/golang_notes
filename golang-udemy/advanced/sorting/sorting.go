package main

import (
	"fmt"
	"sort"
)

type Person struct{
	Name string
	Age int
}

type By func(p1, p2 *Person) bool

type personSorter struct{
	people []Person
	by func(p1, p2 *Person) bool
}

func (s *personSorter) Len() int{
	return len(s.people)
}

func (s *personSorter) Less(i, j int) bool{
	return s.by(&s.people[i], &s.people[j])
}

func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

func (by By) Sort(people []Person){
	ps := &personSorter{
		people: people,
		by: by,
	}
	sort.Sort(ps)
}

func main() {

	numbers := []int{5,3,4,1,2}
	sort.Ints(numbers)
	fmt.Println("Sorted Numbers:", numbers)

	stringSlices := []string{"John", "Anthony", "Steve", "Victor", "Walter"}
	sort.Strings(stringSlices)
	fmt.Println("Sorted Strings:",stringSlices)

	fmt.Println("-------- CUSTOM SORTING ---------")
	people := []Person{
		{"Alice", 30},
		{"Bob", 24},
		{"Anna", 31},
	}

	fmt.Println("Unsorted by Age:", people)

	ageAsc := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	By(ageAsc).Sort(people)
	/* Since age have the same function signature as the
		By type, so we have converted the age function to
		type `By` function and called the `Sort()` method 
		associated to type `By`.
	*/
	fmt.Println("Sorted by age (ascending):", people)

	ageDesc := func(p1, p2 *Person) bool {
		return p1.Age > p2.Age
	}
	By(ageDesc).Sort(people)
	fmt.Println("Sorted by Age (descending):", people)


	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}
	By(name).Sort(people)
	fmt.Println("Sorted by Name:", people)

	lenName := func(p1, p2 *Person) bool {
		return len(p1.Name) < len(p2.Name)
	}
	By(lenName).Sort(people)
	fmt.Println("Sorted by length of Name:", people)


	// ---------- SORT.SLICE --------------
	stringSlice := []string{"banana", "apple", "cherry", "grapes", "gauva"}
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})
	fmt.Println("Sorted by last character:", stringSlice)
}