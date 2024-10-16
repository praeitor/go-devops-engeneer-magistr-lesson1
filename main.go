package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	url := "http://srv.msk01.gigacorp.local/_stats"
	client := &http.Client{}

	for {
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("Error fetching data:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		fmt.Println("Data received:", string(body))

		time.Sleep(10 * time.Second)
	}
}
