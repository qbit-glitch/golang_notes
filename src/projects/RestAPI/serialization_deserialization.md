# Serialization / Deserialization - Marshal/Unmarshal - Encode/Decode

Serialization is the process of converting a Go object into a json string. A json string is a byte slice and we are converting a Go object, an instance of a streuct into a JSON string which is a byte slice. Deserialization is a reverse process. It converts a JSON string into a Go object. Go provides two primary ways to handle json :
- `json.Marshal` and `json.Unmarshal`: these functions are straight forward and commonly used for in-memory json processing.
- `json.NewEncoder` and `json.NewDecoder`: these methods are used for streaming JSON data. These are ideal for handling large datasets or working with network connections.



json.Marshal and json.Unmarshal are best suited for situations where you need to quickly serialize data in memory. They are simple to use and perfect for small to medium sized datasets.

json.NewDecoder creates a new decoder that reads from io.Reader. It's particularly used for streaming data such as reading json from a network connection or a file.

json.NewEncoder and json.NewDecoder, are ideal for situations involving large datasets or for streaming data. They are more efficient in terms of memory usage, especially when dealing with data that is being read from or written to an external source. And mostly when we are making APIs, we are reading or writing data from or to an external source and dealing with large datasets.

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{Name:"Alice", Email: "alice@example.com"}
	fmt.Println(user)
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))

	var user1 User
	err = json.Unmarshal(jsonData, &user1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User created from json data:", user1)

	
	// json.NewDecoder and json.NewEncoder
	jsonData1 := `{"name": "John", "email": "john@example.com"}`
	reader := strings.NewReader(jsonData1)
	decoder := json.NewDecoder(reader)

	var user2 User
	err = decoder.Decode(&user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user2)

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)

	err = encoder.Encode(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encoded json string:", buf.String())
}
```

Encoder: 
- converts a struct to a json string
- first, create an encoder
- then encode the struct

Decoder: 
- converts a json string to a struct
- first create a decoder
- then decode the json string
