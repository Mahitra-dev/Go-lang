package main

import "fmt"

func main() {
	fmt.Println("Welcome to the study of arrays in go lang")
	var fruitlist = [4]string{"Apple" , "Tomato" , "Peach"}
    fmt.Println("Fruit list is: ", fruitlist)
	fmt.Println("Number of fruits is: ", len(fruitlist))
 
	//There is a long way for this too..
	
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	fmt.Println("fruit list is: ", fruitList)
	fmt.Println("Number of fruits is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is : ", len(vegList))
}
