# `sync.Pool`

`sync.Pool` is a type provided by the go standard library in the sync package. It implements a pool of reusable objects. The primary purpose of `sync.Pool` is to reduce the overhead of allocating and deallocating objects frequenty by providing a pool where objects can be reused.

**Why does it matter ?**
Because object allocation and garbage collection can be expensive, especially in high performance applications or scenarios with frequent allocations. `sync.Pool` helps maintaining this by mitigating a pool of objects that can be reused, reducing the need for frequent allocations and garbage collection.

- Key Concepts of `sync.Pool`:
    - Object Pooling
    - Object Retrieval and Return

- Methods of `sync.Pool`:
    - `Get()`
    - `Put(interface{})`
    - `New(Optional)`

- It works on the LIFO principle.
- The new field will create a new instance if the object pool is empty.

## Code

```go
package main

import (
	"fmt"
	"sync"
)

type person struct{
	name string
	age int
}

func main() {
	poolWithNew()
	poolWithoutNew()

}

func poolWithoutNew(){

	var pool = sync.Pool{}
	pool.Put(&person{name: "John", age: 26})
	person1 := pool.Get().(*person)

	fmt.Println("Person 1:", person1)


	fmt.Printf("Person1: Name: %s | Age: %d\n", person1.name, person1.age)

	pool.Put(person1)
	fmt.Println("Returned Person to Pool")

	person2 := pool.Get().(*person)
	fmt.Println("Got Person 2:", person2)

	person3 := pool.Get()
	if person3 != nil {
	fmt.Println("Got Person 3:", person3)
	person3.(*person).name = "James"
	} else {
		fmt.Println("Sync Pool is empty. So person3 is not assigned anything")
	}

	// Returning object to the pool again
	pool.Put(person2)
	pool.Put(person3)

	person4 := pool.Get().(*person)
	fmt.Println("Got Person 4:", person4)

	person5 := pool.Get()
	if person5 != nil {
	fmt.Println("Got Person 3:", person5)
	person5.(*person).name = "James"
	} else {
		fmt.Println("Sync Pool is empty. So person5 is not assigned anything")
	}
}

func poolWithNew(){

	var pool = sync.Pool{
		New: func() interface{}{
			fmt.Println("Creating a new Person")
			return &person{}
		},
	}

	// Get an Object from the pool
	person1 := pool.Get().(*person)
	person1.name = "John"
	person1.age = 18
	fmt.Println("Person 1:", person1)

	fmt.Printf("Person1: Name: %s | Age: %d\n", person1.name, person1.age)

	pool.Put(person1)
	fmt.Println("Returned Person to Pool")

	person2 := pool.Get().(*person)
	fmt.Println("Got Person 2:", person2)

	person3 := pool.Get().(*person)
	fmt.Println("Got Person 3:", person3)
	person3.name = "James"

	// Returning object to the pool again
	pool.Put(person2)
	pool.Put(person3)

	person4 := pool.Get().(*person)
	fmt.Println("Got Person 4:", person4)

	person5 := pool.Get().(*person)
	fmt.Println("Got Person 5:", person5)
}

```

## Key Notes:

- **Best Practices for using `sync.Pool`**:
    - Use for expensive object allocations
    - Keep Objects in Pool Clean
    - Avoid Complex Objects

- **Advanced Use Cases**
    - Reusing Buffers
    - Managing Database Connections
    - High Performance Applications

- **Considerations and Limitations**
    - Garbafe Collection
    - Not for Long-Lived Objects
    - Thread Safety