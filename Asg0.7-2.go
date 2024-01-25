package main

import (
	"fmt"
	"runtime"
	"sync"
)

func reverseString(s string) string {
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}

func main() {
	originalString := "test123"

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		revString := reverseString(originalString)
		fmt.Printf("Output string: %s - %d\n", revString, runtime.NumGoroutine())
	}()

	wg.Wait()
}
