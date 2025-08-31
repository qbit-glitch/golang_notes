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


## Packages in Protocol Buffers

- Packages
- Package Naming Conventions
    - Lowercase
    - Dot Notation
    - Consistency

- Importing Packages

File: `person.proto`
```proto
syntax = "proto3"
package example;
// Message definition
message Person{
    string name = 1;
    int32 id = 2;
}
```

File: `main.proto`
```proto
syntax = "proto3"
package main;
// Importing another .proto file
import "example/person.proto";
message Company{
    repeated example.Person employees = 1; // using the person message from the example package
}
```

When you generate code from your .proto files, the package declaration influences the namespace of the generated code. The code gets generated using the protoc compiler. For instance, if you define a package as example, the generated classes will be organized under that namespace in the target programming language. 


- Avoiding Naming Conflicts

File: `user.proto`
```proto
syntax= "proto3"
package user;
message User {
    string username = 1;
}
```

File: `admin.proto`
```proto
syntax = "proto3";
package admin;
message User {
    string adminId = 1;
}
```

Using packages is essential to prevent naming conflicts in large code bases. For instance if you have multiple messages with the same name in different .proto files, using ensures that they can coexist without issues. In this case, both user messages can exist because they are in different packages.

In summary, packages in Protocol buffers are crucila for organizing your data structures, preventing naming conflicts and maintaining a clean code base. By following best practices in package naming and usagem you can create well structured and maintainable Protocol Buffers projects.


## Messages in Protocol Buffers

Messages are the code data structures used in Protocol buffers to represent and serialize structured data. Understanding how to define and use messages is fundamental to effectively utilizing protocol buffers in our applications. 

A Message in Protocol Buffer is a logical container for structured data. It allows you to define a set of fields, each with a specific type and purpose. Messages can represent complex data structures and are used to facilitate communication between systems. To define a message in a .proto file, use the `message` keyword followed by the message name and a block containing its fields. 

**Messages**
```proto
syntax = "proto3"
package example;

// Defining a message
message Person{
    string name = 1;    // Field 1
    int32 id = 2;       // Field 2
    string email = 3;   // Field 3
}
```

*A message can have the following components :*
- Field Declarations
- Nested Messages
- Enumerations

*Message Options*
```proto
message OldPerson{
    option deprecated = true;   // This message is deprecated
    string name = 1;
}
```

*Best Practices for Messages*
- use meaningful names
- keep messages focused
- plan for evolution
