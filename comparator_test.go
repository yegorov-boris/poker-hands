package main

import (
	"testing"
	"log"
	"math/rand"
	"github.com/stretchr/testify/assert"
)

func TestGetCombination(t *testing.T) {
	log.Println("GetCombination")

	log.Println("should return the combination index and the reordered hand")
	randomHandGenerators := []func () Hand {
		HandRoyalFlush,
		HandStraightFlush,
		HandFour,
		HandFullHouse,
		HandFlush,
		HandStraight,
		HandWithThree,
		HandWithTwoPairs,
		HandWithPair,
		HandNoPairs,
	}
	expectedIndex := rand.Intn(len(randomHandGenerators))
	hand := randomHandGenerators[expectedIndex]()
	expectedHand := reorderHand(expectedIndex, hand)


	actualIndex, actualHand := GetCombination(hand)
	assert.Equal(t, expectedIndex, actualIndex)
	assert.Equal(t, expectedHand, actualHand)
}

func reorderHand(combinationIndex int, hand Hand) Hand {
	if combinationIndex == 8 {
		return ReorderOnePair(hand)
	}
	if combinationIndex == 7 {
		return ReorderTwoPairs(hand)
	}
	if combinationIndex == 6 {
		return ReorderThree(hand)
	}
	if combinationIndex == 3 {
		return ReorderFullHouse(hand)
	}
	if combinationIndex == 2 {
		return ReorderFour(hand)
	}

	return hand
}
