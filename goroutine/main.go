package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	links := []string{
		"https://google.com",
		"https://amazon.com",
		"https://stackoverflow.com",
		"https://facebook.com",
		"https://golang.org",
	}

	fmt.Println("Running process...")
	var wg sync.WaitGroup
	wg.Add(len(links))

	for _, link := range links {
		go func(link string) {
			_, err := http.Get(link)
			defer wg.Done()

			if err != nil {
				fmt.Println(link, "might be down!")
				return
			}
			fmt.Println(link, "is up!")
		}(link)
	}
	fmt.Println("Doing other stuff")

	wg.Wait()
	fmt.Println("Process finished...")
}
