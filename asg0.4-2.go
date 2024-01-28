//2. The given program accepts 2 values from the user, length and breadth of a rectangle respectively. 
//Complete the program by writing methods "Area" and "Perimeter" on Rectangle type to calculate the area and perimeter of a rectangle.


package main

import "fmt"

// structs
type RectangleAsg2 struct {
	Length, Width int
}

func (r RectangleAsg2) Area() int {
	return r.Length * r.Width
}
func (r RectangleAsg2) Perimeter() int {
	return 2 * (r.Length + r.Width)
}

func main() {
	var length, width int

	fmt.Print("Enter length of rectangle: ")
	fmt.Scan(&length)
	fmt.Print("Enter width of rectangle: ")
	fmt.Scan(&width)

	calculationOfRectangle := RectangleAsg2{Length: length, Width: width}

	fmt.Println("Area:", calculationOfRectangle.Area())
	fmt.Println("Perimeter:", calculationOfRectangle.Perimeter())
}
