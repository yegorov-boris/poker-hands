package main

import (
	"testing"
	"log"
	"github.com/stretchr/testify/assert"
)

func TestCreateChecker(t *testing.T) {
	log.Println("CreateChecker")

	log.Println("should return a goroutine that takes a string and returns an error when something fails")
	func () {
		input := make(chan string)
		output := make(chan EitherBool)
		CreateChecker(input, output)

		input <- RandomString(20, 50)
		result := <- output
		assert.Error(t, result.Right)
	}()

	log.Println("should return a goroutine that takes a string and returns true when the first player wins")
	func () {
		input := make(chan string)
		output := make(chan EitherBool)
		CreateChecker(input, output)

		input <- "5D 8C 9S JS AC 2C 5C 7D 8S QH"
		result := <- output
		assert.Nil(t, result.Right)
		assert.True(t, result.Left)
	}()

	log.Println("should return a goroutine that takes a string and returns false when the second player wins")
	func () {
		input := make(chan string)
		output := make(chan EitherBool)
		CreateChecker(input, output)

		input <- "5H 5C 6S 7S KD 2C 3S 8S 8D TD"
		result := <- output
		assert.Nil(t, result.Right)
		assert.False(t, result.Left)
	}()
}
