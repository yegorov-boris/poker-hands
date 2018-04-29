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

func PickOneWithout(slice []string, exceptions []string) string {
	var diff []string
	exceptionsLength := len(exceptions)
	for _, element := range slice {
		i := 0
		for ; i < exceptionsLength; i++ {
			if element == exceptions[i] {
				break
			}
		}

		if i == exceptionsLength {
			diff = append(diff, element)
		}
	}

	return diff[rand.Intn(len(diff))]
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

func HandNoPairs() Hand {
	cardValues := strings.Split(CardValues, Separator)
	indexes := rand.Perm(len(cardValues))[:5]
	sort.Ints(indexes)

	var hand Hand
	for i, index := range indexes {
		hand[i] = Card{Value: cardValues[index], Suit: ValidSuit()}
	}

	return hand
}

func HandWithPair(i int) Hand {
	suits := strings.Split(Suits, Separator)
	hand := HandNoPairs()
	hand[i + 1] = Card{
		Value: hand[i].Value,
		Suit: PickOneWithout(suits, []string{hand[i].Suit}),
	}

	return hand
}

func HandWithTwoPairs(i, j int) Hand {
	hand := HandWithPair(i)
	hand[j + 1] = Card{
		Value: hand[j].Value,
		Suit: PickOneWithout(strings.Split(Suits, Separator), []string{hand[i].Suit}),
	}

	return hand
}

func HandWithThree(i int) Hand {
	hand := HandWithPair(i)
	hand[i + 2] = Card{
		Value: hand[i].Value,
		Suit: PickOneWithout(strings.Split(Suits, Separator), []string{hand[i].Suit, hand[i + 1].Suit}),
	}

	return hand
}

func HandStraight() Hand {
	cardValues := strings.Split(CardValues, Separator)
	i := rand.Intn(len(cardValues) - 4)

	var hand Hand
	for j, value := range cardValues[i:i + 5] {
		hand[j] = Card{Value: value, Suit: ValidSuit()}
	}

	return hand
}
