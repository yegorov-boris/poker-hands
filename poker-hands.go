package main

import (
	"fmt"
	"net/http"
	"os"
	"bufio"
)

func main() {
	const maxChunkSize = 10

	resp, errGet := http.Get("https://projecteuler.net/project/resources/p054_poker.txt")
	if errGet != nil {
		fmt.Printf("Failed to download the combinations: %s\n", errGet)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to download the combinations. Status %d\n", resp.StatusCode);
		os.Exit(1)
	}

	defer resp.Body.Close()

	var inputs [maxChunkSize]chan string
	var outputs [maxChunkSize]chan bool

	for i := range inputs {
		inputs[i] = make(chan string)
		outputs[i] = make(chan bool)
		createChecker(inputs[i], outputs[i]);
	}

	scanner := bufio.NewScanner(resp.Body)
	var notEof = true

	for notEof {
		var currentChunkSize = 0;

		for i := range inputs {
			if !scanner.Scan() {
				notEof = false
				break
			}

			inputs[i] <- scanner.Text()
			currentChunkSize++
		}

		for i := 0; i < currentChunkSize; i++ {
			fmt.Println(<- outputs[i])
		}
	}
}

func createChecker(input chan string, output chan bool) {
	go func() {
		for {
			msg := <- input
			if msg == "foo" {
				output <- true
			} else {
				output <- false
			}
		}
	}()
}
