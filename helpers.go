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

	// check if it's not straight
	if indexes[4] == indexes[0] + 4 {
		return HandNoPairs()
	}

	var hand Hand
	for i, index := range indexes {
		hand[i] = Card{Value: cardValues[index], Suit: ValidSuit()}
	}

	// check if it's not flush
	if isFlush(hand) {
		return HandNoPairs()
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

	// check if it's not flush
	if isFlush(hand) {
		return HandWithPair(i)
	}

	return hand
}

func HandWithTwoPairs(i, j int) Hand {
	hand := HandWithPair(i)
	hand[j + 1] = Card{
		Value: hand[j].Value,
		Suit: PickOneWithout(strings.Split(Suits, Separator), []string{hand[i].Suit}),
	}

	// check if it's not flush
	if isFlush(hand) {
		return HandWithTwoPairs(i, j)
	}

	return hand
}

func HandWithThree(i int) Hand {
	hand := HandWithPair(i)
	hand[i + 2] = Card{
		Value: hand[i].Value,
		Suit: PickOneWithout(strings.Split(Suits, Separator), []string{hand[i].Suit, hand[i + 1].Suit}),
	}

	// check if it's not flush
	if isFlush(hand) {
		return HandWithThree(i)
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

	// check if it's not flush
	if isFlush(hand) {
		return HandStraight()
	}

	return hand
}

func isFlush(hand Hand) bool {
	expected := [5]string{hand[0].Suit, hand[0].Suit, hand[0].Suit, hand[0].Suit, hand[0].Suit}
	actual := [5]string{hand[0].Suit, hand[1].Suit, hand[2].Suit, hand[3].Suit, hand[4].Suit}
	return expected == actual
}

func HandFlush() Hand {
	hand := HandNoPairs()
	for i := 1; i < 5; i++ {
		hand[i].Suit = hand[0].Suit
	}

	return hand
}

func HandFullHouse() Hand {
	cardValues := strings.Split(CardValues, Separator)
	suits := strings.Split(Suits, Separator)
	valuesIndexes := rand.Perm(len(cardValues))[:2]
	suitsIndexes := rand.Perm(len(suits))[:3]

	var hand Hand
	if valuesIndexes[0] < valuesIndexes[1] {
		for i := 0; i < 3; i++ {
			hand[i] = Card{Value: cardValues[valuesIndexes[0]], Suit: suits[suitsIndexes[i]]}
		}
		for i := 0; i < 2; i++ {
			hand[i + 3] = Card{Value: cardValues[valuesIndexes[1]], Suit: suits[suitsIndexes[i]]}
		}
	} else {
		for i := 0; i < 2; i++ {
			hand[i] = Card{Value: cardValues[valuesIndexes[1]], Suit: suits[suitsIndexes[i]]}
		}
		for i := 0; i < 3; i++ {
			hand[i + 2] = Card{Value: cardValues[valuesIndexes[0]], Suit: suits[suitsIndexes[i]]}
		}
	}

	return hand
}

func HandFour() Hand {
	cardValues := strings.Split(CardValues, Separator)
	suits := strings.Split(Suits, Separator)
	valuesIndexes := rand.Perm(len(cardValues))[:2]

	var hand Hand
	if valuesIndexes[0] < valuesIndexes[1] {
		for i, suit := range suits {
			hand[i] = Card{Value: cardValues[valuesIndexes[0]], Suit: suit}
		}
		hand[4] = Card{Value: cardValues[valuesIndexes[1]], Suit: ValidSuit()}
	} else {
		hand[0] = Card{Value: cardValues[valuesIndexes[1]], Suit: ValidSuit()}
		for i, suit := range suits {
			hand[i + 1] = Card{Value: cardValues[valuesIndexes[0]], Suit: suit}
		}
	}

	return hand
}

func HandStraightFlush() Hand {
	cardValues := strings.Split(CardValues, Separator)
	i := rand.Intn(len(cardValues) - 4)
	suit := PickOne(strings.Split(Suits, Separator))

	var hand Hand
	for j, value := range cardValues[i:i + 5] {
		hand[j] = Card{Value: value, Suit: suit}
	}

	return hand
}
