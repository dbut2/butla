package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	buf := bytes.NewBufferString("{\"url\":\"https://example.com\"}")

	resp, err := http.Post("https://but.la/shorten", "application/json", buf)
	if err != nil {
		log.Fatalln(err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(resp.StatusCode, resp.Status, string(body))

}
