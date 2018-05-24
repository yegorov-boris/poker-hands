package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"math/rand"
	"strings"
	"testing"
)

func TestSplitHandsString(t *testing.T) {
	log.Println("splitHandsString")

	log.Println("should split a string by a separator")
	func() {
		expected := []string{}
		separator := randomString(1, 3)
		for i := 0; i < 10; i++ {
			expected = append(expected, randomStringWithout(1, 3, separator))
		}

		fakeConfig := config{separator: separator}
		actual, err := (splitter{config: fakeConfig}).splitHandsString(strings.Join(expected, separator))
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	}()

	log.Println("should fail when the number of the substrings doesn't equal 10")
	func() {
		separator := randomString(1, 3)
		fakeConfig := config{separator: separator}
		_, err := (splitter{config: fakeConfig}).splitHandsString(randomString(1, 3))
		assert.Errorf(t, err, "failed to parse a line with hands: wrong length")
	}()
}

func TestParseCardString(t *testing.T) {
	log.Println("parseCardString")

	log.Println("should parse a valid card string")
	func() {
		suit := randomString(1, 1)
		cardValue := randomString(1, 1)
		expected := Card{Suit: suit, Value: cardValue}
		cardString := strings.Join([]string{cardValue, suit}, "")
		fakeConfig := config{
			suits:      suit,
			cardValues: cardValue,
		}

		actual, err := (cardParser{config: fakeConfig}).parseCardString(cardString)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	}()

	log.Println("should fail when a card string is longer than 2 bytes")
	func() {
		_, err := (cardParser{config: defaultConfig()}).parseCardString(randomString(3, 5))
		assert.Errorf(t, err, "failed to parse an encoded card: wrong length")
	}()

	log.Println("should fail when a card string is shorter than 2 bytes")
	func() {
		_, err := (cardParser{config: defaultConfig()}).parseCardString(randomString(0, 1))
		assert.Errorf(t, err, "failed to parse an encoded card: wrong length")
	}()

	log.Println("should fail when a card string has invalid suit")
	func() {
		suit := randomString(1, 1)
		cardValue := randomString(1, 1)
		cardString := strings.Join([]string{cardValue, suit}, "")
		fakeConfig := config{
			suits:      randomStringWithout(1, 1, suit),
			cardValues: cardValue,
		}

		_, err := (cardParser{config: fakeConfig}).parseCardString(cardString)
		assert.Errorf(t, err, "failed to parse an encoded card: wrong suit")
	}()

	log.Println("should fail when a card string has invalid card value")
	func() {
		suit := randomString(1, 1)
		cardValue := randomString(1, 1)
		cardString := strings.Join([]string{cardValue, suit}, "")
		fakeConfig := config{
			suits:      suit,
			cardValues: randomStringWithout(1, 1, cardString),
		}

		_, err := (cardParser{config: fakeConfig}).parseCardString(cardString)
		assert.Errorf(t, err, "failed to parse an encoded card: wrong card value")
	}()
}

func TestSortByValue(t *testing.T) {
	log.Println("sortByValue")

	log.Println("should sort a hand by card values")

	separator := randomString(1, 1)
	cardValues := []string{}
	for i := 0; i < 5; i++ {
		cardValues = append(cardValues, randomStringWithout(1, 1, separator))
	}

	unsortedHand := Hand{}
	for i, index := range rand.Perm(5) {
		unsortedHand[i] = Card{
			Suit:  randomString(1, 1),
			Value: cardValues[index],
		}
	}

	fakeConfig := config{
		cardValues: strings.Join(cardValues, separator),
		separator:  separator,
	}

	for i, card := range (sorter{config: fakeConfig}).sortByValue(unsortedHand) {
		assert.Equal(t, cardValues[i], card.Value)
	}
}

func TestParseHands(t *testing.T) {
	log.Println("parseHands")

	log.Println("should parse a string which contains two hands")
	func() {
		fakeCardStrings := strings.Split("0123456789", "")
		fakeCard := Card{Value: "1", Suit: "a"}
		fakeHand := Hand{fakeCard, fakeCard, fakeCard, fakeCard, fakeCard}

		fakeSplitter := &mockHandsStringSplitter{}
		fakeSplitter.On("splitHandsString", mock.AnythingOfType("string")).
			Return(fakeCardStrings, nil).
			Once()

		fakeParser := &mockCardStringParser{}
		fakeParser.On("parseCardString", mock.AnythingOfType("string")).
			Return(fakeCard, nil).
			Times(10)

		fakeSorter := &mockSorterByValue{}
		fakeSorter.On("sortByValue", fakeHand).
			Return(fakeHand).
			Twice()

		fakeHandsParser := handsParser{
			splitter:   fakeSplitter,
			cardParser: fakeParser,
			sorter:     fakeSorter,
		}

		first, second, err := fakeHandsParser.parseHands(randomString(1, 3))
		assert.Nil(t, err)
		assert.Equal(t, fakeHand, first)
		assert.Equal(t, fakeHand, second)
		fakeSplitter.AssertExpectations(t)
		fakeParser.AssertExpectations(t)
		fakeSorter.AssertExpectations(t)
	}()

	log.Println("should fail when the splitHandsString fails")
	func() {
		msg := randomString(10, 20)

		fakeSplitter := &mockHandsStringSplitter{}
		fakeSplitter.On("splitHandsString", mock.AnythingOfType("string")).
			Return([]string{}, errors.New(msg)).
			Once()

		_, _, err := (handsParser{splitter: fakeSplitter}).parseHands(randomString(1, 3))
		assert.Errorf(t, err, msg)
		fakeSplitter.AssertExpectations(t)
	}()

	log.Println("should fail when the cards string contains duplicated encoded cards")
	func() {
		fakeCardString := randomString(1, 3)
		fakeHandsString := strings.Join([]string{fakeCardString, fakeCardString}, " ")

		fakeSplitter := &mockHandsStringSplitter{}
		fakeSplitter.On("splitHandsString", fakeHandsString).
			Return([]string{fakeCardString}, nil).
			Once()

		_, _, err := (handsParser{splitter: fakeSplitter}).parseHands(fakeHandsString)
		assert.Errorf(t, err, "failed to parse a line with hands: %s is not unique", fakeCardString)
		fakeSplitter.AssertExpectations(t)
	}()

	log.Println("should fail when the parseCardString fails")
	func() {
		msg := randomString(10, 20)

		fakeSplitter := &mockHandsStringSplitter{}
		fakeSplitter.On("splitHandsString", mock.AnythingOfType("string")).
			Return([]string{randomString(1, 3)}, nil).
			Once()

		fakeParser := &mockCardStringParser{}
		fakeParser.On("parseCardString", mock.AnythingOfType("string")).
			Return(Card{}, errors.New(msg)).
			Once()

		fakeHandsParser := handsParser{
			splitter:   fakeSplitter,
			cardParser: fakeParser,
		}

		_, _, err := fakeHandsParser.parseHands(randomString(1, 3))
		assert.Errorf(t, err, msg)
		fakeSplitter.AssertExpectations(t)
		fakeParser.AssertExpectations(t)
	}()
}
