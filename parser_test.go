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
	log.Println("SplitHandsString")

	log.Println("should split a string by a separator")
	func() {
		expected := []string{}
		separator := RandomString(1, 3)
		for i := 0; i < 10; i++ {
			expected = append(expected, RandomStringWithout(1, 3, separator))
		}

		fakeConfig := config{separator: separator}
		actual, err := (splitter{config: fakeConfig}).SplitHandsString(strings.Join(expected, separator))
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	}()

	log.Println("should fail when the number of the substrings doesn't equal 10")
	func() {
		separator := RandomString(1, 3)
		fakeConfig := config{separator: separator}
		_, err := (splitter{config: fakeConfig}).SplitHandsString(RandomString(1, 3))
		assert.Errorf(t, err, "failed to parse a line with hands: wrong length")
	}()
}

func TestParseCardString(t *testing.T) {
	log.Println("ParseCardString")

	log.Println("should parse a valid card string")
	func() {
		suit := RandomString(1, 1)
		cardValue := RandomString(1, 1)
		expected := Card{Suit: suit, Value: cardValue}
		cardString := strings.Join([]string{cardValue, suit}, "")
		fakeConfig := config{
			suits:      suit,
			cardValues: cardValue,
		}

		actual, err := (cardParser{config: fakeConfig}).ParseCardString(cardString)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	}()

	log.Println("should fail when a card string is longer than 2 bytes")
	func() {
		_, err := (cardParser{config: defaultConfig()}).ParseCardString(RandomString(3, 5))
		assert.Errorf(t, err, "failed to parse an encoded card: wrong length")
	}()

	log.Println("should fail when a card string is shorter than 2 bytes")
	func() {
		_, err := (cardParser{config: defaultConfig()}).ParseCardString(RandomString(0, 1))
		assert.Errorf(t, err, "failed to parse an encoded card: wrong length")
	}()

	log.Println("should fail when a card string has invalid suit")
	func() {
		suit := RandomString(1, 1)
		cardValue := RandomString(1, 1)
		cardString := strings.Join([]string{cardValue, suit}, "")
		fakeConfig := config{
			suits:      RandomStringWithout(1, 1, suit),
			cardValues: cardValue,
		}

		_, err := (cardParser{config: fakeConfig}).ParseCardString(cardString)
		assert.Errorf(t, err, "failed to parse an encoded card: wrong suit")
	}()

	log.Println("should fail when a card string has invalid card value")
	func() {
		suit := RandomString(1, 1)
		cardValue := RandomString(1, 1)
		cardString := strings.Join([]string{cardValue, suit}, "")
		fakeConfig := config{
			suits:      suit,
			cardValues: RandomStringWithout(1, 1, cardString),
		}

		_, err := (cardParser{config: fakeConfig}).ParseCardString(cardString)
		assert.Errorf(t, err, "failed to parse an encoded card: wrong card value")
	}()
}

func TestSortByValue(t *testing.T) {
	log.Println("SortByValue")

	log.Println("should sort a hand by card values")

	separator := RandomString(1, 1)
	cardValues := []string{}
	for i := 0; i < 5; i++ {
		cardValues = append(cardValues, RandomStringWithout(1, 1, separator))
	}

	unsortedHand := Hand{}
	for i, index := range rand.Perm(5) {
		unsortedHand[i] = Card{
			Suit:  RandomString(1, 1),
			Value: cardValues[index],
		}
	}

	fakeConfig := config{
		cardValues: strings.Join(cardValues, separator),
		separator:  separator,
	}

	for i, card := range (sorter{config: fakeConfig}).SortByValue(unsortedHand) {
		assert.Equal(t, cardValues[i], card.Value)
	}
}

func TestParseHands(t *testing.T) {
	log.Println("ParseHands")

	log.Println("should parse a string which contains two hands")
	func() {
		fakeCardStrings := strings.Split("0123456789", "")
		fakeCard := Card{Value: "1", Suit: "a"}
		fakeHand := Hand{fakeCard, fakeCard, fakeCard, fakeCard, fakeCard}

		fakeSplitter := &mockHandsStringSplitter{}
		fakeSplitter.On("SplitHandsString", mock.AnythingOfType("string")).
			Return(fakeCardStrings, nil).
			Once()

		fakeParser := &mockCardStringParser{}
		fakeParser.On("ParseCardString", mock.AnythingOfType("string")).
			Return(fakeCard, nil).
			Times(10)

		fakeSorter := &mockSorterByValue{}
		fakeSorter.On("SortByValue", fakeHand).
			Return(fakeHand).
			Twice()

		fakeHandsParser := handsParser{
			splitter:   fakeSplitter,
			cardParser: fakeParser,
			sorter:     fakeSorter,
		}

		first, second, err := fakeHandsParser.ParseHands(RandomString(1, 3))
		assert.Nil(t, err)
		assert.Equal(t, fakeHand, first)
		assert.Equal(t, fakeHand, second)
		fakeSplitter.AssertExpectations(t)
		fakeParser.AssertExpectations(t)
		fakeSorter.AssertExpectations(t)
	}()

	log.Println("should fail when the SplitHandsString fails")
	func() {
		msg := RandomString(10, 20)

		fakeSplitter := &mockHandsStringSplitter{}
		fakeSplitter.On("SplitHandsString", mock.AnythingOfType("string")).
			Return([]string{}, errors.New(msg)).
			Once()

		_, _, err := (handsParser{splitter: fakeSplitter}).ParseHands(RandomString(1, 3))
		assert.Errorf(t, err, msg)
		fakeSplitter.AssertExpectations(t)
	}()

	log.Println("should fail when the cards string contains duplicated encoded cards")
	func() {
		fakeCardString := RandomString(1, 3)
		fakeHandsString := strings.Join([]string{fakeCardString, fakeCardString}, " ")

		fakeSplitter := &mockHandsStringSplitter{}
		fakeSplitter.On("SplitHandsString", fakeHandsString).
			Return([]string{fakeCardString}, nil).
			Once()

		_, _, err := (handsParser{splitter: fakeSplitter}).ParseHands(fakeHandsString)
		assert.Errorf(t, err, "failed to parse a line with hands: %s is not unique", fakeCardString)
		fakeSplitter.AssertExpectations(t)
	}()

	log.Println("should fail when the ParseCardString fails")
	func() {
		msg := RandomString(10, 20)

		fakeSplitter := &mockHandsStringSplitter{}
		fakeSplitter.On("SplitHandsString", mock.AnythingOfType("string")).
			Return([]string{RandomString(1, 3)}, nil).
			Once()

		fakeParser := &mockCardStringParser{}
		fakeParser.On("ParseCardString", mock.AnythingOfType("string")).
			Return(Card{}, errors.New(msg)).
			Once()

		fakeHandsParser := handsParser{
			splitter:   fakeSplitter,
			cardParser: fakeParser,
		}

		_, _, err := fakeHandsParser.ParseHands(RandomString(1, 3))
		assert.Errorf(t, err, msg)
		fakeSplitter.AssertExpectations(t)
		fakeParser.AssertExpectations(t)
	}()
}
