package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"testing"
)

type scannerMock struct {
	c                int
	fakeHandsStrings []string
}

func (s *scannerMock) Scan() bool {
	s.c++
	return s.c < 4
}

func (s *scannerMock) Text() string {
	return s.fakeHandsStrings[s.c-1]
}

func TestCounter(t *testing.T) {
	log.Println("counter")

	log.Println("should send an error when it fails to process a hands string")
	func() {
		requestsToRead := make(chan Either, 2)
		counts := make(chan int, 2)
		msg := RandomString(10, 20)
		fakeCmp := &mockComparator{}
		fakeCmp.
			On("IsFirstPlayerWinner", mock.AnythingOfType("string")).
			Return(false, errors.New(msg)).
			Once()
		go counter(fakeCmp, requestsToRead, counts)

		initialRequestToRead := <-requestsToRead
		initialRequestToRead.Left <- RandomString(10, 20)
		finalRequestToRead := <-requestsToRead
		finalRequestToRead.Left <- stop

		assert.Nil(t, initialRequestToRead.Right)
		assert.Errorf(t, finalRequestToRead.Right, msg)
		fakeCmp.AssertExpectations(t)
	}()

	log.Println("should send the number of the processed hands strings")
	func() {
		requestsToRead := make(chan Either, 2)
		counts := make(chan int, 2)
		handsWin := RandomString(10, 20)
		handsLoss := RandomString(10, 20)
		fakeCmp := &mockComparator{}
		fakeCmp.
			On("IsFirstPlayerWinner", handsWin).
			Return(true, nil).
			Once().
			On("IsFirstPlayerWinner", handsLoss).
			Return(false, nil).
			Once()
		go counter(fakeCmp, requestsToRead, counts)

		initialRequestToRead := <-requestsToRead
		initialRequestToRead.Left <- handsLoss
		nextRequestToRead := <-requestsToRead
		nextRequestToRead.Left <- handsWin
		finalRequestToRead := <-requestsToRead
		finalRequestToRead.Left <- stop
		count := <-counts
		assert.Nil(t, initialRequestToRead.Right)
		assert.Nil(t, nextRequestToRead.Right)
		assert.Nil(t, finalRequestToRead.Right)
		assert.Equal(t, 1, count)
		fakeCmp.AssertExpectations(t)
	}()
}

func TestCountWins(t *testing.T) {
	log.Println("CountWins")

	log.Println("should fail when it fails to process a hands string")
	func() {
		fakeHandsString := RandomString(10, 15)
		fakeScanner := &mockScanner{}
		fakeScanner.On("Scan").Return(true).Twice()
		fakeScanner.On("Text").Return(fakeHandsString).Twice()

		msg := RandomString(10, 20)
		fakeCmp := &mockComparator{}
		fakeCmp.
			On("IsFirstPlayerWinner", fakeHandsString).
			Return(false, errors.New(msg)).
			Maybe()

		result, err := countWins(fakeScanner, fakeCmp, 2)
		assert.Equal(t, 0, result)
		assert.Errorf(t, err, msg)
		fakeScanner.AssertExpectations(t)
		fakeCmp.AssertExpectations(t)
	}()

	log.Println("should count the first player's wins")
	func() {
		fakeHandsStrings := []string{
			RandomString(10, 15),
			RandomString(10, 15),
			RandomString(10, 15),
		}

		fakeCmp := &mockComparator{}
		fakeCmp.
			On("IsFirstPlayerWinner", fakeHandsStrings[0]).
			Return(true, nil).
			Once().
			On("IsFirstPlayerWinner", fakeHandsStrings[1]).
			Return(false, nil).
			Once().
			On("IsFirstPlayerWinner", fakeHandsStrings[2]).
			Return(true, nil).
			Once()

		result, err := countWins(&scannerMock{fakeHandsStrings: fakeHandsStrings}, fakeCmp, 2)
		assert.Equal(t, 2, result)
		assert.Nil(t, err)
		fakeCmp.AssertExpectations(t)
	}()
}
