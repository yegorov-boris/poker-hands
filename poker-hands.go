package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)

type Card struct {Suit, Value string}
type Hand [5]Card
type EitherBool struct {Left bool; Right error}

func main() {
	resp, errGet := http.Get(Url)
	if errGet != nil {
		log.Fatalf("Failed to download the combinations: %s\n", errGet)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to download the combinations. Status %d\n", resp.StatusCode);
	}

	defer resp.Body.Close()

	inputs, outputs := CreateCheckers(MaxChunkSize)

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
			result := <- outputs[i]
			if result.Right != nil {
				log.Fatal(result.Right)
			}
			if result.Left {
				firstPlayerWinsCount++
			}
		}
	}

	fmt.Printf("The first player won %d times\n", firstPlayerWinsCount)
}

func CreateCheckers(chunkSize int) ([]chan string, []chan EitherBool) {
	var inputs []chan string
	var outputs []chan EitherBool
	for i := 0; i < chunkSize; i++ {
		input := make(chan string)
		output := make(chan EitherBool)
		inputs = append(inputs, input)
		outputs = append(outputs, output)
		go func() {
			for {
				result, err := IsFirstPlayerWinner(<- input)
				if err == nil {
					output <- EitherBool{Left: result}
				} else {
					output <- EitherBool{Right: err}
				}
			}
		}()
	}

	return inputs, outputs
}
