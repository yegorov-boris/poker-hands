package main

import (
	"errors"
	"fmt"
	"strings"
)

type handsStringSplitter interface {
	splitHandsString(string) ([]string, error)
}

type sorterByValue interface {
	sortByValue(Hand) Hand
}

type cardStringParser interface {
	parseCardString(string) (Card, error)
}

type handsStringParser interface {
	parseHands(string) (Hand, Hand, error)
}

type splitter struct {
	config config
}

type sorter struct {
	config config
}

type cardParser struct {
	config config
}

type handsParser struct {
	splitter   handsStringSplitter
	cardParser cardStringParser
	sorter     sorterByValue
}

func (s splitter) splitHandsString(hands string) ([]string, error) {
	cardStrings := strings.Split(hands, s.config.separator)
	if len(cardStrings) != 10 {
		return cardStrings, errors.New("failed to parse a line with hands: wrong length")
	}

	return cardStrings, nil
}

func (p cardParser) parseCardString(cardString string) (Card, error) {
	length := len(cardString)

	if length != 2 {
		return Card{}, errors.New("failed to parse an encoded card: wrong length")
	}

	suit := cardString[1:]
	value := cardString[:1]

	if !strings.Contains(p.config.suits, suit) {
		return Card{}, errors.New("failed to parse an encoded card: wrong suit")
	}
	if !strings.Contains(p.config.cardValues, value) {
		return Card{}, errors.New("failed to parse an encoded card: wrong card value")
	}

	return Card{Suit: suit, Value: value}, nil
}

func (s sorter) sortByValue(hand Hand) Hand {
	sortedCardsCount := 0
	for _, value := range strings.Split(s.config.cardValues, s.config.separator) {
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

func (p handsParser) parseHands(hands string) (Hand, Hand, error) {
	var first, second Hand
	cardStrings, err := p.splitter.splitHandsString(hands)
	if err != nil {
		return Hand{}, Hand{}, err
	}

	for i, cardString := range cardStrings {
		if strings.Count(hands, cardString) > 1 {
			return Hand{}, Hand{}, fmt.Errorf("failed to parse a line with hands: %s is not unique", cardString)
		}

		card, err := p.cardParser.parseCardString(cardString)
		if err != nil {
			return Hand{}, Hand{}, err
		}

		if i < 5 {
			first[i] = card
		} else {
			second[i%5] = card
		}

	}
	return p.sorter.sortByValue(first), p.sorter.sortByValue(second), nil
}
