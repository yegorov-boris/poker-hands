package main

import (
	"strings"
	"log"
	"errors"
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

		card, err := ParseCardString(cardString)
		if err != nil {
			log.Fatal(err)
		}

		if i < 5 {
			first[i] = card
		} else {
			second[i % 5] = card
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

func ParseCardString(cardString string) (Card, error) {
	length := len(cardString)

	if length != 2 {
		return  Card{}, errors.New("failed to parse an encoded card: wrong length")
	}

	suit := cardString[1:]
	value := cardString[:1]

	if !strings.Contains(Suits, suit) {
		return  Card{}, errors.New("failed to parse an encoded card: wrong suit")
	}
	if !strings.Contains(CardValues, value) {
		return  Card{}, errors.New("failed to parse an encoded card: wrong card value")
	}

	return Card{Suit: suit, Value: value}, nil
}
