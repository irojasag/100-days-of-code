package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello! Lets create a rock band name!")
	reader := bufio.NewReader(os.Stdin)

	var color string
	var pet string
	
	fmt.Print("Enter your favorite color: ")
	colorInput, _ := reader.ReadString('\n')
	color = strings.TrimSpace(colorInput)

	fmt.Print("Enter your pet's name: ")
	petInput, _ := reader.ReadString('\n')
	pet = strings.TrimSpace(petInput)

	fmt.Printf("Your rock band name could be: %s %s!\n", color, pet)
}