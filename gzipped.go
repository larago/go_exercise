package main 

import (
	"fmt"
	"bufio"
	"os"
	"compress/gzip"
)

func main() {
	fName := "MyFile.gz"
	var r *bufio.Reader 
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Println("Error occured", err)
		os.Exit(1)
	}
	fz ,err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("done")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}