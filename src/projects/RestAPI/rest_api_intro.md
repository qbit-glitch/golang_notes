# What is REST API

## Introduction

**API (Application Programming Interface)**:
- Set of rules and protocols that allows different software applications to communicate with each other.
- It's a standardized way for applications to interact and exchange data.

## REST (Representational State Transfer)
REST is an architectural style for designing networked applications. RESTful systems communicate via HTTP and use a stateless client server and cacheable communication protocol.

**Key Components**:
- Statelessness: each request from a client to the server must contain all the information the server needs to fulfill the request. The server does not store any state about the client session between requests.
- Client Server Architecture: Client server architecture focuses on separation of concerns between client and server. The client is responsible for the user interface while the server manages data and business logic.
- Uniform Interface: A consistent interface accross the system that simplifies and decouples the architecture, allowing for independent evolution of components is Uniform Interface.
- Resource Based: REST is resource based which are data objects are identified by URLs. Each resource is accessed using standard http methods like GET, POST, PUT, DELETE.
- Stateless Communication: Every request from a client to server must contain all necessary information example authentication tokens, query parameters, etc. The server does not store any information about previous requests. 
- Cacheability: Responses from the server must define whether they are cacheable or not. Proper caching reduces the need for request and improve performance.

## RESTFUL API

RESTful API is an API that adheres to the principles of REST. It uses standard http methods to perform operations on resources identified by URLs. It provides a way for different applications or services to interact with each other using a uniform interface.

### Key Components
- Resources: Resources are the objects or data that the API exposes. Each resource is identified by a URL.
- Endpoints: Specific URLs where resources are accesses or manipulated are called endpoints. Only the server has the endpoints and not the frontend.
- HTTP Methods: incudes GET, POST, PUT, PATCH and DELETE.
- Request and Response Formats: Client sends request to the server, request include https methods, urls, headers and optionally a body. Next the server sends response to the client. Response includes status code, headers and a body containing the requested data or the result of the operation.

### Benefits of RESTful APIs 
- Scalability: Support scalability by supporting client and server concers and allowing for distributed systems. They are stateless, which simplifies server design and scaling. 
- Interoperability
- Flexibility
- Explanation
- Cacheability

### Limitations
While statelessness simplifies server design, it can increase the complexity of client side state management especially for applications requiring complex interactions. REST APIs can involve additional overhead due to the need for multiple HTTP requests, headers and status codes.

REST APIs are request response based and may not be suitable for real time applications where immediate are required. Alterantives like WebSockets or GraphQL might be used in those cases.

### Summary 
REST APIs are a powerful way to enable communication between web applications and services. They follow the principles of REST such as statelessness, uniform interface and resource based interaction using standard http methods to perform operations on resources.