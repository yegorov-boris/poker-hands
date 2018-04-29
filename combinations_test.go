package main

import (
	"testing"
	"log"
	"github.com/stretchr/testify/assert"
	"math/rand"
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
		handWithPairs := HandWithTwoPairs(i, j)

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

func TestIsThreeKind(t *testing.T) {
	log.Println("IsThreeKind")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	func () {
		handNoPairs := HandNoPairs()
		hasThree, actualHand := IsThreeKind(handNoPairs)
		assert.Equal(t, false, hasThree)
		assert.Equal(t, handNoPairs, actualHand)
	}()

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	func () {
		i := rand.Intn(4)
		handWithPair := HandWithPair(i)
		hasThree, actualHand := IsThreeKind(handWithPair)
		assert.Equal(t, false, hasThree)
		assert.Equal(t, handWithPair, actualHand)
	}()

	log.Println("should return false and a hand when the hand contains 2 pairs")
	func () {
		i := rand.Intn(2)
		j := i + 2 + rand.Intn(2)
		handWithPairs := HandWithTwoPairs(i, j)

		expectedHand := Hand{handWithPairs[i], handWithPairs[i + 1], handWithPairs[j], handWithPairs[j + 1]}
		for k :=0; k < 5; k++ {
			if (k < i) || (k > j + 1) || ((k > i + 1) && (k < j)) {
				expectedHand[4] = handWithPairs[k]
			}
		}

		hasThree, actualHand := IsThreeKind(handWithPairs)
		assert.Equal(t, false, hasThree)
		assert.Equal(t, expectedHand, actualHand)
	}()

	log.Println("should return true and a reordered hand when the hand contains 3 cards of the same value")
	func () {
		i := rand.Intn(3)
		handWithThree := HandWithThree(i)

		expectedHand := Hand{handWithThree[i], handWithThree[i + 1], handWithThree[i + 2]}
		for j := 0; j < i; j++ {
			expectedHand[j + 3] = handWithThree[j]
		}
		for j := i + 3; j < 5; j++ {
			expectedHand[j] = handWithThree[j]
		}

		hasThree, actualHand := IsThreeKind(handWithThree)
		assert.Equal(t, true, hasThree)
		assert.Equal(t, expectedHand, actualHand)
	}()
}
