package main

import (
	"strings"
)

type combinationMatcher interface {
	IsOnePair(hand Hand) (bool, Hand)
	IsTwoPairs(hand Hand) (bool, Hand)
	IsThreeKind(hand Hand) (bool, Hand)
	IsStraight(hand Hand) (bool, Hand)
	IsFlush(hand Hand) (bool, Hand)
	IsFullHouse(hand Hand) (bool, Hand)
	IsFourKind(hand Hand) (bool, Hand)
	IsStraightFlush(hand Hand) (bool, Hand)
	IsRoyalFlush(hand Hand) (bool, Hand)
}

type combMatcher struct {
	config config
}

func (m combMatcher) IsRoyalFlush(hand Hand) (bool, Hand) {

	if result, _ := m.IsFlush(hand); !result {
		return false, hand
	}

	for i, value := range strings.Split(m.config.cardValues, m.config.separator)[:5] {
		if hand[i].Value != value {
			return false, hand
		}
	}

	return true, hand
}

func (m combMatcher) IsStraightFlush(hand Hand) (bool, Hand) {
	flush, _ := m.IsFlush(hand)
	straight, _ := m.IsStraight(hand)

	return flush && straight, hand
}

func (m combMatcher) IsFourKind(hand Hand) (bool, Hand) {
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

func (m combMatcher) IsFullHouse(hand Hand) (bool, Hand) {
	if (hand[0].Value == hand[2].Value) && (hand[3].Value == hand[4].Value) {
		return true, hand
	}
	if (hand[0].Value == hand[1].Value) && (hand[2].Value == hand[4].Value) {
		return true, Hand{hand[2], hand[3], hand[4], hand[0], hand[1]}
	}

	return false, hand
}

func (m combMatcher) IsFlush(hand Hand) (bool, Hand) {
	for _, card := range hand {
		if card.Suit != hand[0].Suit {
			return false, hand
		}
	}

	return true, hand
}

func (m combMatcher) IsStraight(hand Hand) (bool, Hand) {
	var values []string

	for _, card := range hand {
		values = append(values, card.Value)
	}

	return strings.Contains(m.config.cardValues, strings.Join(values, m.config.separator)), hand
}

func (m combMatcher) IsThreeKind(hand Hand) (bool, Hand) {
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

func (m combMatcher) IsTwoPairs(hand Hand) (bool, Hand) {
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

func (m combMatcher) IsOnePair(hand Hand) (bool, Hand) {
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
