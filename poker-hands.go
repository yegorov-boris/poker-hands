package main

import (
	"fmt"
	"net/http"
	"os"
	"bufio"
)

func main() {
	const chunkSize = 10

	resp, errGet := http.Get("https://projecteuler.net/project/resources/p054_poker.txt")
	if errGet != nil {
		fmt.Printf("Failed to download the combinations: %s\n", errGet)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to download the combinations. Status %d\n", resp.StatusCode);
		os.Exit(1)
	}

	defer resp.Body.Close()

	var i = 0;
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		i++
		if (i == chunkSize) {
			fmt.Println("------------")
			i = 0
		}
	}
}
