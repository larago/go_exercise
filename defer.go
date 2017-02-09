package main

import "fmt"

func main() {
	function1()
}

func function1() {
	fmt.Println("In function1 at the top")
	defer function2()
	fmt.Println("In function1 at the bottom")
}

func function2() {
	fmt.Println("function2: Deferred until the end of the calling function")
}
