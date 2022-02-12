package Cmd

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Timeout waits for n seconds or keypress
func Timeout(seconds int) {
	input := make(chan string, 1)

	go timeoutGetInput(input)

	for seconds > 0 {

		fmt.Printf("Waiting for %d seconds, press a key to continue ...\r", seconds)

		select { // checks for input or times out
		case <-input:
			seconds = 0
		case <-time.After(time.Second):
			seconds--
		}
	}
	fmt.Println()
}

// timeoutGetInput reads from os.Stdin and sends it to the channel input
func timeoutGetInput(input chan string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		in, err := reader.ReadString('\n')
		if err != nil {
			in = ""
		}

		input <- in
	}
}
