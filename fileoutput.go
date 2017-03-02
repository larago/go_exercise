package main 

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY | os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Println("Error occured")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "Hello!"

	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
	fmt.Println("Done")
}
