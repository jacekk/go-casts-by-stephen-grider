package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string, index int, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		c <- fmt.Sprintf("failed %d", index)
		fmt.Println(index, "[x]", link, "seems down!")
	} else {
		c <- fmt.Sprintf("done %d", index)
		fmt.Println(index, "[ ]", link, "is online")
	}
}

func main() {
	links :=
		[]string{
			"http://amazon.com",
			"http://google.com",
			"http://invalid-subdomain.kapias.net",
			"http://ua.kapias.net",
		}

	c := make(chan string)

	for index, link := range links {
		go checkLink(link, index+1, c)
	}

	for index := range links {
		fmt.Println("print", index, "-", <-c)
	}
}
