package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)

type Card struct {Suit, Value string}
type Hand [5]Card

func main() {
	resp, errGet := http.Get(Url)
	if errGet != nil {
		log.Fatalf("Failed to download the combinations: %s\n", errGet)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to download the combinations. Status %d\n", resp.StatusCode);
	}

	defer resp.Body.Close()

	var inputs [MaxChunkSize]chan string
	var outputs [MaxChunkSize]chan bool

	for i := range inputs {
		inputs[i] = make(chan string)
		outputs[i] = make(chan bool)
		createChecker(inputs[i], outputs[i]);
	}

	scanner := bufio.NewScanner(resp.Body)
	notEof := true
	firstPlayerWinsCount := 0;

	for notEof {
		currentChunkSize := 0;

		for i := range inputs {
			if !scanner.Scan() {
				notEof = false
				break
			}

			inputs[i] <- scanner.Text()
			currentChunkSize++
		}

		for i := 0; i < currentChunkSize; i++ {
			if <- outputs[i] {
				firstPlayerWinsCount++
			}
		}
	}

	fmt.Printf("The first player won %d times\n", firstPlayerWinsCount)
}

func createChecker(input chan string, output chan bool) {
	go func() {
		for {
			result, err := IsFirstPlayerWinner(<- input)
			if err != nil {
				log.Fatal(err)
			}

			output <- result
		}
	}()
}
