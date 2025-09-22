// 2. Web Scraping
// Example: Fetching Data from a Website

// You can automate the process of gathering data from websites using HTTP requests.

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("https://webscraper.io/test-sites/e-commerce/allinone/computers")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
