package main

import (
	"fmt"
)

func main() {
	array := []string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}
	
	var index1, index2 int
	fmt.Print("Enter two indexes [saperated by white space]: ")
	fmt.Scan(&index1, &index2)

	if index1<0 || index2 >= len(array) || index1 >= len(array) || index2<0 {
		fmt.Println("incorrect indexes")
	return
}

slice1 := array[:index1+1]
slice2 := array[index1:index2+1]
slice3 := array[index2:]

fmt.Println("Output: ")
fmt.Println(slice1)
fmt.Println(slice2)
fmt.Println(slice3)
}
