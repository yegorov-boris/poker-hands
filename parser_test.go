package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"math/rand"
	"sort"
)

func TestParseCardString(t *testing.T) {
	log.Println("ParseCardString")

	log.Println("should parse a valid card string")
	expected := Card {Suit: "D", Value: "9"}
	actual, _ := ParseCardString("9D")
	assert.Equal(t, expected, actual)

	log.Println("should fail when a card string is longer than 2 bytes")
	_, errTooLong := ParseCardString(RandomString(3, 5))
	assert.Errorf(t, errTooLong, "failed to parse an encoded card: wrong length")

	log.Println("should fail when a card string is shorter than 2 bytes")
	_, errTooShort := ParseCardString(RandomString(0, 1))
	assert.Errorf(t, errTooShort, "failed to parse an encoded card: wrong length")

	log.Println("should fail when a card string has invalid suit")
	_, errInvalidSuit := ParseCardString(strings.Join([]string{ValidCardValue(), InvalidSuit()}, ""))
	assert.Errorf(t, errInvalidSuit, "failed to parse an encoded card: wrong suit")

	log.Println("should fail when a card string has invalid card value")
	_, errInvalidCardValue := ParseCardString(strings.Join([]string{InvalidCardValue(), ValidSuit()}, ""))
	assert.Errorf(t, errInvalidCardValue, "failed to parse an encoded card: wrong card value")
}

func TestSortByValue(t *testing.T) {
	log.Println("SortByValue")

	log.Println("should sort a hand by card values")

	cardValues := strings.Split(CardValues, Separator)
	randomCardValuesIndexes := rand.Perm(len(cardValues))[:5]
	sort.Ints(randomCardValuesIndexes)
	randomHandIndexes := rand.Perm(5)

	var sortedHand Hand
	var unsortedHand Hand
	for i, cardValuesIndex := range randomCardValuesIndexes {
		card := Card{Value: cardValues[cardValuesIndex], Suit: ValidSuit()}
		sortedHand[i] = card
		unsortedHand[randomHandIndexes[i]] = card
	}

	assert.Equal(t, sortedHand, SortByValue(unsortedHand))
}
