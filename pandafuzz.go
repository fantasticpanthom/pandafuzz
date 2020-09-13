package main

import (
    "flag"
    "fmt"
    "os"
    "log"
    "net/http"
    "bufio"
)

func main() {
    required := []string{"url", "wl"}
    url := flag.String("url", "", "Target URL")
    wordlist := flag.String("wl","","Wordlist file path")

    flag.Parse()

    seen := make(map[string]bool)
    flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
    for _, req := range required {
        if !seen[req] {
            fmt.Fprintf(os.Stderr, "missing required -%s argument/flag\n", req)
            os.Exit(2)
        }
    }

    
    file, err := os.Open(*wordlist)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()


    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        resp, err := http.Get(*url+"/"+scanner.Text())
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(*url+"/"+scanner.Text(),":", resp.StatusCode)
    }

    
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}