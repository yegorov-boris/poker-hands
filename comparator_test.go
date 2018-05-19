package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"strings"
	"testing"
)

var randomHandGenerators = []func() Hand{
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

func TestGetCombination(t *testing.T) {
	log.Println("GetCombination")

	log.Println("should return the combination index and the reordered hand")

	expectedIndex := rand.Intn(len(randomHandGenerators))
	hand := randomHandGenerators[expectedIndex]()
	expectedHand := reorderHand(expectedIndex, hand)

	actualIndex, actualHand := GetCombination(hand)
	assert.Equal(t, expectedIndex, actualIndex)
	assert.Equal(t, expectedHand, actualHand)
}

func TestIsFirstPlayerWinner(t *testing.T) {
	log.Println("IsFirstPlayerWinner")

	log.Println("should fail when the parsing fails")
	func() {
		_, err := IsFirstPlayerWinner(RandomString(30, 50))
		assert.Error(t, err)
	}()

	log.Println("should return true when the first player's combination is stronger")
	func() {
		firstCombinationIndex := rand.Intn(len(randomHandGenerators) - 1)
		secondCombinationIndex := firstCombinationIndex + 1 +
			rand.Intn(len(randomHandGenerators)-firstCombinationIndex-1)

		actual, err := IsFirstPlayerWinner(validHandsString(firstCombinationIndex, secondCombinationIndex))
		assert.Nil(t, err)
		assert.True(t, actual)
	}()

	log.Println("should return false when the second player's combination is stronger")
	func() {
		firstCombinationIndex := 1 + rand.Intn(len(randomHandGenerators)-1)
		secondCombinationIndex := rand.Intn(firstCombinationIndex)

		actual, err := IsFirstPlayerWinner(validHandsString(firstCombinationIndex, secondCombinationIndex))
		assert.Nil(t, err)
		assert.False(t, actual)
	}()

	log.Println("should return true when the combinations have equal ranks ",
		"but the first player's cards are stronger")
	func() {
		actual, err := IsFirstPlayerWinner("5D 8C 9S JS AC 2C 5C 7D 8S QH")
		assert.Nil(t, err)
		assert.True(t, actual)
	}()

	log.Println("should return false when the combinations have equal ranks ",
		"but the first player's cards are not stronger")
	func() {
		actual, err := IsFirstPlayerWinner("5H 5C 6S 7S KD 2C 3S 8S 8D TD")
		assert.Nil(t, err)
		assert.False(t, actual)
	}()
}

func validHandsString(firstCombinationIndex, secondCombinationIndex int) string {
	firstHand := randomHandGenerators[firstCombinationIndex]()
	secondHand := randomHandGenerators[secondCombinationIndex]()

	for _, firstCard := range firstHand {
		for _, secondCard := range secondHand {
			if firstCard == secondCard {
				return validHandsString(firstCombinationIndex, secondCombinationIndex)
			}
		}
	}

	var values []string
	for _, card := range firstHand {
		values = append(values, strings.Join([]string{card.Value, card.Suit}, ""))
	}
	for _, card := range secondHand {
		values = append(values, strings.Join([]string{card.Value, card.Suit}, ""))
	}

	return strings.Join(values, Separator)
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
