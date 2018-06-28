package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	//Commandline Arguments and their default values...
	url := flag.String("url", "http://status.pfmtest.com", "URL to check..")
	respText := flag.String("respText", "", "Text to check for in response..")
	respCode := flag.String("respCode", "200", "Response code to check for..")
	verbose := flag.Bool("v", false, "Verbose output..")

	flag.Parse()

	c := &http.Client{
		Timeout: 15 * time.Second,
	}
	response, err := c.Get(*url)
	if err != nil {
		fmt.Println("fail")
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if *verbose {
			fmt.Printf("Response: %s", contents)
			fmt.Println(response.Status)
		}
		if err != nil {
			fmt.Println("fail")
			os.Exit(1)
		}
		if strings.Contains(response.Status, *respCode) && strings.Contains(string(contents), *respText) {
			fmt.Println("pass")
		} else {
			fmt.Println("fail")
			os.Exit(1)
		}
	}
	os.Exit(0)
}
