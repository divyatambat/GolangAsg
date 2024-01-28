package main

import (
	"errors"
	"fmt"
)

func accessSlice(slice []int, index int) (int, error) {

	if index < 0 || index >= len(slice) {
		return 0, errors.New("Length of the slice should be more than index!")
	}
	return slice[index], nil
}

func main() {
	slice := []int{10, 20, 30, 40, 50}
	var index int
	fmt.Print("Input: ")
	fmt.Scanln(&index)

	value, err := accessSlice(slice, index)
	if err != nil {
		fmt.Printf("Internal error: %s\n", err)
		return
	} 
	fmt.Printf("Item: %d, Value: %d\n", index, value)

}
