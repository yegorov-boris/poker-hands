package main

import (
	"math/rand"
	"strings"
	"sort"
)

func RandomString(minLength int, maxLength int) string {
	length := minLength + rand.Intn(maxLength - minLength + 1)
	var randomBytes []byte
	for i := 0; i < length; i++ {
		randomBytes = append(randomBytes, byte(32 + rand.Intn(95)))
	}

	return string(randomBytes)
}

func RandomStringWithout(minLength int, maxLength int, exceptions string) string {
	length := minLength + rand.Intn(maxLength - minLength + 1)
	var randomBytes []byte
	for i := 0; i < length; {
		randomByte := byte(32 + rand.Intn(95))
		if strings.IndexByte(exceptions, randomByte) == -1 {
			randomBytes = append(randomBytes, randomByte)
			i++
		}
	}

	return string(randomBytes)
}

func PickOne(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func ValidCardValue() string {
	return PickOne(strings.Split(CardValues, Separator))
}

func InvalidCardValue() string {
	return RandomStringWithout(1, 1, CardValues)
}

func ValidSuit() string {
	return PickOne(strings.Split(Suits, Separator))
}

func InvalidSuit() string {
	return RandomStringWithout(1, 1, Suits)
}

func SortedHand(exceptions []Card) Hand {
	var deck []Card
	exceptionsLength := len(exceptions)
	for _, value := range strings.Split(CardValues, Separator) {
		for _, suit := range strings.Split(Suits, Separator) {
			card := Card{Value: value, Suit: suit}
			i := 0;
			for ; i < exceptionsLength; i++ {
				if card == exceptions[i] {
					break
				}
			}
			if i == exceptionsLength {
				deck = append(deck, card)
			}
		}
	}

	indexes := rand.Perm(len(deck))[:5]
	sort.Ints(indexes)

	var sortedHand Hand
	for i, index := range indexes {
		sortedHand[i] = deck[index]
	}

	return sortedHand
}


