package main

import "fmt"

func main() {
	var username string = "mahitra"
	fmt.Println("username")
	fmt.Printf("variable is of type :%T \n", username)
	//printing a boolean value
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of the type : %T \n", isLoggedIn)
	//printing an integeral value
	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("variable is of the type: %T \n", smallval)
}
