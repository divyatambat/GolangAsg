package main
import "fmt"

func main(){
	var input string
	fmt.Println("Input: ")
	fmt.Scanln(&input)
	
	output := romenToInteger(input)
	fmt.Println("Output: ", output)
}

func romenToInteger(input string) int {
	output := 0
	for i := 0 ; i < len(input) ; i++ {
		switch input[i] {
			case 'I':
				if i < len(input)-1 && (input[i+1] == 'V' || input[i+1] == 'X') {
					output -= 1
				} else {
					output += 1
				}
			case 'V':
				output += 5
			case 'X':
				if i < len(input)-1 && (input[i+1] == 'L' || input[i+1] == 'C') {
					output -= 10
				} else {
					output += 10
				}
			case 'L':
				output += 50
			case 'C':
				if i < len(input)-1 && (input[i+1] == 'D' || input[i+1] == 'M') {
					output -= 100
				} else {
					output += 100
				}
			case 'D':
				output += 500
			case 'M':
				output += 1000
			default:
				fmt.Println("Invalid input!")
		}
	}
	
	return output;
}
