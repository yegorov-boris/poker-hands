package main

import (
	"strings"
	"log"
)

func ParseHands(hands string) (Hand, Hand) {
	var first, second Hand
	cardStrings := strings.Split(hands, Separator)

	if len(cardStrings) != 10 {
		log.Fatal("Failed to parse a line with hands: wrong length!/n")
	}

	for i, cardString := range cardStrings {
		if strings.Count(hands, cardString) > 1 {
			log.Fatalf("Failed to parse a line with hands: %s is not unique!/n", cardString)
		}

		if i < 5 {
			first[i] = ParseCardString(cardString)
		} else {
			second[i % 5] = ParseCardString(cardString)
		}

	}

	return SortByValue(first), SortByValue(second)
}

func SortByValue(hand Hand) Hand {
	var sortedHand Hand
	i := 0

	for _, value := range strings.Split(CardValues, Separator) {
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

func ParseCardString(cardString string) Card {
	length := len(cardString)

	if length != 2 {
		log.Fatal("Failed to parse an encoded card: wrong length!/n")
	}

	suit := cardString[1:]
	value := cardString[:1]

	if !strings.Contains(Suits, suit) {
		log.Fatal("Failed to parse an encoded card: wrong suit!/n")
	}
	if !strings.Contains(CardValues, value) {
		log.Fatal("Failed to parse an encoded card: wrong card value!/n")
	}

	return Card{Suit: suit, Value: value}
}
