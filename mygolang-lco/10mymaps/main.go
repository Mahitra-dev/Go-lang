package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all the languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all the languages: ", languages)

	// loops are interesting in go lang

	//for key, value :=range languages {
	//fmt.Printf("For key %v, value is %v\n" , key, value)
	//For comma ok syntax
	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}
}
