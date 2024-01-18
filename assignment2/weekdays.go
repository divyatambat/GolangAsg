package main

import (
	"fmt"
)

func main() {

	daysMap := map[int]string {
		0 : "Monday",
		1 : "Tuesday",
		2 : "Wednesday",
		3 : "Thusday",
		4 : "Friday",
		5 : "Saturday",
		6 : "Sunday",
	}	
	
	var index int
	
	fmt.Println("Input: ")
	fmt.Scanln(&index)
	
	if day := daysMap[index]; day != "" {
		fmt.Println("Output: ", day)
	} else {
		fmt.Println("Output: Not a day!")
	}
	
	
}

