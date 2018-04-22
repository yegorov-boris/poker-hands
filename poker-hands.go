package main

import (
	"fmt"
	"log"
	"strings"
)

//const maxChunkSize = 10
//const url = "https://projecteuler.net/project/resources/p054_poker.txt"
const cardValues = "A K Q J T 9 8 7 6 5 4 3 2"
const suits = "D C H S"
const separator = " "
//var cardValues = [13]string {"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
//var suits = [4]string {"D", "C", "H", "S"}

type Card struct {Suit, Value string}
type Hand [5]Card

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
	isFirstPlayerWinner("JH QH TH AH KH 3D 6D 7D TD QD")
}

//func createChecker(input chan string, output chan bool) {
//	go func() {
//		for {
//			hands := <- input
//			output <- isFirstPlayerWinner(hands)
//		}
//	}()
//}

func isFirstPlayerWinner(hands string) bool {
	first, second := parseHands(hands)
	firstCombination := getCombination(first)
	secondCombination := getCombination(second)
	fmt.Println(firstCombination)
	fmt.Println(secondCombination)

	return firstCombination < secondCombination
}

func getCombination(hand Hand) int {
	combinationCheckers := []func(Hand) bool {
		isRoyalFlush,
		isStraightFlush,
		isFourKind,
		isFullHouse,
		isFlush,
		isStraight,
		isThreeKind,
		isTwoPairs,
		isOnePair,
	}

	for i, checker := range combinationCheckers {
		if checker(hand) {
			return i
		}
	}

	return len(combinationCheckers) - 1
}

func isRoyalFlush(hand Hand) bool {
	if !isFlush(hand) {
		return false
	}

	values := [5]string {"A", "K", "Q", "J", "T"}

	for i, value := range values {
		if hand[i].Value != value {
			return false
		}
	}

	return true
}

func isStraightFlush(hand Hand) bool {
	return isFlush(hand) && isStraight(hand)
}

func isFourKind(hand Hand) bool {
	return (hand[0].Value == hand[3].Value) || (hand[1].Value == hand[4].Value)
}

func isFullHouse(hand Hand) bool {
	twoAndThree := (hand[0].Value == hand[1].Value) && (hand[2].Value == hand[4].Value)
	threeAndTwo := (hand[0].Value == hand[2].Value) && (hand[3].Value == hand[4].Value)
	return twoAndThree || threeAndTwo
}

func isFlush(hand Hand) bool {
	for _, card := range hand {
		if card.Suit != hand[0].Suit {
			return false
		}
	}

	return true
}

func isStraight(hand Hand) bool {
	var values [5]string

	for i, card := range hand {
		values[i] = card.Value
	}

	return strings.Contains(cardValues, strings.Join(values[:], separator))
}

func isThreeKind(hand Hand) bool {
	left := hand[0].Value == hand[2].Value
	middle := hand[1].Value == hand[3].Value
	right := hand[2].Value == hand[4].Value

	return left || middle || right
}

func isTwoPairs(hand Hand) bool {
	left := (hand[1].Value == hand[2].Value) && (hand[3].Value == hand[4].Value)
	middle := (hand[0].Value == hand[1].Value) && (hand[3].Value == hand[4].Value)
	right := (hand[0].Value == hand[1].Value) && (hand[2].Value == hand[3].Value)

	return left || middle || right
}

func isOnePair(hand Hand) bool {
	for i := 0; i < 4; i++ {
		if hand[i].Value == hand[i + 1].Value {
			return true
		}
	}

	return false
}

func parseHands(hands string) (Hand, Hand) {
	var first, second Hand
	cardStrings := strings.Split(hands, separator)

	if len(cardStrings) != 10 {
		log.Fatal("Failed to parse a line with hands: wrong length!/n")
	}

	for i, cardString := range cardStrings[:5] {
		first[i] = parseCardString(cardString)
	}
	for i, cardString := range cardStrings[5:] {
		second[i] = parseCardString(cardString)
	}

	return sortByValue(first), sortByValue(second)
}

func sortByValue(hand Hand) Hand {
	var sortedHand Hand
	i := 0

	for _, value := range strings.Split(cardValues, separator) {
		for _, card := range hand {
			if card.Value == value {
				sortedHand[i] = card
				i++
			}
		}

		if i == 5 {
			break
		}
	}

	return sortedHand
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
