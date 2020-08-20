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
	"strings"
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

func downloader(httpLink string) {

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
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
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
		input, err := reader.ReadString('\n')
		// Print error (if any)
		if err != nil {
			log.Fatal(err)
		} else if strings.TrimRight(input, "\n") == "1" {
			fmt.Println("Please input the link or type 'b' to go back to option")
			read, _ := reader.ReadString('\n')
			if strings.TrimRight(read, "\n") == "b" || strings.TrimRight(read, "\n") == "B" {
				CallClear()
			} else {
				getHTTPRequest(read)
			}
		}
	}
}
