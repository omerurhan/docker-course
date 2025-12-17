package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://raw.githubusercontent.com/" +
		"omerurhan/docker-course/refs/heads/main/README.MD"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println("response:\n", string(body))
}
