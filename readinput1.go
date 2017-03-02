// Read from control platform 

package main 

import (
	"fmt"
)

var (
	firstName, lastName, s string 
	i int 
	f float64 
	input = "56.12 / 5212 / Go"
	format = "%f / %d / %s"
)

func main() {
	fmt.Println("Please enter your full name")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("from the string we read: ", f, i, s)
}