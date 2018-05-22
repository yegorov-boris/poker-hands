package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

type Card struct{ Suit, Value string }
type Hand [5]Card
type EitherBool struct {
	Left  bool
	Right error
}
type Scanner interface {
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

	firstPlayerWinsCount, errCount := CountWins(bufio.NewScanner(resp.Body), MaxChunkSize)
	if errCount != nil {
		log.Fatal(errCount)
	}

	fmt.Printf("The first player won %d times\n", firstPlayerWinsCount)
}

func CreateCheckers(chunkSize int) ([]chan string, []chan EitherBool) {
	var inputs []chan string
	var outputs []chan EitherBool
	defaultComparator := comparator{
		config: defaultConfig(),
		parser: handsParser{
			splitter:   splitter{config: defaultConfig()},
			cardParser: cardParser{config: defaultConfig()},
			sorter:     sorter{config: defaultConfig()},
		},
		matcher: combMatcher{config: defaultConfig()},
	}
	for i := 0; i < chunkSize; i++ {
		input := make(chan string)
		output := make(chan EitherBool)
		inputs = append(inputs, input)
		outputs = append(outputs, output)
		go func() {
			for {
				result, err := defaultComparator.IsFirstPlayerWinner(<-input)
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

func ScanChunk(scanner Scanner, size int) (bool, []string) {
	var chunk []string
	for i := 0; i < size; i++ {
		if !scanner.Scan() {
			return true, chunk
		}

		chunk = append(chunk, scanner.Text())
	}

	return false, chunk
}

func CountWins(scanner Scanner, chunkSize int) (int, error) {
	firstPlayerWinsCount := 0
	inputs, outputs := CreateCheckers(chunkSize)

	for {
		eof, chunk := ScanChunk(scanner, chunkSize)

		for i, cardString := range chunk {
			inputs[i] <- cardString
		}

		for i := range chunk {
			result := <-outputs[i]
			if result.Right != nil {
				return 0, result.Right
			}
			if result.Left {
				firstPlayerWinsCount++
			}
		}

		if eof {
			return firstPlayerWinsCount, nil
		}
	}
}
