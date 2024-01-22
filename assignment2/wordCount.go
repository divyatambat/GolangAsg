package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a string: ")
	scanner.Scan()
	input := scanner.Text()

	words := strings.Fields(input)

	wordFreq := make(map[string]int)
	maxFreq := 0
	
	var result []string

	for _, word := range words {
		wordFreq[word]++
		if wordFreq[word] > maxFreq {
			maxFreq = wordFreq[word]
			result = []string{word}
		} else if wordFreq[word] == maxFreq {
			result = append(result, word)
		}
	}

	fmt.Println("Output: ", result)
}

