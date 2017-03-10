package main

import (
    "fmt"
    "os/exec"
    "sync"
    "strings"
)

func exe_cmd(cmd string, wg *sync.WaitGroup) {
    fmt.Println(cmd)
            parts := strings.Fields(cmd)
    out, err := exec.Command(parts[0],parts[1]).Output()
    if err != nil {
        fmt.Println("error occured")
        fmt.Printf("%s", err)
    }
    fmt.Printf("%s", out)
    wg.Done()
}

func main() {
    wg := new(sync.WaitGroup)
    commands := []string{"grep ERROR"}
    for _, str := range commands {
        wg.Add(1)
        go exe_cmd(str, wg)
    }
    wg.Wait()
}