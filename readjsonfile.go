package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("vcard.json")
	if err != nil {
		log.Println("open error:", err)
		return
	}
	jsonDecoder := json.NewDecoder(file)

	var jsonContent interface{}
	err = jsonDecoder.Decode(&jsonContent)
	if err != nil {
		log.Println("decode error:", err)
		return
	}
	fmt.Println(jsonContent)

	type b interface{}
	var bb b
	bb = {"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}
	// err = json.Unmarshal(b, &jsonContent)
	// if err != nil {
	//  log.Println("Unmarshal error:", err)
	//  return
	// }
}
