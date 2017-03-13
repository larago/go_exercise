package main 

import "time"
import "fmt"

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(2 * 1e9)
		x := <- c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10 
	fmt.Println("sent", 10)
}