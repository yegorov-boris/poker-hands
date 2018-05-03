package main

import (
	"testing"
	"log"
	"github.com/stretchr/testify/assert"
	"math/rand"
)

type ScannerSpy struct {
	C int
	ShouldFail bool
}

func (s *ScannerSpy) Scan() bool {
	s.C++
	return s.C < 4
}

func (s *ScannerSpy) Text() string {
	dummy := []string{
		"4D 6S 9H QH QC 3D 6D 7H QD QS",
		"2D 9C AS AH AC 3D 6D 7D TD QD",
		"2H 2D 4C 4D 4S 3C 3D 3S 9S 9D",
	}
	if s.ShouldFail {
		return RandomString(20, 50)
	}
	return dummy[s.C - 1]
}

func TestCreateCheckers(t *testing.T) {
	log.Println("CreateCheckers")

	log.Println("should create goroutines that take a string and return an error when something fails")
	func () {
		inputs, outputs := CreateCheckers(MaxChunkSize)
		i := rand.Intn(MaxChunkSize)

		inputs[i] <- RandomString(20, 50)
		result := <- outputs[i]
		assert.Error(t, result.Right)
	}()

	log.Println("should create goroutines that take a string and return true when the first player wins")
	func () {
		inputs, outputs := CreateCheckers(MaxChunkSize)
		i := rand.Intn(MaxChunkSize)

		inputs[i] <- "5D 8C 9S JS AC 2C 5C 7D 8S QH"
		result := <- outputs[i]
		assert.Nil(t, result.Right)
		assert.True(t, result.Left)
	}()

	log.Println("should create goroutines that take a string and return false when the second player wins")
	func () {
		inputs, outputs := CreateCheckers(MaxChunkSize)
		i := rand.Intn(MaxChunkSize)

		inputs[i] <- "5H 5C 6S 7S KD 2C 3S 8S 8D TD"
		result := <- outputs[i]
		assert.Nil(t, result.Right)
		assert.False(t, result.Left)
	}()
}

func TestCountWins(t *testing.T) {
	log.Println("CountWins")

	log.Println("should fail when a checker fails")
	func () {
		scannerSpy := ScannerSpy{0, true}
		result, err := CountWins(&scannerSpy, 2)
		assert.Equal(t, 0, result)
		assert.Error(t, err)
	}()

	log.Println("should count the first player's wins")
	func () {
		scannerSpy := ScannerSpy{0, false}
		result, err := CountWins(&scannerSpy, 2)
		assert.Nil(t, err)
		assert.Equal(t, 2, result)
	}()
}
