package main

import "fmt"

func accessSlice(slice []int, index int) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Internal error: %s\n", err)
		}
	}()

	value := slice[index]
	fmt.Printf("Item: %d, Value: %d\n", index, value)
}

func main() {
	slice := []int{10, 20, 30, 40, 50}
	var index int
	fmt.Print("Input: ")
	fmt.Scanln(&index)

	accessSlice(slice, index)
}
