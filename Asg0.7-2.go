package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	originalString := "test123"

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		runes := []rune(originalString)

		for currentIndex := 0; currentIndex < len(runes)/2; currentIndex++ {
			oppositeIndex := len(runes) - 1 - currentIndex
			runes[currentIndex], runes[oppositeIndex] = runes[oppositeIndex], runes[currentIndex]
		}

		reversedString := string(runes)
		fmt.Printf("Reversed string: %s - %d\n", reversedString, runtime.NumGoroutine()) // Print in a single line for desired output
	}()

	wg.Wait()
}
