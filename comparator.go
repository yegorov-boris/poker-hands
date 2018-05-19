package main

import (
	"strings"
)

func IsFirstPlayerWinner(hands string) (bool, error) {
	first, second, err := ParseHands(hands)
	if err != nil {
		return false, err
	}

	firstCombinationRank, firstReordered := GetCombination(first)
	secondCombinationRank, secondReordered := GetCombination(second)

	if firstCombinationRank != secondCombinationRank {
		return firstCombinationRank < secondCombinationRank, nil
	}

	for i, card := range firstReordered {
		firstIndex := strings.Index(CardValues, card.Value)
		secondIndex := strings.Index(CardValues, secondReordered[i].Value)

		if firstIndex != secondIndex {
			return firstIndex < secondIndex, nil
		}
	}

	return false, nil
}

func GetCombination(hand Hand) (int, Hand) {
	combinationCheckers := []func(Hand) (bool, Hand){
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

	i := 0
	for ; i < len(combinationCheckers); i++ {
		if isMatching, reorderedHand := combinationCheckers[i](hand); isMatching {
			return i, reorderedHand
		}
	}

	return i, hand
}
