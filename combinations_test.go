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
	checkNoCombinations(t, IsOnePair)

	log.Println("should return true and a reordered hand when the hand contains a pair")
	checkOnePair(t, IsOnePair, true)
}

func TestIsTwoPairs(t *testing.T) {
	log.Println("IsTwoPairs")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, IsTwoPairs)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, IsTwoPairs, false)

	log.Println("should return true and a reordered hand when the hand contains 2 pairs")
	checkTwoPairs(t, IsTwoPairs, true)
}

func TestIsThreeKind(t *testing.T) {
	log.Println("IsThreeKind")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, IsThreeKind)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, IsThreeKind, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, IsThreeKind, false)

	log.Println("should return true and a reordered hand when the hand contains 3 cards of the same value")
	checkThree(t, IsThreeKind, true)
}

func TestIsStraight(t *testing.T) {
	log.Println("IsStraight")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, IsStraight)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, IsStraight, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, IsStraight, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, IsStraight, false)

	log.Println("should return true and a hand when the hand is straight")
	checkStraight(t, IsStraight, true)
}

func TestIsFlush(t *testing.T) {
	log.Println("IsFlush")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, IsFlush)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, IsFlush, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, IsFlush, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, IsFlush, false)

	log.Println("should return false and a hand when the hand is straight")
	checkStraight(t, IsFlush, false)

	log.Println("should return true and a hand when the hand is flush")
	checkFlush(t, IsFlush, true)
}

func TestIsFullHouse(t *testing.T) {
	log.Println("IsFullHouse")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, IsFullHouse)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, IsFullHouse, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, IsFullHouse, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, IsFullHouse, false)

	log.Println("should return false and a hand when the hand is straight")
	checkStraight(t, IsFullHouse, false)

	log.Println("should return false and a hand when the hand is flush")
	checkFlush(t, IsFullHouse, false)

	log.Println("should return true and a reordered hand when the hand is full house")
	checkFullHouse(t, IsFullHouse, true)
}

func TestIsFourKind(t *testing.T) {
	log.Println("IsFourKind")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, IsFourKind)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, IsFourKind, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, IsFourKind, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, IsFourKind, false)

	log.Println("should return false and a hand when the hand is straight")
	checkStraight(t, IsFourKind, false)

	log.Println("should return false and a hand when the hand is flush")
	checkFlush(t, IsFourKind, false)

	log.Println("should return false and a hand when the hand is full house")
	checkFullHouse(t, IsFourKind, false)

	log.Println("should return true and a reordered hand when the hand has four cards of the same value")
	checkFour(t, IsFourKind, true)
}

func TestIsStraightFlush(t *testing.T) {
	log.Println("IsStraightFlush")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, IsStraightFlush)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, IsStraightFlush, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, IsStraightFlush, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, IsStraightFlush, false)

	log.Println("should return false and a hand when the hand is straight")
	checkStraight(t, IsStraightFlush, false)

	log.Println("should return false and a hand when the hand is flush")
	checkFlush(t, IsStraightFlush, false)

	log.Println("should return false and a hand when the hand is full house")
	checkFullHouse(t, IsStraightFlush, false)

	log.Println("should return false and a hand when the hand has four cards of the same value")
	checkFour(t, IsStraightFlush, false)

	log.Println("should return true and a hand when the hand is a straight flush")
	checkStraightFlush(t, IsStraightFlush, true)
}

func TestIsRoyalFlush(t *testing.T) {
	log.Println("IsRoyalFlush")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, IsRoyalFlush)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, IsRoyalFlush, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, IsRoyalFlush, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, IsRoyalFlush, false)

	log.Println("should return false and a hand when the hand is straight")
	checkStraight(t, IsRoyalFlush, false)

	log.Println("should return false and a hand when the hand is flush")
	checkFlush(t, IsRoyalFlush, false)

	log.Println("should return false and a hand when the hand is full house")
	checkFullHouse(t, IsRoyalFlush, false)

	log.Println("should return false and a hand when the hand has four cards of the same value")
	checkFour(t, IsRoyalFlush, false)

	log.Println("should return false and a hand when the hand is a straight flush")
	checkStraightFlush(t, IsRoyalFlush, false)

	log.Println("should return true and a hand when the hand is a royal flush")
	func () {
		handRoyalFlush := HandRoyalFlush()

		isRoyalFlush, actualHand := IsRoyalFlush(handRoyalFlush)
		assert.Equal(t, true, isRoyalFlush)
		assert.Equal(t, handRoyalFlush, actualHand)
	}()
}

func checkNoCombinations(t *testing.T, checker func (Hand) (bool, Hand)) {
	handNoPairs := HandNoPairs()
	actualResult, actualHand := checker(handNoPairs)
	assert.Equal(t, false, actualResult)
	assert.Equal(t, handNoPairs, actualHand)
}

func checkOnePair(t *testing.T, checker func (Hand) (bool, Hand), expectedResult bool) {
	i := rand.Intn(4)
	handWithPair := HandWithPair(i)

	if expectedResult {
		expectedHand := Hand{handWithPair[i], handWithPair[i + 1]}
		for k, card := range handWithPair[:i] {
			expectedHand[2 + k] = card
		}
		for k, card := range handWithPair[i + 2:] {
			expectedHand[i + 2 + k] = card
		}

		actualResult, actualHand := checker(handWithPair)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, expectedHand, actualHand)
	} else {
		actualResult, actualHand := checker(handWithPair)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, handWithPair, actualHand)
	}
}

func checkTwoPairs(t *testing.T, checker func (Hand) (bool, Hand), expectedResult bool) {
	i := rand.Intn(2)
	var j int
	if i == 0 {
		j = 2 + rand.Intn(2)
	} else {
		j = 3
	}
	handWithPairs := HandWithTwoPairs(i, j)

	if expectedResult {
		expectedHand := Hand{handWithPairs[i], handWithPairs[i + 1], handWithPairs[j], handWithPairs[j + 1]}
		for k :=0; k < 5; k++ {
			if (k < i) || (k > j + 1) || ((k > i + 1) && (k < j)) {
				expectedHand[4] = handWithPairs[k]
			}
		}

		hasTwoPairs, actualHand := checker(handWithPairs)
		assert.Equal(t, expectedResult, hasTwoPairs)
		assert.Equal(t, expectedHand, actualHand)
	} else {
		hasTwoPairs, actualHand := checker(handWithPairs)
		assert.Equal(t, expectedResult, hasTwoPairs)
		assert.Equal(t, handWithPairs, actualHand)
	}
}

func checkThree(t *testing.T, checker func (Hand) (bool, Hand), expectedResult bool) {
	i := rand.Intn(3)
	handWithThree := HandWithThree(i)

	if expectedResult {
		expectedHand := Hand{handWithThree[i], handWithThree[i + 1], handWithThree[i + 2]}
		for j := 0; j < i; j++ {
			expectedHand[j + 3] = handWithThree[j]
		}
		for j := i + 3; j < 5; j++ {
			expectedHand[j] = handWithThree[j]
		}

		hasThree, actualHand := checker(handWithThree)
		assert.Equal(t, expectedResult, hasThree)
		assert.Equal(t, expectedHand, actualHand)
	} else {
		hasThree, actualHand := checker(handWithThree)
		assert.Equal(t, expectedResult, hasThree)
		assert.Equal(t, handWithThree, actualHand)
	}
}

func checkStraight(t *testing.T, checker func (Hand) (bool, Hand), expectedResult bool) {
	handStraight := HandStraight()
	isStraight, actualHand := checker(handStraight)
	assert.Equal(t, expectedResult, isStraight)
	assert.Equal(t, handStraight, actualHand)
}

func checkFlush(t *testing.T, checker func (Hand) (bool, Hand), expectedResult bool) {
	handFlush := HandFlush()

	isFlush, actualHand := checker(handFlush)
	assert.Equal(t, expectedResult, isFlush)
	assert.Equal(t, handFlush, actualHand)

}

func checkFullHouse(t *testing.T, checker func (Hand) (bool, Hand), expectedResult bool) {
	handFullHouse := HandFullHouse()

	expectedHand := handFullHouse
	if handFullHouse[0].Value != handFullHouse[2].Value {
		expectedHand = Hand{handFullHouse[2], handFullHouse[3], handFullHouse[4], handFullHouse[0], handFullHouse[1]}
	}

	isFullHouse, actualHand := checker(handFullHouse)
	assert.Equal(t, expectedResult, isFullHouse)
	assert.Equal(t, expectedHand, actualHand)
}

func checkFour(t *testing.T, checker func (Hand) (bool, Hand), expectedResult bool) {
	handFour := HandFour()

	expectedHand := handFour
	if handFour[0].Value != handFour[3].Value {
		expectedHand = Hand{handFour[1], handFour[2], handFour[3], handFour[4], handFour[0]}
	}

	hasFour, actualHand := checker(handFour)
	assert.Equal(t, expectedResult, hasFour)
	assert.Equal(t, expectedHand, actualHand)
}

func checkStraightFlush(t *testing.T, checker func (Hand) (bool, Hand), expectedResult bool) {
	handStraightFlush := HandStraightFlush()

	isStraightFlush, actualHand := checker(handStraightFlush)
	assert.Equal(t, expectedResult, isStraightFlush)
	assert.Equal(t, handStraightFlush, actualHand)
}
