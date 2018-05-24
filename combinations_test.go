package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var defaultMatcher = combMatcher{config: defaultConfig()}

func TestIsOnePair(t *testing.T) {
	log.Println("isOnePair")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.isOnePair)

	log.Println("should return true and a reordered hand when the hand contains a pair")
	checkOnePair(t, defaultMatcher.isOnePair, true)
}

func TestIsTwoPairs(t *testing.T) {
	log.Println("isTwoPairs")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.isTwoPairs)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.isTwoPairs, false)

	log.Println("should return true and a reordered hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.isTwoPairs, true)
}

func TestIsThreeKind(t *testing.T) {
	log.Println("isThreeKind")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.isThreeKind)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.isThreeKind, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.isThreeKind, false)

	log.Println("should return true and a reordered hand when the hand contains 3 cards of the same value")
	checkThree(t, defaultMatcher.isThreeKind, true)
}

func TestIsStraight(t *testing.T) {
	log.Println("isStraight")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.isStraight)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.isStraight, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.isStraight, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, defaultMatcher.isStraight, false)

	log.Println("should return true and a hand when the hand is straight")
	checkStraight(t, defaultMatcher.isStraight, true)
}

func TestIsFlush(t *testing.T) {
	log.Println("isFlush")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.isFlush)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.isFlush, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.isFlush, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, defaultMatcher.isFlush, false)

	log.Println("should return false and a hand when the hand is straight")
	checkStraight(t, defaultMatcher.isFlush, false)

	log.Println("should return true and a hand when the hand is flush")
	checkFlush(t, defaultMatcher.isFlush, true)
}

func TestIsFullHouse(t *testing.T) {
	log.Println("isFullHouse")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.isFullHouse)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.isFullHouse, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.isFullHouse, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, defaultMatcher.isFullHouse, false)

	log.Println("should return false and a hand when the hand is straight")
	checkStraight(t, defaultMatcher.isFullHouse, false)

	log.Println("should return false and a hand when the hand is flush")
	checkFlush(t, defaultMatcher.isFullHouse, false)

	log.Println("should return true and a reordered hand when the hand is full house")
	checkFullHouse(t, defaultMatcher.isFullHouse, true)
}

func TestIsFourKind(t *testing.T) {
	log.Println("isFourKind")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.isFourKind)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.isFourKind, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.isFourKind, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, defaultMatcher.isFourKind, false)

	log.Println("should return false and a hand when the hand is straight")
	checkStraight(t, defaultMatcher.isFourKind, false)

	log.Println("should return false and a hand when the hand is flush")
	checkFlush(t, defaultMatcher.isFourKind, false)

	log.Println("should return false and a hand when the hand is full house")
	checkFullHouse(t, defaultMatcher.isFourKind, false)

	log.Println("should return true and a reordered hand when the hand has four cards of the same value")
	checkFour(t, defaultMatcher.isFourKind, true)
}

func TestIsStraightFlush(t *testing.T) {
	log.Println("isStraightFlush")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.isStraightFlush)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.isStraightFlush, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.isStraightFlush, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, defaultMatcher.isStraightFlush, false)

	log.Println("should return false and a hand when the hand is straight")
	checkStraight(t, defaultMatcher.isStraightFlush, false)

	log.Println("should return false and a hand when the hand is flush")
	checkFlush(t, defaultMatcher.isStraightFlush, false)

	log.Println("should return false and a hand when the hand is full house")
	checkFullHouse(t, defaultMatcher.isStraightFlush, false)

	log.Println("should return false and a hand when the hand has four cards of the same value")
	checkFour(t, defaultMatcher.isStraightFlush, false)

	log.Println("should return true and a hand when the hand is a straight flush")
	checkStraightFlush(t, defaultMatcher.isStraightFlush, true)
}

func TestIsRoyalFlush(t *testing.T) {
	log.Println("isRoyalFlush")

	log.Println("should return false and a hand when the hand doesn't contain a pair")
	checkNoCombinations(t, defaultMatcher.isRoyalFlush)

	log.Println("should return false and a hand when the hand contains exactly 1 pair")
	checkOnePair(t, defaultMatcher.isRoyalFlush, false)

	log.Println("should return false and a hand when the hand contains 2 pairs")
	checkTwoPairs(t, defaultMatcher.isRoyalFlush, false)

	log.Println("should return false and a hand when the hand contains 3 cards of the same value")
	checkThree(t, defaultMatcher.isRoyalFlush, false)

	log.Println("should return false and a hand when the hand is straight")
	checkStraight(t, defaultMatcher.isRoyalFlush, false)

	log.Println("should return false and a hand when the hand is flush")
	checkFlush(t, defaultMatcher.isRoyalFlush, false)

	log.Println("should return false and a hand when the hand is full house")
	checkFullHouse(t, defaultMatcher.isRoyalFlush, false)

	log.Println("should return false and a hand when the hand has four cards of the same value")
	checkFour(t, defaultMatcher.isRoyalFlush, false)

	log.Println("should return false and a hand when the hand is a straight flush")
	checkStraightFlush(t, defaultMatcher.isRoyalFlush, false)

	log.Println("should return true and a hand when the hand is a royal flush")
	func() {
		handRoyalFlush := royalFlushHand()

		result, actualHand := defaultMatcher.isRoyalFlush(handRoyalFlush)
		assert.Equal(t, true, result)
		assert.Equal(t, handRoyalFlush, actualHand)
	}()
}

func checkNoCombinations(t *testing.T, checker func(Hand) (bool, Hand)) {
	handNoPairs := noPairsHand()
	actualResult, actualHand := checker(handNoPairs)
	assert.Equal(t, false, actualResult)
	assert.Equal(t, handNoPairs, actualHand)
}

func checkOnePair(t *testing.T, checker func(Hand) (bool, Hand), expectedResult bool) {
	handWithPair := onePairHand()

	if expectedResult {
		expectedHand := reorderOnePair(handWithPair)

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
	handWithPairs := twoPairsHand()

	if expectedResult {
		expectedHand := reorderTwoPairs(handWithPairs)

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
	handWithThree := threeHand()

	if expectedResult {
		expectedHand := reorderThree(handWithThree)

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
	handStraight := straightHand()
	straight, actualHand := checker(handStraight)
	assert.Equal(t, expectedResult, straight)
	assert.Equal(t, handStraight, actualHand)
}

func checkFlush(t *testing.T, checker func(Hand) (bool, Hand), expectedResult bool) {
	handFlush := flushHand()

	result, actualHand := checker(handFlush)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, handFlush, actualHand)

}

func checkFullHouse(t *testing.T, checker func(Hand) (bool, Hand), expectedResult bool) {
	handFullHouse := fullHouseHand()
	expectedHand := reorderFullHouse(handFullHouse)

	result, actualHand := checker(handFullHouse)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedHand, actualHand)
}

func checkFour(t *testing.T, checker func(Hand) (bool, Hand), expectedResult bool) {
	handFour := fourHand()
	expectedHand := reorderFour(handFour)

	hasFour, actualHand := checker(handFour)
	assert.Equal(t, expectedResult, hasFour)
	assert.Equal(t, expectedHand, actualHand)
}

func checkStraightFlush(t *testing.T, checker func(Hand) (bool, Hand), expectedResult bool) {
	handStraightFlush := straightFlushHand()

	result, actualHand := checker(handStraightFlush)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, handStraightFlush, actualHand)
}
