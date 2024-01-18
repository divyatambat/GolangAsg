package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	const PI float64 = 3.14

	fmt.Println("Enter radius: ")
	fmt.Scanln(&r)

	area := PI * math.Pow(r, 2)
	updatedArea := math.Round(area * 100) / 100
	fmt.Println("Area of circle: ", updatedArea)
}
