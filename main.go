package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
)

func getHTTP(httpLink string) {
	response, error := http.Get(httpLink)

	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	link, _, _ := reader.ReadLine()
	getHTTP(string(link))
}
