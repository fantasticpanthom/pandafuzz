package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	required := []string{"url", "wl"}
	url := flag.String("url", "", "Target URL")
	wordlist := flag.String("wl", "", "Wordlist file path")
	fuzzword := flag.String("fw", "fuzz", "fuzz keyword")
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

	jobs := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()

			for URL := range jobs {
				resp, err := http.Get(URL)
				if err != nil {
					continue
				}
				fmt.Println(URL, ":", resp.StatusCode)
			}
		}()
	}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		URL := strings.ReplaceAll(*url, *fuzzword, sc.Text())
		jobs <- URL
	}
	close(jobs)

	wg.Wait()
}
