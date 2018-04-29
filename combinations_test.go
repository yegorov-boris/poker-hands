package main

import (
	"testing"
	"log"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strings"
)

func TestIsOnePair(t *testing.T) {
	log.Println("IsOnePair")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	func () {
		handNoPairs := HandNoPairs()
		isPair, actualHand := IsOnePair(handNoPairs)
		assert.Equal(t, false, isPair)
		assert.Equal(t, handNoPairs, actualHand)
	}()

	log.Println("should return true and a reordered hand when the hand contains a pair")
	func () {
		i := rand.Intn(4)
		handWithPair := HandWithPair(i)

		expectedHand := Hand{handWithPair[i], handWithPair[i + 1]}
		for k, card := range handWithPair[:i] {
			expectedHand[2 + k] = card
		}
		for k, card := range handWithPair[i + 2:] {
			expectedHand[i + 2 + k] = card
		}

		isPair, actualHand := IsOnePair(handWithPair)
		assert.Equal(t, true, isPair)
		assert.Equal(t, expectedHand, actualHand)
	}()
}

func TestIsTwoPairs(t *testing.T) {
	log.Println("IsTwoPairs")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	func () {
		handNoPairs := HandNoPairs()
		hasTwoPairs, actualHand := IsTwoPairs(handNoPairs)
		assert.Equal(t, false, hasTwoPairs)
		assert.Equal(t, handNoPairs, actualHand)
	}()

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	func () {
		i := rand.Intn(4)
		handWithPair := HandWithPair(i)
		hasTwoPairs, actualHand := IsTwoPairs(handWithPair)
		assert.Equal(t, false, hasTwoPairs)
		assert.Equal(t, handWithPair, actualHand)
	}()

	log.Println("should return true and a reordered hand when the hand contains 2 pairs")
	func () {
		i := rand.Intn(3)
		j := i + 1 + rand.Intn(2)
		handWithPairs := HandWithPair(i)
		handWithPairs[j + 1] = Card{
			Value: handWithPairs[j].Value,
			Suit: PickOneWithout(strings.Split(Suits, Separator), []string{handWithPairs[i].Suit}),
		}

		expectedHand := Hand{handWithPairs[i], handWithPairs[i + 1], handWithPairs[j], handWithPairs[j + 1]}
		for k :=0; k < 5; k++ {
			if (k < i) || (k > j + 1) || ((k > i + 1) && (k < j)) {
				expectedHand[4] = handWithPairs[k]
			}
		}

		hasTwoPairs, actualHand := IsTwoPairs(handWithPairs)
		assert.Equal(t, true, hasTwoPairs)
		assert.Equal(t, expectedHand, actualHand)
	}()
}
