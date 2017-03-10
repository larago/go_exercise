package main

import (
	"os"
	"log"
	"os/exec"
	"fmt"
	"strings"
	"bytes"
	"regexp"
	"encoding/json"
	"strconv"
)

type LineData struct {
	Total 		int
	Content		map[string]*LineContent
}

type LineContent struct {
	Value		string 
	Level		string
}



// func LogPath() string {
// 	file, err := exec.LookPath(os.Args[0])
// 	if err != nil {
// 		Log(err)
// 	}
// 	path, err := filepath.Abs(file)
// 	if err != nil {
// 		Log(err)
// 	}
// 	execPath := strings.Split(path, "/")
// 	pathLength := len(execPath)
// 	SeabedPath := strings.Join(execPath[:pathLength-2], "/")
// 	SeabedLogPath := SeabedPath + "/" + "seabed.log"
// 	return SeabedLogPath
// }


func ReadFile(filePath, line string, logLevel int) []byte {

	var output bytes.Buffer
	cmd := exec.Command("tail", filePath, "-n", line)
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		Log(err)
	}
	eachLine := strings.Split(output.String(), "\n")
	lineJSON := &LineData{Content: make(map[string]*LineContent),
		Total: logLevel}

	pats := []string{"<ERROR>", "<INFO>", "<WARNING>", 
		"<DEBUG>", "<CRITICAL>", "<NOTSET>"}
	levelNum := make(map[string]int)
	levelNum["<NOTSET>"] = 0
	levelNum["<DEBUG>"] = 10
	levelNum["<INFO>"] = 20
	levelNum["<WARNING>"] = 30
	levelNum["<ERROR>"] = 40
	levelNum["<CRITICAL>"] = 50

	for i, each := range eachLine {
		for _, pat := range pats {
			if ok, _ := regexp.MatchString(pat, each); ok {
				lineLevel := levelNum[pat]
				if lineLevel >= logLevel {
					lineJSON.Content[string(i)] = &LineContent{Value: each, Level: pat}
				}
				break
			}
		}
	}
	js, err := json.Marshal(lineJSON)
	if err != nil {
		Log(err)
	}
	return js
}

func Log(msg error) {
	log.Fatal(msg)
	panic(msg)
}

func ParseCommand() (logFile, lineNum, logLevel string) {
	logFile = os.Args[1]
	lineNum = os.Args[2]
	logLevel = os.Args[3]
	_, err := exec.LookPath(logFile)
	if err != nil {
		Log(err)
	}
	return
}

func main() {
	if len(os.Args[1:]) != 4 {
		filePath, lineNum, logLevel := ParseCommand()
		level, err := strconv.Atoi(logLevel)
		if err != nil {
			Log(err)
		}
		result := ReadFile(filePath, lineNum, level)
		fmt.Printf("%s", result)
	} else {
		log.Fatal("Number of parameters incorrect.")
	}
}