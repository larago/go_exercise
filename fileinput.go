package main 

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Println("Error occured.")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return 
		}
		fmt.Printf("The input was: %s", inputString)
	}
}