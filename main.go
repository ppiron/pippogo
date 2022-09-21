package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// fmt.Println("Hello, World!")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(io.ReadAll(resp.Body))
}
