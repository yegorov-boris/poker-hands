package main

import (
	"testing"
	"log"
	"math/rand"
	"strings"
	"sort"
	"github.com/stretchr/testify/assert"
)

func TestIsOnePair(t *testing.T) {
	log.Println("IsOnePair")

	cardValues := strings.Split(CardValues, Separator)
	suits := strings.Split(Suits, Separator)
	indexes := rand.Perm(len(cardValues))[:5]
	sort.Ints(indexes)

	var handNoPairs Hand
	for i, index := range indexes {
		handNoPairs[i] = Card{Value: cardValues[index], Suit: ValidSuit()}
	}

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	isPair, actualHand := IsOnePair(handNoPairs)
	assert.Equal(t, false, isPair)
	assert.Equal(t, handNoPairs, actualHand)

	log.Println("should return true and a reordered hand when the hand contains a pair")

	for i := 0; i < 4; i++ {
		handWithPair := CopyHand(handNoPairs)
		dup := Card{
			Value: handWithPair[i].Value,
			Suit: PickOneWithout(suits, []string{handWithPair[i].Suit}),
		}
		handWithPair[i + 1] = dup

		expectedHand := Hand{handNoPairs[i], dup}
		for k, card := range handNoPairs[:i] {
			expectedHand[2 + k] = card
		}
		for k, card := range handNoPairs[i + 2:] {
			expectedHand[i + 2 + k] = card
		}

		isPair, actualHand := IsOnePair(handWithPair)
		assert.Equal(t, true, isPair)
		assert.Equal(t, expectedHand, actualHand)
	}
}
