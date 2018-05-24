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

var checkerNames = []string{
	"IsRoyalFlush",
	"IsStraightFlush",
	"IsFourKind",
	"IsFullHouse",
	"IsFlush",
	"IsStraight",
	"IsThreeKind",
	"IsTwoPairs",
	"IsOnePair",
}

func TestGetCombination(t *testing.T) {
	log.Println("GetCombination")

	log.Println("should return the combination index and the reordered hand")
	func() {
		expectedIndex := rand.Intn(len(checkerNames))
		fakeMatcher := &mockCombinationMatcher{}
		for i, checkerName := range checkerNames {
			fakeMatcher.On(checkerName, Hand{}).Return(i == expectedIndex, Hand{}).Maybe()
		}

		actualIndex, actualHand := (cmp{matcher: fakeMatcher}).GetCombination(Hand{})
		assert.Equal(t, expectedIndex, actualIndex)
		assert.Equal(t, Hand{}, actualHand)
		fakeMatcher.AssertExpectations(t)
	}()

	log.Println("should return 9 and the hand when there's no combination")
	func() {
		fakeMatcher := &mockCombinationMatcher{}
		for _, checkerName := range checkerNames {
			fakeMatcher.On(checkerName, Hand{}).Return(false, Hand{}).Once()
		}

		actualIndex, actualHand := (cmp{matcher: fakeMatcher}).GetCombination(Hand{})
		assert.Equal(t, 9, actualIndex)
		assert.Equal(t, Hand{}, actualHand)
		fakeMatcher.AssertExpectations(t)
	}()
}

func TestIsFirstPlayerWinner(t *testing.T) {
	log.Println("IsFirstPlayerWinner")

	log.Println("should fail when the ParseHands fails")
	func() {
		msg := RandomString(10, 20)
		fakeParser := &mockHandsStringParser{}
		fakeParser.On("ParseHands", mock.AnythingOfType("string")).
			Return(Hand{}, Hand{}, errors.New(msg)).
			Once()

		_, err := (cmp{parser: fakeParser}).IsFirstPlayerWinner(RandomString(30, 50))
		assert.Errorf(t, err, msg)
		fakeParser.AssertExpectations(t)
	}()

	log.Println("should return true when the first player's combination is stronger")
	func() {
		firstHand := Hand{Card{Value: "first"}}
		secondHand := Hand{Card{Value: "second"}}
		fakeParser := &mockHandsStringParser{}
		fakeParser.On("ParseHands", mock.AnythingOfType("string")).
			Return(firstHand, secondHand, nil).
			Once()

		secondCombinationIndex := 1 + rand.Intn(len(checkerNames))
		firstCombinationIndex := rand.Intn(secondCombinationIndex)
		fakeMatcher := &mockCombinationMatcher{}
		for i, checkerName := range checkerNames {
			fakeMatcher.
				On(checkerName, firstHand).
				Return(i == firstCombinationIndex, Hand{}).
				Maybe().
				On(checkerName, secondHand).
				Return(i == secondCombinationIndex, Hand{}).
				Maybe()
		}

		fakeComparator := cmp{
			parser:  fakeParser,
			matcher: fakeMatcher,
		}

		actual, err := fakeComparator.IsFirstPlayerWinner(RandomString(20, 30))
		assert.Nil(t, err)
		assert.True(t, actual)
		fakeParser.AssertExpectations(t)
		fakeMatcher.AssertExpectations(t)
	}()

	log.Println("should return false when the second player's combination is stronger")
	func() {
		firstHand := Hand{Card{Value: "first"}}
		secondHand := Hand{Card{Value: "second"}}
		fakeParser := &mockHandsStringParser{}
		fakeParser.On("ParseHands", mock.AnythingOfType("string")).
			Return(firstHand, secondHand, nil).
			Once()

		firstCombinationIndex := 1 + rand.Intn(len(checkerNames))
		secondCombinationIndex := rand.Intn(firstCombinationIndex)
		fakeMatcher := &mockCombinationMatcher{}
		for i, checkerName := range checkerNames {
			fakeMatcher.
				On(checkerName, firstHand).
				Return(i == firstCombinationIndex, Hand{}).
				Maybe().
				On(checkerName, secondHand).
				Return(i == secondCombinationIndex, Hand{}).
				Maybe()
		}

		fakeComparator := cmp{
			parser:  fakeParser,
			matcher: fakeMatcher,
		}

		actual, err := fakeComparator.IsFirstPlayerWinner(RandomString(20, 30))
		assert.Nil(t, err)
		assert.False(t, actual)
		fakeParser.AssertExpectations(t)
		fakeMatcher.AssertExpectations(t)
	}()

	log.Println("should return true when the combinations have equal ranks ",
		"but the first player's cards are stronger")
	func() {
		firstValue := RandomStringWithout(1, 1, " ")
		secondValue := RandomStringWithout(1, 1, strings.Join([]string{" ", firstValue}, ""))
		position := rand.Intn(5)
		firstHand := Hand{}
		firstHand[position].Value = firstValue
		secondHand := Hand{}
		secondHand[position].Value = secondValue
		fakeParser := &mockHandsStringParser{}
		fakeParser.On("ParseHands", mock.AnythingOfType("string")).
			Return(firstHand, secondHand, nil).
			Once()

		index := rand.Intn(len(checkerNames) + 1)
		fakeMatcher := &mockCombinationMatcher{}
		for i, checkerName := range checkerNames {
			fakeMatcher.
				On(checkerName, firstHand).
				Return(i == index, firstHand).
				Maybe().
				On(checkerName, secondHand).
				Return(i == index, secondHand).
				Maybe()
		}

		fakeCardValues := strings.Join([]string{firstValue, secondValue}, " ")
		fakeComparator := cmp{
			config:  config{cardValues: fakeCardValues},
			parser:  fakeParser,
			matcher: fakeMatcher,
		}

		actual, err := fakeComparator.IsFirstPlayerWinner(RandomString(20, 30))
		assert.Nil(t, err)
		assert.True(t, actual)
		fakeParser.AssertExpectations(t)
		fakeMatcher.AssertExpectations(t)
	}()

	log.Println("should return false when the combinations have equal ranks ",
		"but the first player's cards are not stronger")
	func() {
		firstValue := RandomStringWithout(1, 1, " ")
		secondValue := RandomStringWithout(1, 1, strings.Join([]string{" ", firstValue}, ""))
		position := rand.Intn(5)
		firstHand := Hand{}
		firstHand[position].Value = firstValue
		secondHand := Hand{}
		secondHand[position].Value = secondValue
		fakeParser := &mockHandsStringParser{}
		fakeParser.On("ParseHands", mock.AnythingOfType("string")).
			Return(firstHand, secondHand, nil).
			Once()

		index := rand.Intn(len(checkerNames) + 1)
		fakeMatcher := &mockCombinationMatcher{}
		for i, checkerName := range checkerNames {
			fakeMatcher.
				On(checkerName, firstHand).
				Return(i == index, firstHand).
				Maybe().
				On(checkerName, secondHand).
				Return(i == index, secondHand).
				Maybe()
		}

		fakeCardValues := strings.Join([]string{secondValue, firstValue}, " ")
		fakeComparator := cmp{
			config:  config{cardValues: fakeCardValues},
			parser:  fakeParser,
			matcher: fakeMatcher,
		}

		actual, err := fakeComparator.IsFirstPlayerWinner(RandomString(20, 30))
		assert.Nil(t, err)
		assert.False(t, actual)
		fakeParser.AssertExpectations(t)
		fakeMatcher.AssertExpectations(t)
	}()
}
