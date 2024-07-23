package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	//var ptr *int
	//fmt.Println("Value of ther pointer is ", ptr)
	myNumber := 46

	var ptr = &myNumber
	fmt.Println("Value of the actual pointer is", ptr)
	fmt.Println("Value of the actual pointer is", *ptr)

	*ptr = *ptr - 2
	fmt.Println("New value is : ", myNumber)

}
