package main

type config struct {
	cardValues string
	suits      string
	separator  string
}

const poolSzie = 10
const Url = "https://projecteuler.net/project/resources/p054_poker.txt"
const CardValues = "A K Q J T 9 8 7 6 5 4 3 2"
const Suits = "D C H S"
const Separator = " "
const stop = "stop"

func defaultConfig() config {
	return config{
		cardValues: CardValues,
		suits:      Suits,
		separator:  Separator,
	}
}

func defaultComparator() comparator {
	return cmp{
		config: defaultConfig(),
		parser: handsParser{
			splitter:   splitter{config: defaultConfig()},
			cardParser: cardParser{config: defaultConfig()},
			sorter:     sorter{config: defaultConfig()},
		},
		matcher: combMatcher{config: defaultConfig()},
	}
}
