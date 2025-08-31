## Protol Buffers

Language agnostic binary serialization format developed by Google. One of the key features of Protocol Buffers is that they can be used accross various programming languages.

*Key Features*:
- Efficiency
- Speed
- Cross-Platform Compatibility

*Use Cases*:
- Microservices Communication
- APIs
- Data Storage
- Game Development

*Advantages of Using Protocol Buffers*
- Backward and Forward Compatibility
- Strongly Typed
- Support for Multiple Languages

After creating the dot proto file, we use the Protoc compiler to generate source code in various programming. This generated source code includes classes or structures that can be used to serialize and deserialize our data. Once you have the generated code, you can use it to easily convert your data objects to and from the Protocol Buffers binary format, allowing for efficient data transmission or storage. 


## Syntax and Structure of `.proto` files

A .proto file is a text file that defines the structures of your data in a clear and concise way. It uses a specific syntax to describe the messages and the fields that compose them. The .proto file serves as a blueprint for generating the corresponding code in your preferred programming language. 

Basic structure of a .proto file:

```go
syntax = "proto3";  // Syntax Version
package example;

// Message definition
message Person{
    string name = 1;
    int32 id = 2;
    string email = 3;
}
```

*Defining Fields*
- `<fiedl_type> <field_name> = <field_number>;`

*Basic Field Types*
- int32, int64: Signed integers of varying sizes.
- uint32, uint64: Unsigned integers
- float, double: Floating point numbers
- bool: Boolean values
- string: A sequence of characters
- bytes: A sequence of raw bytes


*Enumerations*
```bash
enum Gender{
    MALE=0;
    FEMALE=1;
    OTHER=2;
}
```

*Nested Messages*
```bash
message Address{
    string street = 1;
    string city = 2;
}
```

```bash
message Person{
    string name=1;
    Address address = 2;    // Nested message
}
```

*Field Options* : Repeated Fields, Required and Optional 
```bash
message Person {
    repeated string phone_numbers = 1;  // List of phone numbers
}
```

*Comments*

In summary, .proto files serve as the foundation structure for defining your data in Protocol Buffers. Understanding the syntax and structure is essential for effectively creating messages and ensuring proper data serialization.
