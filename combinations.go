package main

import (
	"strings"
)

func IsRoyalFlush(hand Hand) (bool, Hand) {

	if result, _ := IsFlush(hand); !result {
		return false, hand
	}

	values := [5]string {"A", "K", "Q", "J", "T"}

	for i, value := range values {
		if hand[i].Value != value {
			return false, hand
		}
	}

	return true, hand
}

func IsStraightFlush(hand Hand) (bool, Hand) {
	flush, _ := IsFlush(hand)
	straight, _ := IsStraight(hand)

	return flush && straight, hand
}

func IsFourKind(hand Hand) (bool, Hand) {
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

func IsFullHouse(hand Hand) (bool, Hand) {
	if (hand[0].Value == hand[2].Value) && (hand[3].Value == hand[4].Value) {
		return true, hand
	}
	if (hand[0].Value == hand[1].Value) && (hand[2].Value == hand[4].Value) {
		return true, Hand {hand[2], hand[3], hand[4], hand[0], hand[1]}
	}

	return false, hand
}

func IsFlush(hand Hand) (bool, Hand) {
	for _, card := range hand {
		if card.Suit != hand[0].Suit {
			return false, hand
		}
	}

	return true, hand
}

func IsStraight(hand Hand) (bool, Hand) {
	var values [5]string

	for i, card := range hand {
		values[i] = card.Value
	}

	return strings.Contains(CardValues, strings.Join(values[:], Separator)), hand
}

func IsThreeKind(hand Hand) (bool, Hand) {
	if hand[0].Value == hand[2].Value {
		return true, hand
	}
	if hand[1].Value == hand[3].Value {
		return true, Hand {hand[1], hand[2], hand[3], hand[0], hand[4]}
	}
	if hand[2].Value == hand[4].Value {
		return true, Hand {hand[2], hand[3], hand[4], hand[0], hand[1]}
	}

	return false, hand
}

func IsTwoPairs(hand Hand) (bool, Hand) {
	if (hand[0].Value == hand[1].Value) && (hand[2].Value == hand[3].Value) {
		return true, hand
	}
	if (hand[1].Value == hand[2].Value) && (hand[3].Value == hand[4].Value) {
		return true, Hand {hand[1], hand[2], hand[3], hand[4], hand[0]}
	}
	if (hand[0].Value == hand[1].Value) && (hand[3].Value == hand[4].Value) {
		return true, Hand {hand[0], hand[1], hand[3], hand[4], hand[2]}
	}

	return false, hand
}

func IsOnePair(hand Hand) (bool, Hand) {
	index := 0
	for ; index < 4; index++ {
		if hand[index].Value == hand[index + 1].Value {
			break
		}
	}

	if index == 4 {
		return false, hand
	}

	var reorderedHand Hand
	reorderedHand[0] = hand[index]
	reorderedHand[1] = hand[index + 1]
	for i, card := range hand[:index] {
		reorderedHand[i + 2] = card
	}
	for i, card := range hand[index + 2:] {
		reorderedHand[i + index + 2] = card
	}

	return true, reorderedHand
}
