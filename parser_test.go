package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"math/rand"
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
	hand := SortedHand([]Card{})
	unsortedHand := Hand{}
	for i, index := range rand.Perm(5) {
		unsortedHand[i] = hand[index]
	}
	assert.Equal(t, hand, SortByValue(unsortedHand))
}

func TestParseHands(t *testing.T) {
	log.Println("ParseHands")

	expectedFirst := SortedHand([]Card{})
	expectedSecond := SortedHand(expectedFirst[:])
	var cardStrings []string
	for _, card := range expectedFirst {
		cardStrings = append(cardStrings, strings.Join([]string{card.Value, card.Suit}, ""))
	}
	for _, card := range expectedSecond {
		cardStrings = append(cardStrings, strings.Join([]string{card.Value, card.Suit}, ""))
	}

	log.Println("should parse a string which contains two hands")
	actualFirst, actualSecond, _ := ParseHands(strings.Join(cardStrings, Separator))
	assert.Equal(t, expectedFirst, actualFirst)
	assert.Equal(t, expectedSecond, actualSecond)

	log.Println("should fail when the string is not a ten valid encoded cards (e.g. 9D) separated by a whitespace")
	_, _, errWrongLength := ParseHands(RandomString(10, 50))
	assert.Errorf(t, errWrongLength, "failed to parse a line with hands: wrong length")

	log.Println("should fail when the string contains duplicated encoded cards")
	_, _, errNotUnique := ParseHands(strings.Join(append(cardStrings[:9], cardStrings[0]), Separator))
	assert.Errorf(t, errNotUnique, "failed to parse a line with hands: %s is not unique", cardStrings[0])

	log.Println("should fail when the card string parser fails")
	_, _, errCardString := ParseHands(strings.Join(append(cardStrings[:9], RandomString(1, 5)), Separator))
	assert.Error(t, errCardString)
}
