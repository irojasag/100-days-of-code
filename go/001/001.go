package main

import "fmt"

func main() {
	fmt.Println("Hello! Lets create a rock band name!")
	var color string
	var pet string
	
	fmt.Print("Enter your favorite color: ")
	fmt.Scan(&color)

	fmt.Print("Enter your pet's name: ")
	fmt.Scan(&pet)

	fmt.Printf("Your rock band name could be: %s %s!\n", color, pet)
}