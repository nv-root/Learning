package main

import "fmt"
import "net/url"

func main() {

	rawURL := "https://example.com:8080/path?query=param&query=param#fragment"
	parsedURL, err := url.Parse(rawURL)

	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println(parsedURL)
	fmt.Println(parsedURL.Scheme)
	fmt.Println(parsedURL.Host)
	fmt.Println(parsedURL.Port())
	fmt.Println(parsedURL.Path)
	fmt.Println(parsedURL.RawQuery)
	fmt.Println(parsedURL.Fragment)

	rawURL = "https://example.com/path?name=john&age=30"
	parsedURL2, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	queryParams := parsedURL2.Query()
	fmt.Println(queryParams) // stored as a map
	fmt.Println("Name:", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))

	// building url
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := baseURL.Query()
	query.Set("name", "justin")
	query.Set("age", "30")
	baseURL.RawQuery = query.Encode()

	fmt.Println("built url:", baseURL.String())

	values := url.Values{}

	values.Add("name", "ariana")
	values.Add("age", "32")

	fmt.Println(values)

	encodedQuery := values.Encode()
	
	fmt.Println("fullURL:", "https://example.com/user"+"?"+encodedQuery)

	fmt.Println(encodedQuery)

}
