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

type LineContent struct {
	Value	string 	`json:"value"`
	Level	string	`json:"level"`
}


func ReadFile(filePath, line string, logLevel, levelFilter int) string {

	var output bytes.Buffer
	// var params []string = make([]string, 0)
	// levelString := strconv.Itoa(logLevel)
	var cmd *exec.Cmd
	if levelFilter == -1 {
		params := []string{"tail", filePath, "-n", line}
		paramString := strings.Join(params, " ")
		cmd = exec.Command("bash", "-c", paramString)
	} else {
		aboveLevel := strconv.Itoa(levelFilter)
		levelMap := make(map[string]string)
		levelMap["0"] = "'CRITICAL|ERROR|WARNING|INFO|DEBUG|NOTEST'"
		levelMap["10"] = "'CRITICAL|ERROR|WARNING|INFO|DEBUG'"
		levelMap["20"] = "'CRITICAL|ERROR|WARNING|INFO'"
		levelMap["30"] = "'CRITICAL|ERROR|WARNING'"
		levelMap["40"] = "'CRITICAL|ERROR'"
		levelMap["50"] = "'CRITICAL'"
		levelString := levelMap[aboveLevel]
		if len(levelString) == 0 {
			levelString = "no_level"
		}
		params := []string{"grep", "-E" , levelString, filePath, "|", "tail", "-n", line}
		paramString := strings.Join(params, " ")
		cmd = exec.Command("bash", "-c", paramString)
	}
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		Log(err)
	}
	eachLine := strings.Split(output.String(), "\n")
	
	var dataSlice []*LineContent
	var colour string

	if logLevel != -1 {
		pats := []string{"<ERROR>", "<INFO>", "<WARNING>", 
				"<DEBUG>", "<CRITICAL>", "<NOTSET>"}
		levelNum := make(map[string]int)
		levelNum["<NOTSET>"] = 0
		levelNum["<DEBUG>"] = 10
		levelNum["<INFO>"] = 20
		levelNum["<WARNING>"] = 30
		levelNum["<ERROR>"] = 40
		levelNum["<CRITICAL>"] = 50
		for _, each := range eachLine {
			found := false
			for _, pat := range pats {
				if ok, _ := regexp.MatchString(pat, each); ok {
					found = true
					lineLevel := levelNum[pat]
					if lineLevel == 30 {
						colour = "yellow"  
					} else if lineLevel > 30 {
						colour = "red"
					} else {
						colour = "white"
					}
					dataSlice = append(dataSlice, 
									   &LineContent{Value: each, 
								       Level: colour})
					break
				}
			}
			if found == false && len(each) != 0 {
				dataSlice = append(dataSlice, 
					&LineContent{Value: each, 
					Level: "no_level"})
			}
		}
	} else {
		for _, each := range eachLine {
			dataSlice = append(dataSlice, 
							   &LineContent{Value: each, 
							   Level: "no_level"})
		}
	}
	js, err := JSONMarshal(dataSlice,true)
	if err != nil {
		Log(err)
	}
	return string(js)
}

func JSONMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)

	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
        b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
        b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}

func Log(msg error) {
	log.Fatal(msg)
	panic(msg)
}

func ParseCommand() (logFile, lineNum, logLevel, levelFilter string) {
	logFile = os.Args[1]
	lineNum = os.Args[2]
	logLevel = os.Args[3]
	levelFilter = os.Args[4]
	_, err := exec.LookPath(logFile)
	if err != nil {
		Log(err)
	}
	return
}

func main() {
	if len(os.Args[1:]) == 4 {
		filePath, lineNum, logLevel, levelFilter := ParseCommand()
		level, err := strconv.Atoi(logLevel)
		if err != nil {
			Log(err)
		}
		filter, err := strconv.Atoi(levelFilter)
		if  err != nil {
			Log(err)
		}
		result := ReadFile(filePath, lineNum, level, filter)
		fmt.Printf("%s", result)
	} else {
		log.Fatal("Number of parameters incorrect.")
	}
}