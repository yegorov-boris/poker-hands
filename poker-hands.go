package main

import (
	"log"
	"strings"
	"fmt"
	"net/http"
	"bufio"
)

const maxChunkSize = 10
const url = "https://projecteuler.net/project/resources/p054_poker.txt"
const cardValues = "A K Q J T 9 8 7 6 5 4 3 2"
const suits = "D C H S"
const separator = " "

type Card struct {Suit, Value string}
type Hand [5]Card

func main() {
	resp, errGet := http.Get(url)
	if errGet != nil {
		log.Fatalf("Failed to download the combinations: %s\n", errGet)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to download the combinations. Status %d\n", resp.StatusCode);
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
//	fmt.Println(isFirstPlayerWinner("5H 5C 6S 7S KD 2C 3S 8S 8D TD"))
//	fmt.Println(isFirstPlayerWinner("5D 8C 9S JS AC 2C 5C 7D 8S QH"))
//	fmt.Println(isFirstPlayerWinner("2D 9C AS AH AC 3D 6D 7D TD QD"))
//	fmt.Println(isFirstPlayerWinner("4D 6S 9H QH QC 3D 6D 7H QD QS"))
//	fmt.Println(isFirstPlayerWinner("2H 2D 4C 4D 4S 3C 3D 3S 9S 9D"))
//}

func createChecker(input chan string, output chan bool) {
	go func() {
		for {
			output <- isFirstPlayerWinner(<- input)
		}
	}()
}

func isFirstPlayerWinner(hands string) bool {
	first, second := parseHands(hands)
	firstCombinationRank, firstReordered := getCombination(first)
	secondCombinationRank, secondReordered := getCombination(second)

	if firstCombinationRank != secondCombinationRank {
		return firstCombinationRank < secondCombinationRank
	}

	for i, card := range firstReordered {
		firstIndex := strings.Index(cardValues, card.Value)
		secondIndex := strings.Index(cardValues, secondReordered[i].Value)

		if firstIndex != secondIndex {
			return firstIndex < secondIndex
		}
	}

	return false
}

func getCombination(hand Hand) (int, Hand) {
	combinationCheckers := []func(Hand) (bool, Hand) {
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
		if isMatching, reorderedHand := checker(hand); isMatching {
			return i, reorderedHand
		}
	}

	return len(combinationCheckers) - 1, hand
}

func isRoyalFlush(hand Hand) (bool, Hand) {

	if result, _ := isFlush(hand); !result {
		return false, hand
	}

	values := [5]string {"A", "K", "Q", "J", "T"}

	for i, value := range values {
		if hand[i].Value != value {
			return false, hand
		}
	}

	return true, hand
}

func isStraightFlush(hand Hand) (bool, Hand) {
	flush, _ := isFlush(hand)
	straight, _ := isStraight(hand)

	return flush && straight, hand
}

func isFourKind(hand Hand) (bool, Hand) {
	if hand[0].Value == hand[3].Value {
		return true, hand
	}
	if hand[1].Value == hand[4].Value {
		var reorderedHand Hand
		for i, card := range hand[1:] {
			reorderedHand[i] = card
		}
		reorderedHand[4] = hand[0]

		return true, reorderedHand
	}

	return false, hand
}

func isFullHouse(hand Hand) (bool, Hand) {
	if (hand[0].Value == hand[2].Value) && (hand[3].Value == hand[4].Value) {
		return true, hand
	}
	if (hand[0].Value == hand[1].Value) && (hand[2].Value == hand[4].Value) {
		return true, Hand {hand[2], hand[3], hand[4], hand[0], hand[1]}
	}

	return false, hand
}

func isFlush(hand Hand) (bool, Hand) {
	for _, card := range hand {
		if card.Suit != hand[0].Suit {
			return false, hand
		}
	}

	return true, hand
}

func isStraight(hand Hand) (bool, Hand) {
	var values [5]string

	for i, card := range hand {
		values[i] = card.Value
	}

	return strings.Contains(cardValues, strings.Join(values[:], separator)), hand
}

func isThreeKind(hand Hand) (bool, Hand) {
	if hand[0].Value == hand[2].Value {
		return true, hand
	}
	if hand[1].Value == hand[3].Value {
		return true, Hand {hand[1], hand[2], hand[3], hand[0], hand[4]}
	}
	if hand[2].Value == hand[4].Value {
		return true, Hand {hand[2], hand[3], hand[4], hand[0], hand[1]}
	}

	return false, hand
}

func isTwoPairs(hand Hand) (bool, Hand) {
	if (hand[0].Value == hand[1].Value) && (hand[2].Value == hand[3].Value) {
		return true, hand
	}
	if (hand[1].Value == hand[2].Value) && (hand[3].Value == hand[4].Value) {
		return true, Hand {hand[1], hand[2], hand[3], hand[4], hand[0]}
	}
	if (hand[0].Value == hand[1].Value) && (hand[3].Value == hand[4].Value) {
		return true, Hand {hand[0], hand[1], hand[3], hand[4], hand[2]}
	}

	return false, hand
}

func isOnePair(hand Hand) (bool, Hand) {
	index := 0
	for ; index < 4; index++ {
		if hand[index].Value == hand[index + 1].Value {
			break
		}
	}

	if index == 4 {
		return false, hand
	}

	var reorderedHand Hand
	reorderedHand[0] = hand[index]
	reorderedHand[1] = hand[index + 1]
	for i, card := range hand[:index] {
		reorderedHand[i + 2] = card
	}
	for i, card := range hand[index + 2:] {
		reorderedHand[i + index + 2] = card
	}

	return true, reorderedHand
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
