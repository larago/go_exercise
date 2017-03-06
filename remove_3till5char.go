package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("goprogram")
	if err != nil {
		fmt.Println(err)
		return
	}
	outputFile, err := os.OpenFile("goProgramT", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			outputWriter.Flush()
			fmt.Println("Conversion done.")
			fmt.Println("EOF")
			return
		}
		outputString := string(inputString[2:5]) + "\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
