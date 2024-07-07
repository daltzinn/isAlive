package main

import (
    "fmt"
    "bufio"
    "os"
    "net/http"
    "strings"
)


func readFile() []string {
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Printf("Error opening the file `%s`\n%s\n", os.Args[1], err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text()) 
    }
    return lines
}


func usage() {
    if len(os.Args) >= 4 || len(os.Args) <= 1 {
        fmt.Println("Usage: ./isAlive <file> <output>")
        os.Exit(1)
    }
}


func main() {
    usage()
    line := readFile()
    for _, url := range line {

        var fullUrl string
        if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
            fullUrl = url
        } else {
            fullUrl = "https://"+url
        }
        req, err := http.Get(fullUrl)
        if err != nil {
            fmt.Println(fullUrl + "-> not up")
            continue
        }

        fmt.Println(fullUrl, "->", req.StatusCode)
        if len(os.Args) == 3 {
            f, err := os.OpenFile(os.Args[2], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
            if err != nil {
                fmt.Println(err)
            }

            if req.StatusCode != 404 {
                fmt.Fprintln(f, fullUrl, "->", req.StatusCode)
                defer f.Close()
            }
        }
    }
}

