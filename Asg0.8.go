package main

import (
	"fmt"
	"sync"
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

	wg.Add(1)
	go func() {
		defer wg.Done()	
		nIsEven := isEven(n, &mu)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
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
