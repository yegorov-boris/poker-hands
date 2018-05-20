package main

import (
	"errors"
	"fmt"
	"strings"
)

type cardStringParser interface {
	SplitHandsString(string) ([]string, error)
	ParseCardString(string) (Card, error)
	SortByValue(Hand) Hand
}

type CardStringParser struct {
	Suits      string
	CardValues string
	Separator  string
}

func (p CardStringParser) SplitHandsString(hands string) ([]string, error) {
	cardStrings := strings.Split(hands, p.Separator)
	if len(cardStrings) != 10 {
		return cardStrings, errors.New("failed to parse a line with hands: wrong length")
	}

	return cardStrings, nil
}

func (p CardStringParser) ParseCardString(cardString string) (Card, error) {
	length := len(cardString)

	if length != 2 {
		return Card{}, errors.New("failed to parse an encoded card: wrong length")
	}

	suit := cardString[1:]
	value := cardString[:1]

	if !strings.Contains(p.Suits, suit) {
		return Card{}, errors.New("failed to parse an encoded card: wrong suit")
	}
	if !strings.Contains(p.CardValues, value) {
		return Card{}, errors.New("failed to parse an encoded card: wrong card value")
	}

	return Card{Suit: suit, Value: value}, nil
}

func (p CardStringParser) SortByValue(hand Hand) Hand {
	sortedCardsCount := 0
	for _, value := range strings.Split(p.CardValues, p.Separator) {
		i := sortedCardsCount
		for j, card := range hand[i:] {
			if card.Value == value {
				currentCardIndex := j + i
				hand[currentCardIndex] = hand[sortedCardsCount]
				hand[sortedCardsCount] = card
				sortedCardsCount++
			}
		}
		if sortedCardsCount == 5 {
			break
		}
	}
	return hand
}

func ParseHands(parser cardStringParser, hands string) (Hand, Hand, error) {
	var first, second Hand
	cardStrings, err := parser.SplitHandsString(hands)
	if err != nil {
		return Hand{}, Hand{}, err
	}

	for i, cardString := range cardStrings {
		if strings.Count(hands, cardString) > 1 {
			return Hand{}, Hand{}, fmt.Errorf("failed to parse a line with hands: %s is not unique", cardString)
		}

		card, err := parser.ParseCardString(cardString)
		if err != nil {
			return Hand{}, Hand{}, err
		}

		if i < 5 {
			first[i] = card
		} else {
			second[i%5] = card
		}

	}
	return parser.SortByValue(first), parser.SortByValue(second), nil
}
