package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct{}

func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://ua.kapias.net")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	/*
		// way #1
		bs := make([]byte, 999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/

	/*
		// way #2
		io.Copy(os.Stdout, resp.Body)
	*/

	// way #3 - implemented Writer interface
	logger := logWritter{}
	io.Copy(logger, resp.Body)
}
