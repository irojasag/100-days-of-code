package main

import (
	"fmt"
	
)

func main() {
	fmt.Println("Hello! Did you enjoy the dinner?")
	
	var total float64 = 0.0
	var tip float64 = 0.0

	fmt.Println("Time to pay the bill. The total is: ")

	fmt.Scanf("%f\n", &total)

	fmt.Print("The recommended tip is 15%, which is: ")
	fmt.Printf("%.2f\n", total*0.15)
	fmt.Println("How much would you like to tip?")

	fmt.Scanf("%f\n", &tip)

	fmt.Printf("The total amount to be paid is: %.2f\n", total+tip)
}