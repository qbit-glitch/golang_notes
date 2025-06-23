# Endpoints

An endpoint is a specific URL or URI where an API interacts with clients. It represents a specific resource or collection of resources exposed by an API. Endpoints define where an how the API can be accessed and what operations can be performed on the resources.

### Components of an API Endpoint
- Base URL - `https://api.example.com/v1/`
- Path - `/users`, `/orders`, `/products`
- Query Parameters - `?status=active&limit=10`
- HTTP Method - `GET /users`, `POST /orders`

### Types of Endpoints
- **Resource Endpoints**: represents a single resource or a collection of resources. They are used to retrieve, create, update or delete resources.
    - Single Resource: `/users/123`
    - Collection: `/users`
- **Action-Based Endpoints**: perform specific actions or operations that are not necessarily related to CRUD operations.
    - `/users/123/activate`
    - `/orders/checkout`
- **Query-Based Endpoints**: uses query parameters to filter or modify the data returned.
    - `/products?category=electronics`
    - `/orders?status=shipped&limit=10`

### Designing API Endpoints principles:
- Resource Naming
- Consitent and Predictabel URLs
- Versioning
- Error Handling

#### Best Practices
- Use RESTful Principles
- Ensure Security
- Optimize Performance
- Document Endpoints

In summary, API endpoints are crucial for defining how clients interact with an API. They represent specific resources or actions and use URLs, HTTP methods and query parameters to facilitate communication. Proper endpoint design and adherence to best practices ensure a reliable and efficient API.