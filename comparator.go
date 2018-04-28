package main

import (
	"strings"
	"log"
)

func IsFirstPlayerWinner(hands string) bool {
	first, second, err := ParseHands(hands)
	if err != nil {
		log.Fatal(err)
	}

	firstCombinationRank, firstReordered := GetCombination(first)
	secondCombinationRank, secondReordered := GetCombination(second)

	if firstCombinationRank != secondCombinationRank {
		return firstCombinationRank < secondCombinationRank
	}

	for i, card := range firstReordered {
		firstIndex := strings.Index(CardValues, card.Value)
		secondIndex := strings.Index(CardValues, secondReordered[i].Value)

		if firstIndex != secondIndex {
			return firstIndex < secondIndex
		}
	}

	return false
}

func GetCombination(hand Hand) (int, Hand) {
	combinationCheckers := []func(Hand) (bool, Hand) {
		IsRoyalFlush,
		IsStraightFlush,
		IsFourKind,
		IsFullHouse,
		IsFlush,
		IsStraight,
		IsThreeKind,
		IsTwoPairs,
		IsOnePair,
	}

	for i, checker := range combinationCheckers {
		if isMatching, reorderedHand := checker(hand); isMatching {
			return i, reorderedHand
		}
	}

	return len(combinationCheckers) - 1, hand
}
