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

//func main() {
//	fmt.Println(IsFirstPlayerWinner("5H 5C 6S 7S KD 2C 3S 8S 8D TD"))
//	fmt.Println(IsFirstPlayerWinner("5D 8C 9S JS AC 2C 5C 7D 8S QH"))
//	fmt.Println(IsFirstPlayerWinner("2D 9C AS AH AC 3D 6D 7D TD QD"))
//	fmt.Println(IsFirstPlayerWinner("4D 6S 9H QH QC 3D 6D 7H QD QS"))
//	fmt.Println(IsFirstPlayerWinner("2H 2D 4C 4D 4S 3C 3D 3S 9S 9D"))
//}

func createChecker(input chan string, output chan bool) {
	go func() {
		for {
			output <- IsFirstPlayerWinner(<- input)
		}
	}()
}
