package main

import (
	"bufio"
	"fmt"
	browser "github.com/EDDYCJY/fake-useragent"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func getHTTPRequest(httpLink string) {
	random := browser.Random()

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

func downloader(httpLink string) {
	random := browser.Random()

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

	// Create output file
	outFile, err := os.Create("output.html")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Copy data from HTTP response to file
	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func option() {
	fmt.Println("What do you want to do ?")
	fmt.Println("[1] See website source code")
	fmt.Println("[2] Download a URL")
	fmt.Println("Please input the corresponding number of the action you would like to do")
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("reset") //Linux example, its tested
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	for true {
		reader := bufio.NewReader(os.Stdin)
		option()
		input, _, err := reader.ReadLine()
		// Print error (if any)
		if err != nil {
			log.Fatal(err)
		} else if string(input) == "1" {
			fmt.Println("Please input the link or type 'b' to go back to option")
			read, _, _ := reader.ReadLine()
			if string(read) == "b" || string(read) == "B" {
				CallClear()
			} else {
				getHTTPRequest(string(read))
			}
		} else if string(input) == "2" {
			fmt.Println("Please input the link or type 'b' to go back to option")
			read, _, _ := reader.ReadLine()
			if string(read) == "b" || string(read) == "B" {
				CallClear()
			} else {
				downloader(string(read))
			}
		}
	}
}
