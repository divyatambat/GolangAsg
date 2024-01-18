package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"	
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Enter a string: ")
	scanner.Scan()
	input := scanner.Text()
		
	words := strings.Fields(input)
	
	wordFreq := make(map[string]int)
	
	for _, word := range words {
		wordFreq[word]++
	}
	
	var maxFreq int
	for _, freq := range wordFreq {
		if freq > maxFreq {
			maxFreq = freq
		}
	}
	
	var result []string
	for word, freq := range wordFreq {
		if freq == maxFreq {
			result = append(result, word)
		}
	}
	
	fmt.Println("Output: ", result)
}
