package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("\nURL Parsing ")

	// [scheme://][username@]host[:port][/path][?query][#fragment]

	rawURL := "https://example.com:8080/path?query=param#fragment"

	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("error parsing the url: ", err)
		return
	}

	fmt.Println(parsedUrl.Scheme)
	fmt.Println(parsedUrl.Host)
	fmt.Println(parsedUrl.Port())
	fmt.Println(parsedUrl.Path)
	fmt.Println(parsedUrl.Query())
	fmt.Println(parsedUrl.Fragment)

	fmt.Println("-----------------")

	rawURL2 := "https://example.com:8080/path?name=John&age=20"
	parsedUrl2, err := url.Parse(rawURL2)
	if err != nil {
		fmt.Println("error parsing url:", err)
		return
	}

	queryParams := parsedUrl2.Query()
	fmt.Println(queryParams)

	fmt.Println(queryParams.Get("name"))
	fmt.Println(queryParams.Get("age"))

	queryParams.Add("school","MIT")
	fmt.Println(queryParams, queryParams.Get("school"))

	// Building the URL

	baseUrl := &url.URL{
		Scheme: "https",
		Host: "example.com",
		Path: "/path",
	}

	query := baseUrl.Query()
	// query.Set(key, value)
	query.Set("name", "Smith")
	query.Set("age", "35")

	fmt.Println(query)

	baseUrl.RawQuery = query.Encode()
	fmt.Println("Built URL:", baseUrl)

	fmt.Println("-----------------")
	values := &url.Values{}
	// Add key value pairs to the values object
	values.Add("name", "Alice")
	values.Add("city", "Wonderland")
	values.Add("age", "30")
	fmt.Println(values)
	// Encode the query
	encodedQuery := values.Encode()
	fmt.Println(encodedQuery)

	// build a URL
	baseUrl1 := "https://example.com/search"
	fullUrl := baseUrl1 + "?" + encodedQuery
	fmt.Println(fullUrl)

}