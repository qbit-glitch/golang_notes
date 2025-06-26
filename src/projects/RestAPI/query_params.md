# Query Params

Query paramters are a common way to pass data to a server via URL. They are different from path parameters because path parameter is a single value. Query parameters are typically used in GET requests to send data to the server such as filters or search criteria, or any other data that doesn't need to be in the request body in a URL.

We can provide default values for some of the query parameters which are not received by us.

Confidential information cannot be passes through URL parameters, query and path. These are both URL parameters. We cannot use confidential information in the URL. So that's why we will use username and password through JSON format.

The query parameters are used with GET requests most of the time. Query parameters are a powerful way to pass data to our server via URLs and the standard library in Go provides simple and effective tools to extract and work with query parameters. Used to implement features like filtering, sorting, pagination, etc.


***Add this code to the GET method:***

```go
// teachers/?key=value&query=value2&sortby=email&sortorder=ASC

queryParams := r.URL.Query()
fmt.Println("Query Params:", queryParams)
sortby := queryParams.Get("sortby")
key := queryParams.Get("key")
sortorder := queryParams.Get("sortorder")
if sortorder == ""{
    sortorder = "DESC"
}
fmt.Printf("Sortby: %v | Sort Order: %v | Key: %v", sortby, sortorder, key)
```
