package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
//:= is called walurus operator
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")
	// comma ok // err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating, ", input)
	fmt.Printf("type of this rating is %T", input)
}
