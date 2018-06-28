package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	//Commandline Arguments and their default values...
	dst := flag.String("dst", "127.0.0.1:8080", "URL to check..")
	prot := flag.String("prot", "tcp", "Text to check for in response..")
	sendTxt := flag.String("sendTxt", "Hello", "Text to send on socket..")
	respTxt := flag.String("respTxt", "", "Response txt to check for..")
	verbose := flag.Bool("v", false, "Verbose output..")

	flag.Parse()

	//Set timeout for socket

	timeoutDuration := 10 * time.Second
	d := net.Dialer{
		Timeout: 10 * time.Second,
	}
	// connect to this socket
	conn, err := d.Dial(*prot, *dst)
	if err == nil {
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))
		// send to socket
		fmt.Fprintf(conn, *sendTxt+"\n")
		// listen for reply
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err == nil {
			if len(*respTxt) > 1 {
				if strings.Contains(message, *respTxt) && len(message) > 1 && *verbose {
					fmt.Print("Message from server: " + message)
				}
			}
			fmt.Println("pass")
			os.Exit(0)
		} else {
			if *verbose {
				fmt.Println(err)
			}
		}
	} else if *verbose {
		fmt.Println(err)
	}
	fmt.Println("fail")
	os.Exit(1)
}
