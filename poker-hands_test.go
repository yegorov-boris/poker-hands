package main

import (
	"testing"
	"log"
	"github.com/stretchr/testify/assert"
	"math/rand"
)

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
