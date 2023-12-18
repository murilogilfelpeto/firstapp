package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// defer will be executed at the end of the function
	// if we add multiple defer statements, they will be executed in LIFO order

	res, err := http.Get("http://www.google.com/robots.txt")
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	robots, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
