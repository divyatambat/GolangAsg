package main

import (
	"fmt"
	"sync"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		nIsEven := isEven(n)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
		mu.Unlock()
	}()
	//wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		n++
		mu.Unlock()
	}()
	wg.Wait()
}
