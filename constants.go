package main

type config struct {
	cardValues string
	suits      string
	separator  string
}

const combinationsUrl = "https://projecteuler.net/project/resources/p054_poker.txt"
const poolSzie = 10
const stop = "stop"

func defaultConfig() config {
	return config{
		cardValues: "A K Q J T 9 8 7 6 5 4 3 2",
		suits:      "D C H S",
		separator:  " ",
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
