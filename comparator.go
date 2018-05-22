package main

import (
	"strings"
)

type comparator struct {
	config  config
	parser  handsStringParser
	matcher combinationMatcher
}

func (c comparator) IsFirstPlayerWinner(hands string) (bool, error) {
	first, second, err := c.parser.ParseHands(hands)
	if err != nil {
		return false, err
	}

	firstCombinationRank, firstReordered := c.GetCombination(first)
	secondCombinationRank, secondReordered := c.GetCombination(second)

	if firstCombinationRank != secondCombinationRank {
		return firstCombinationRank < secondCombinationRank, nil
	}

	for i, card := range firstReordered {
		firstIndex := strings.Index(c.config.cardValues, card.Value)
		secondIndex := strings.Index(c.config.cardValues, secondReordered[i].Value)

		if firstIndex != secondIndex {
			return firstIndex < secondIndex, nil
		}
	}

	return false, nil
}

func (c comparator) GetCombination(hand Hand) (int, Hand) {
	combinationCheckers := []func(Hand) (bool, Hand){
		func(hand Hand) (bool, Hand) {
			return c.matcher.IsRoyalFlush(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.IsStraightFlush(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.IsFourKind(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.IsFullHouse(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.IsFlush(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.IsStraight(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.IsThreeKind(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.IsTwoPairs(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.IsOnePair(hand)
		},
	}

	i := 0
	for ; i < len(combinationCheckers); i++ {
		if isMatching, reorderedHand := combinationCheckers[i](hand); isMatching {
			return i, reorderedHand
		}
	}

	return i, hand
}
