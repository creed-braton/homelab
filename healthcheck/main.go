package main

import (
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	url := os.Args[1]
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		os.Exit(1)
	}

	os.Exit(0)
}
