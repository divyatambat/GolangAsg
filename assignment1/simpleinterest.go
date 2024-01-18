package main

import (
	"fmt"
	"math"
)

func main() {
var principle, rate float64
var time int

fmt.Println("Enter principle amount: ")
fmt.Scanln(&principle)

fmt.Println("Enter interest rate: ")
fmt.Scanln(&rate)

fmt.Println("Enter time period: ")
fmt.Scanln(&time)

interest := (principle * rate * float64(time)) / 100
updatedinterest := math.Round(interest * 100) / 100

fmt.Println("Simple Interest: ", updatedinterest)
}

