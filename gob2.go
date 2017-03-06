package main 

import (
	"encoding/gob"
	"log"
	"os"
)

type Address struct {
	Type 		string 
	City		string 
	Country 	string 
}

type VCard struct {
	FirstName 	string 
	lastName	string 
	Addresses	[]*Address
	Remark		string 
}

var content string 

func main() {
	pa := &Address{"private", "London", "England"}
	wa := &Address{"work", "GuangZhou", "China"}
	vc := VCard{"John", "Peter", []*Address{pa, wa}, "none"}
	
	file, err := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	err = enc.Encode(vc)
	if err != nil {
		log.Fatal("Error in encoding gob")
	}
}