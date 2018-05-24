package main

import (
	"strings"
)

type comparator interface {
	isFirstPlayerWinner(hands string) (bool, error)
}

type cmp struct {
	config  config
	parser  handsStringParser
	matcher combinationMatcher
}

func (c cmp) isFirstPlayerWinner(hands string) (bool, error) {
	first, second, err := c.parser.parseHands(hands)
	if err != nil {
		return false, err
	}

	firstCombinationRank, firstReordered := c.getCombination(first)
	secondCombinationRank, secondReordered := c.getCombination(second)

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

func (c cmp) getCombination(hand Hand) (int, Hand) {
	combinationCheckers := []func(Hand) (bool, Hand){
		func(hand Hand) (bool, Hand) {
			return c.matcher.isRoyalFlush(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.isStraightFlush(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.isFourKind(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.isFullHouse(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.isFlush(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.isStraight(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.isThreeKind(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.isTwoPairs(hand)
		},
		func(hand Hand) (bool, Hand) {
			return c.matcher.isOnePair(hand)
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
