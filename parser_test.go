package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"log"
)

func TestParseCardString(t *testing.T) {
	log.Print("should parse a valid card string");
	expected := Card {Suit: "D", Value: "9"}
	actual := ParseCardString("9D")
	assert.Equal(t, expected, actual)
}
