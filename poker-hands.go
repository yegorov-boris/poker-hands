package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

type Card struct{ Suit, Value string }
type Hand [5]Card
type Either struct {
	Left  chan string
	Right error
}
type scanner interface {
	Scan() bool
	Text() string
}

func main() {
	resp, errGet := http.Get(Url)
	if errGet != nil {
		log.Fatalf("Failed to download the combinations: %s\n", errGet)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to download the combinations. Status %d\n", resp.StatusCode)
	}

	defer resp.Body.Close()

	firstPlayerWinsCount, errCount := countWins(bufio.NewScanner(resp.Body), defaultComparator(), poolSzie)
	if errCount != nil {
		log.Fatal(errCount)
	}

	fmt.Printf("The first player won %d times\n", firstPlayerWinsCount)
}

func counter(comparator comparator, requestsToRead chan Either, counts chan int) {
	count := 0
	reader := make(chan string)
	requestsToRead <- Either{Left: reader, Right: nil}

	for handsString := range reader {
		if handsString == stop {
			close(reader)
			counts <- count
			return
		} else {
			result, err := comparator.IsFirstPlayerWinner(handsString)
			if result {
				count++
			}

			requestsToRead <- Either{Left: reader, Right: err}
		}
	}
}

func countWins(scanner scanner, comparator comparator, poolSize int) (int, error) {
	firstPlayerWinsCount := 0
	requestsToRead := make(chan Either, poolSize)
	counts := make(chan int, poolSize)

	for i := 0; i < poolSize; i++ {
		go counter(comparator, requestsToRead, counts)
	}

	activeCounters := poolSize
	for requestToRead := range requestsToRead {
		if requestToRead.Right != nil {
			return firstPlayerWinsCount, requestToRead.Right
		}

		if scanner.Scan() {
			requestToRead.Left <- scanner.Text()
		} else {
			requestToRead.Left <- stop
			activeCounters--
			if activeCounters == 0 {
				close(requestsToRead)
			}
		}
	}

	for i := 0; i < poolSize; i++ {
		firstPlayerWinsCount = firstPlayerWinsCount + (<-counts)
	}

	return firstPlayerWinsCount, nil
}
