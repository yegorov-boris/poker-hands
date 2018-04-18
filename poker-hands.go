package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

func main() {
	resp, errGet := http.Get("https://projecteuler.net/project/resources/p054_poker.txt")
	if errGet != nil {
		fmt.Print(errGet)
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, errReadAll := ioutil.ReadAll(resp.Body)
		if errReadAll != nil {
			fmt.Print(errReadAll)
			os.Exit(1)
		}

		fmt.Print(string(bodyBytes))
	}
}
