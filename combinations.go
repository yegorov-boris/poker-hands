package main

import (
	"strings"
)

type combinationMatcher interface {
	isOnePair(hand Hand) (bool, Hand)
	isTwoPairs(hand Hand) (bool, Hand)
	isThreeKind(hand Hand) (bool, Hand)
	isStraight(hand Hand) (bool, Hand)
	isFlush(hand Hand) (bool, Hand)
	isFullHouse(hand Hand) (bool, Hand)
	isFourKind(hand Hand) (bool, Hand)
	isStraightFlush(hand Hand) (bool, Hand)
	isRoyalFlush(hand Hand) (bool, Hand)
}

type combMatcher struct {
	config config
}

func (m combMatcher) isRoyalFlush(hand Hand) (bool, Hand) {

	if result, _ := m.isFlush(hand); !result {
		return false, hand
	}

	for i, value := range strings.Split(m.config.cardValues, m.config.separator)[:5] {
		if hand[i].Value != value {
			return false, hand
		}
	}

	return true, hand
}

func (m combMatcher) isStraightFlush(hand Hand) (bool, Hand) {
	flush, _ := m.isFlush(hand)
	straight, _ := m.isStraight(hand)

	return flush && straight, hand
}

func (m combMatcher) isFourKind(hand Hand) (bool, Hand) {
	if hand[0].Value == hand[3].Value {
		return true, hand
	}
	if hand[1].Value == hand[4].Value {
		var reorderedHand Hand
		for i, card := range hand[1:] {
			reorderedHand[i] = card
		}
		reorderedHand[4] = hand[0]

		return true, reorderedHand
	}

	return false, hand
}

func (m combMatcher) isFullHouse(hand Hand) (bool, Hand) {
	if (hand[0].Value == hand[2].Value) && (hand[3].Value == hand[4].Value) {
		return true, hand
	}
	if (hand[0].Value == hand[1].Value) && (hand[2].Value == hand[4].Value) {
		return true, Hand{hand[2], hand[3], hand[4], hand[0], hand[1]}
	}

	return false, hand
}

func (m combMatcher) isFlush(hand Hand) (bool, Hand) {
	for _, card := range hand {
		if card.Suit != hand[0].Suit {
			return false, hand
		}
	}

	return true, hand
}

func (m combMatcher) isStraight(hand Hand) (bool, Hand) {
	var values []string

	for _, card := range hand {
		values = append(values, card.Value)
	}

	return strings.Contains(m.config.cardValues, strings.Join(values, m.config.separator)), hand
}

func (m combMatcher) isThreeKind(hand Hand) (bool, Hand) {
	if hand[0].Value == hand[2].Value {
		return true, hand
	}
	if hand[1].Value == hand[3].Value {
		return true, Hand{hand[1], hand[2], hand[3], hand[0], hand[4]}
	}
	if hand[2].Value == hand[4].Value {
		return true, Hand{hand[2], hand[3], hand[4], hand[0], hand[1]}
	}

	return false, hand
}

func (m combMatcher) isTwoPairs(hand Hand) (bool, Hand) {
	if (hand[0].Value == hand[1].Value) && (hand[2].Value == hand[3].Value) {
		return true, hand
	}
	if (hand[1].Value == hand[2].Value) && (hand[3].Value == hand[4].Value) {
		return true, Hand{hand[1], hand[2], hand[3], hand[4], hand[0]}
	}
	if (hand[0].Value == hand[1].Value) && (hand[3].Value == hand[4].Value) {
		return true, Hand{hand[0], hand[1], hand[3], hand[4], hand[2]}
	}

	return false, hand
}

func (m combMatcher) isOnePair(hand Hand) (bool, Hand) {
	index := 0
	for ; index < 4; index++ {
		if hand[index].Value == hand[index+1].Value {
			break
		}
	}

	if index == 4 {
		return false, hand
	}

	var reorderedHand Hand
	reorderedHand[0] = hand[index]
	reorderedHand[1] = hand[index+1]
	for i, card := range hand[:index] {
		reorderedHand[i+2] = card
	}
	for i, card := range hand[index+2:] {
		reorderedHand[i+index+2] = card
	}

	return true, reorderedHand
}
