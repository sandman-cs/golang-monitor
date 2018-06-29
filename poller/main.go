package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	ver = "0.1"
)

func main() {

	fmt.Println("Monitoring Poller Ver: ", ver)

	go func() {
		for {
			log.Println("Executing Work Loop..")
			//Do the polling here
			osPollExec("net_check.exe", "")
			log.Println("Work Loop Complete...")
			time.Sleep(conf.PollInterval)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		time.Sleep(time.Second)
		text, _ := reader.ReadString('\n')
		if strings.Contains(string(text), "exit") {
			os.Exit(0)
		}
	}

}

func osPollExec(cmdStr string, args string) bool {
	log.Println("Executing: ", cmdStr, " ", args)
	cmd := exec.Command(cmdStr, args)

	out, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
		return false
	}
	if err := cmd.Start(); err != nil {
		log.Println(err)
		return false
	}

	log.Println("Waitig for command to finish...")

	err = cmd.Wait()

	buf := new(bytes.Buffer)
	buf.ReadFrom(out)
	if err != nil {
		log.Println("Application Error: ", err)
		log.Println(buf.String())
		return false
	}

	fmt.Println("OutPut: ", buf.String())

	return false

}
