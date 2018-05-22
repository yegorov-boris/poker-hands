package main

type config struct {
	cardValues string
	suits      string
	separator  string
}

const MaxChunkSize = 10
const Url = "https://projecteuler.net/project/resources/p054_poker.txt"
const CardValues = "A K Q J T 9 8 7 6 5 4 3 2"
const Suits = "D C H S"
const Separator = " "

func defaultConfig() config {
	return config{
		cardValues: CardValues,
		suits:      Suits,
		separator:  Separator,
	}
}
