package main

import (
	"bufio"
	"fmt"
	browser "github.com/EDDYCJY/fake-useragent"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func getHTTPRequest(httpLink string) {
	random := browser.Random()
	log.Printf("Random: %s", random)

	// get request with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Create and modify HTTP request before sending
	request, err := http.NewRequest("Get", httpLink, nil)
	// Print error (if any)
	if err != nil {
		log.Fatal(err, "request fail")
	}
	request.Header.Set("User-Agent", random)

	// Make HTTP GET request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err, "response fail")
	}
	defer response.Body.Close()

	// Copy data from the response to standard output
	_, err = io.Copy(os.Stdout, response.Body) // Will show HTML code
	// Print error (if any)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input a link:")
	link, _, err := reader.ReadLine()
	// Print error (if any)
	if err != nil {
		log.Fatal(err)
	} else {
		getHTTPRequest(string(link))
	}
}
