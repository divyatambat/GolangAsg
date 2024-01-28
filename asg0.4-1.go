//Golang Assignment 0.4 - 1

package main

import (
	"fmt"
)

type Hello struct {
	Message string
}

func AcceptAnything(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("This is a value of type Integer, %d\n", v)
	case string:
		fmt.Printf("This is a value of type String, %s\n", v)
	case bool:
		fmt.Printf("This is a value of type Boolean, %t\n", v)
	case Hello:
		fmt.Printf("This is a value of type Hello, %s\n", v.Message)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	fmt.Println("Choose an option:")
	fmt.Println("1. Integer")
	fmt.Println("2. String")
	fmt.Println("3. Boolean")
	fmt.Println("4. Custom type Hello")

	var option int
	fmt.Print("Enter your choice (1-4): ")
	fmt.Scan(&option)

	switch option {
	case 1:
		var intValue int
		fmt.Print("Enter an integer value: ")
		fmt.Scan(&intValue)
		AcceptAnything(intValue)
	case 2:
		var stringValue string
		fmt.Print("Enter a string value: ")
		fmt.Scan(&stringValue)
		AcceptAnything(stringValue)
	case 3:
		var boolValue bool
		fmt.Print("Enter a boolean value (true/false): ")
		fmt.Scan(&boolValue)
		AcceptAnything(boolValue)
	case 4:
		var helloValue Hello
		fmt.Print("Enter a message for Hello type: ")
		fmt.Scan(&helloValue.Message)
		AcceptAnything(helloValue)
	default:
		fmt.Println("Invalid option")
	}
}
