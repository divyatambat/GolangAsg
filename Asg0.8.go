package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int, m *sync.Mutex) bool {
	m.Lock()
	result := n%2 == 0
	m.Unlock()
	return result
}

func main() {
	n := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	go func() {
		nIsEven := isEven(n, &mu)
		wg.Add(1)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		wg.Done()
		fmt.Println(n, "is odd")
	}()
	wg.Wait()

	go func() {
		mu.Lock()
		n++
		mu.Unlock()
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}
