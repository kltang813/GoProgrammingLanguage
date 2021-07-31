// Exercise 1.10
// Exercise 1.11
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	if _, err := os.Stat("dump/"); os.IsNotExist(err) {
		err := os.Mkdir("dump/", 0777)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error creating dump directory.\n")
			os.Exit(1)
		}
	}
	file, err := os.OpenFile("dump/fetchAll.dump", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error opening dump file.\n")
		os.Exit(1)
	}

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		file.Write([]byte(<-ch))
	}
	file.Write([]byte("----------------------------------------------------------------\n"))

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}
