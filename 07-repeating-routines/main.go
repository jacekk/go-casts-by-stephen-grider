package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("[x]", link, "seems down!")
	} else {
		fmt.Println("[ ]", link, "is online")
	}

	c <- link
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

	for _, link := range links {
		go checkLink(link, c)
	}

	/*
		NOTE:
		Main thread is blocked with the next Sleep,
		and cannot receive messages from channel.
	*/
	// for l := range c {
	// 	time.Sleep(5 * time.Second)
	// 	go checkLink(l, c)
	// }

	// Different/better approach:
	for link := range c {
		go func(linkCopy string) {
			time.Sleep(5 * time.Second)
			checkLink(linkCopy, c)
		}(link)
	}
}
