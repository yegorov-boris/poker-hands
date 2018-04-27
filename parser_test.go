package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

func TestParseCardString(t *testing.T) {
	log.Println("ParseCardString")

	validCardValue := PickOne(strings.Split(CardValues, " "))
	validSuit := PickOne(strings.Split(Suits, " "))
	invalidCardValue := RandomStringWithout(1, 1, CardValues)
	invalidSuit := RandomStringWithout(1, 1, Suits)

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
	_, errInvalidSuit := ParseCardString(strings.Join([]string{validCardValue, invalidSuit}, ""))
	assert.Errorf(t, errInvalidSuit, "failed to parse an encoded card: wrong suit")

	log.Println("should fail when a card string has invalid card value")
	_, errInvalidCardValue := ParseCardString(strings.Join([]string{invalidCardValue, validSuit}, ""))
	assert.Errorf(t, errInvalidCardValue, "failed to parse an encoded card: wrong card value")
}
