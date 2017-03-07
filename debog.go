package main 

import (
	"bufio"
	"fmt"
	"encoding/gob"
	"log"
	"os"
)

type Address struct {
	Type 		string 
	City		string 
	Country		string 
}

type VCard struct {
	FirstName		string
	LastName 		string 
	Addresses		[]*Address 
	Remark 			string 
}

var content string 
var vc VCard 

func main() {

	file, err := os.Open("vcard.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	inReader := bufio.NewReader(file)
	
}