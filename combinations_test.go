package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var defaultMatcher = combMatcher{config: defaultConfig()}

func TestIsOnePair(t *testing.T) {
	log.Println("IsOnePair")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.IsOnePair)

	log.Println("should return true and a reordered hand when the hand contains a pair")
	checkOnePair(t, defaultMatcher.IsOnePair, true)
}

func TestIsTwoPairs(t *testing.T) {
	log.Println("IsTwoPairs")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.IsTwoPairs)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.IsTwoPairs, false)

	log.Println("should return true and a reordered hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.IsTwoPairs, true)
}

func TestIsThreeKind(t *testing.T) {
	log.Println("IsThreeKind")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.IsThreeKind)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.IsThreeKind, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.IsThreeKind, false)

	log.Println("should return true and a reordered hand when the hand contains 3 cards of the same value")
	checkThree(t, defaultMatcher.IsThreeKind, true)
}

func TestIsStraight(t *testing.T) {
	log.Println("IsStraight")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.IsStraight)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.IsStraight, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.IsStraight, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, defaultMatcher.IsStraight, false)

	log.Println("should return true and a hand when the hand is straight")
	checkStraight(t, defaultMatcher.IsStraight, true)
}

func TestIsFlush(t *testing.T) {
	log.Println("IsFlush")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.IsFlush)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.IsFlush, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.IsFlush, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, defaultMatcher.IsFlush, false)

	log.Println("should return false and a hand when the hand is straight")
	checkStraight(t, defaultMatcher.IsFlush, false)

	log.Println("should return true and a hand when the hand is flush")
	checkFlush(t, defaultMatcher.IsFlush, true)
}

func TestIsFullHouse(t *testing.T) {
	log.Println("IsFullHouse")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.IsFullHouse)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.IsFullHouse, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.IsFullHouse, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, defaultMatcher.IsFullHouse, false)

	log.Println("should return false and a hand when the hand is straight")
	checkStraight(t, defaultMatcher.IsFullHouse, false)

	log.Println("should return false and a hand when the hand is flush")
	checkFlush(t, defaultMatcher.IsFullHouse, false)

	log.Println("should return true and a reordered hand when the hand is full house")
	checkFullHouse(t, defaultMatcher.IsFullHouse, true)
}

func TestIsFourKind(t *testing.T) {
	log.Println("IsFourKind")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.IsFourKind)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.IsFourKind, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.IsFourKind, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, defaultMatcher.IsFourKind, false)

	log.Println("should return false and a hand when the hand is straight")
	checkStraight(t, defaultMatcher.IsFourKind, false)

	log.Println("should return false and a hand when the hand is flush")
	checkFlush(t, defaultMatcher.IsFourKind, false)

	log.Println("should return false and a hand when the hand is full house")
	checkFullHouse(t, defaultMatcher.IsFourKind, false)

	log.Println("should return true and a reordered hand when the hand has four cards of the same value")
	checkFour(t, defaultMatcher.IsFourKind, true)
}

func TestIsStraightFlush(t *testing.T) {
	log.Println("IsStraightFlush")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.IsStraightFlush)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.IsStraightFlush, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.IsStraightFlush, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, defaultMatcher.IsStraightFlush, false)

	log.Println("should return false and a hand when the hand is straight")
	checkStraight(t, defaultMatcher.IsStraightFlush, false)

	log.Println("should return false and a hand when the hand is flush")
	checkFlush(t, defaultMatcher.IsStraightFlush, false)

	log.Println("should return false and a hand when the hand is full house")
	checkFullHouse(t, defaultMatcher.IsStraightFlush, false)

	log.Println("should return false and a hand when the hand has four cards of the same value")
	checkFour(t, defaultMatcher.IsStraightFlush, false)

	log.Println("should return true and a hand when the hand is a straight flush")
	checkStraightFlush(t, defaultMatcher.IsStraightFlush, true)
}

func TestIsRoyalFlush(t *testing.T) {
	log.Println("IsRoyalFlush")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.IsRoyalFlush)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.IsRoyalFlush, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.IsRoyalFlush, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, defaultMatcher.IsRoyalFlush, false)

	log.Println("should return false and a hand when the hand is straight")
	checkStraight(t, defaultMatcher.IsRoyalFlush, false)

	log.Println("should return false and a hand when the hand is flush")
	checkFlush(t, defaultMatcher.IsRoyalFlush, false)

	log.Println("should return false and a hand when the hand is full house")
	checkFullHouse(t, defaultMatcher.IsRoyalFlush, false)

	log.Println("should return false and a hand when the hand has four cards of the same value")
	checkFour(t, defaultMatcher.IsRoyalFlush, false)

	log.Println("should return false and a hand when the hand is a straight flush")
	checkStraightFlush(t, defaultMatcher.IsRoyalFlush, false)

	log.Println("should return true and a hand when the hand is a royal flush")
	func() {
		handRoyalFlush := HandRoyalFlush()

		isRoyalFlush, actualHand := defaultMatcher.IsRoyalFlush(handRoyalFlush)
		assert.Equal(t, true, isRoyalFlush)
		assert.Equal(t, handRoyalFlush, actualHand)
	}()
}

func checkNoCombinations(t *testing.T, checker func(Hand) (bool, Hand)) {
	handNoPairs := HandNoPairs()
	actualResult, actualHand := checker(handNoPairs)
	assert.Equal(t, false, actualResult)
	assert.Equal(t, handNoPairs, actualHand)
}

func checkOnePair(t *testing.T, checker func(Hand) (bool, Hand), expectedResult bool) {
	handWithPair := HandWithPair()

	if expectedResult {
		expectedHand := ReorderOnePair(handWithPair)

		actualResult, actualHand := checker(handWithPair)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, expectedHand, actualHand)
	} else {
		actualResult, actualHand := checker(handWithPair)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, handWithPair, actualHand)
	}
}

func checkTwoPairs(t *testing.T, checker func(Hand) (bool, Hand), expectedResult bool) {
	handWithPairs := HandWithTwoPairs()

	if expectedResult {
		expectedHand := ReorderTwoPairs(handWithPairs)

		hasTwoPairs, actualHand := checker(handWithPairs)
		assert.Equal(t, expectedResult, hasTwoPairs)
		assert.Equal(t, expectedHand, actualHand)
	} else {
		hasTwoPairs, actualHand := checker(handWithPairs)
		assert.Equal(t, expectedResult, hasTwoPairs)
		assert.Equal(t, handWithPairs, actualHand)
	}
}

func checkThree(t *testing.T, checker func(Hand) (bool, Hand), expectedResult bool) {
	handWithThree := HandWithThree()

	if expectedResult {
		expectedHand := ReorderThree(handWithThree)

		hasThree, actualHand := checker(handWithThree)
		assert.Equal(t, expectedResult, hasThree)
		assert.Equal(t, expectedHand, actualHand)
	} else {
		hasThree, actualHand := checker(handWithThree)
		assert.Equal(t, expectedResult, hasThree)
		assert.Equal(t, handWithThree, actualHand)
	}
}

func checkStraight(t *testing.T, checker func(Hand) (bool, Hand), expectedResult bool) {
	handStraight := HandStraight()
	isStraight, actualHand := checker(handStraight)
	assert.Equal(t, expectedResult, isStraight)
	assert.Equal(t, handStraight, actualHand)
}

func checkFlush(t *testing.T, checker func(Hand) (bool, Hand), expectedResult bool) {
	handFlush := HandFlush()

	isFlush, actualHand := checker(handFlush)
	assert.Equal(t, expectedResult, isFlush)
	assert.Equal(t, handFlush, actualHand)

}

func checkFullHouse(t *testing.T, checker func(Hand) (bool, Hand), expectedResult bool) {
	handFullHouse := HandFullHouse()
	expectedHand := ReorderFullHouse(handFullHouse)

	isFullHouse, actualHand := checker(handFullHouse)
	assert.Equal(t, expectedResult, isFullHouse)
	assert.Equal(t, expectedHand, actualHand)
}

func checkFour(t *testing.T, checker func(Hand) (bool, Hand), expectedResult bool) {
	handFour := HandFour()
	expectedHand := ReorderFour(handFour)

	hasFour, actualHand := checker(handFour)
	assert.Equal(t, expectedResult, hasFour)
	assert.Equal(t, expectedHand, actualHand)
}

func checkStraightFlush(t *testing.T, checker func(Hand) (bool, Hand), expectedResult bool) {
	handStraightFlush := HandStraightFlush()

	isStraightFlush, actualHand := checker(handStraightFlush)
	assert.Equal(t, expectedResult, isStraightFlush)
	assert.Equal(t, handStraightFlush, actualHand)
}
