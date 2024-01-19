/*3. The given program takes an integer value as input from the user i.e 1 or 2.
Option 1 represents Rectangle and Option 2 represents Square.
Given the metrics of the shapes are hard-coded, complete the given program to achieve the following:
1. Create an interface Quadrilateral which has the following method signatures
* Area() int
* Perimeter() int
2. Implement the Quadrilateral interface for the given shapes i.e Rectangle and Square
3. Define a "Print" function which accepts any shape that implements Quadrilateral interface and Prints Area and Perimeter of shape in the following manner:
"Area :  <value>"
"Perimeter :  <value>"
*/

package main

import (
	"fmt"
)

type Quadrilateral interface {
	Area() int
	Perimeter() int
}

// structs
type Rectangle struct {
	Length, Width int
}
type Square struct {
	Side int
}

func (r Rectangle) Area() int {
	return r.Length * r.Width
}
func (r Rectangle) Perimeter() int {
	return 2 * (r.Length + r.Width)
}

func (s Square) Area() int {
	return s.Side * s.Side
}
func (s Square) Perimeter() int {
	return 4 * s.Side
}

func Print(q Quadrilateral) {
	fmt.Println("Area:", q.Area())
	fmt.Println("Perimeter:", q.Perimeter())
}

func main() {
	var option, length, width, side int

	fmt.Print("1 - Rectangle\n2 - Square\nInput: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		fmt.Print("Enter length of rectangle: ")
		fmt.Scan(&length)
		fmt.Print("Enter width of rectangle: ")
		fmt.Scan(&width)

		calculationOfRectangle := Rectangle{Length: length, Width: width}
		Print(calculationOfRectangle)

	case 2:
		fmt.Print("Enter side of square: ")
		fmt.Scan(&side)

		calculationOfSquare := Square{Side: side}
		Print(calculationOfSquare)
	default:
		fmt.Println("Invalid option")
	}
}
