package main

import (
	"fmt"
	"log"
	"strings"
)

//const maxChunkSize = 10
//const url = "https://projecteuler.net/project/resources/p054_poker.txt"
const cardValues = "2 3 4 5 6 7 8 9 T J Q K A"
var suits = "D C H S"
//var cardValues = [13]string {"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
//var suits = [4]string {"D", "C", "H", "S"}

type Card struct {Suit, Value string}

//func main() {
//	resp, errGet := http.Get(url)
//	if errGet != nil {
//		log.Fatalf("Failed to download the combinations: %s\n", errGet)
//	}
//	if resp.StatusCode != http.StatusOK {
//		log.Fatalf("Failed to download the combinations. Status %d\n", resp.StatusCode);
//	}
//
//	defer resp.Body.Close()
//
//	var inputs [maxChunkSize]chan string
//	var outputs [maxChunkSize]chan bool
//
//	for i := range inputs {
//		inputs[i] = make(chan string)
//		outputs[i] = make(chan bool)
//		createChecker(inputs[i], outputs[i]);
//	}
//
//	scanner := bufio.NewScanner(resp.Body)
//	notEof := true
//	firstPlayerWinsCount := 0;
//
//	for notEof {
//		currentChunkSize := 0;
//
//		for i := range inputs {
//			if !scanner.Scan() {
//				notEof = false
//				break
//			}
//
//			inputs[i] <- scanner.Text()
//			currentChunkSize++
//		}
//
//		for i := 0; i < currentChunkSize; i++ {
//			if firstPlayerWon := <- outputs[i]; firstPlayerWon {
//				firstPlayerWinsCount++
//			}
//		}
//	}
//
//	fmt.Printf("The first player won %d times\n", firstPlayerWinsCount)
//}

func main() {
	isFirstPlayerWinner("5H 5C 6S 7S KD 2C 3S 8S 8D TD")
}

//func createChecker(input chan string, output chan bool) {
//	go func() {
//		for {
//			hands := <- input
//			output <- (hands == "foo")
//		}
//	}()
//}

func isFirstPlayerWinner(hands string) bool {
	first, second := parseHands(hands)

	fmt.Print(first)
	fmt.Print(second)

	return false
}

func parseHands(hands string) ([5]Card, [5]Card) {
	var first, second [5]Card
	cardStrings := strings.Split(hands, " ")

	if len(cardStrings) != 10 {
		log.Fatal("Failed to parse a line with hands: wrong length!/n")
	}

	for i, cardString := range cardStrings[:5] {
		first[i] = parseCardString(cardString)
	}
	for i, cardString := range cardStrings[5:] {
		second[i] = parseCardString(cardString)
	}

	return first, second
}

func parseCardString(cardString string) Card {
	length := len(cardString)

	if length != 2 {
		log.Fatal("Failed to parse an encoded card: wrong length!/n")
	}

	suit := cardString[1:]
	value := cardString[:1]

	if !strings.Contains(suits, suit) {
		log.Fatal("Failed to parse an encoded card: wrong suit!/n")
	}
	if !strings.Contains(cardValues, value) {
		log.Fatal("Failed to parse an encoded card: wrong card value!/n")
	}

	return Card{Suit: suit, Value: value}
}
