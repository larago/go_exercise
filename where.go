package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}

	fmt.Println("Start here")
	where()
	fmt.Println("Go on")
	where()

	log.SetFlags(log.Llongfile)
	log.Print("sdfsdf")

	var where2 = log.Print

	where2("hee")
}
