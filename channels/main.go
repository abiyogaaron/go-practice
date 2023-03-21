package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://amazon.com",
		"https://stackoverflow.com",
		"https://facebook.com",
		"https://golang.org",
	}
	// Create a channel to communicate between go routine
	c := make(chan string)

	for _, url := range urls {
		go checkLink(url, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down!")
		c <- url
		return
	}
	c <- url
	fmt.Println(url, "is up!")
}
