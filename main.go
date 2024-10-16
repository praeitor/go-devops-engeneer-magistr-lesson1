package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "http://srv.msk01.gigacorp.local/_stats"
	client := &http.Client{}

	for {
		resp, err := client.Get(ulr)
		if err != nil {
			fmt.Println("Error fetching data:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nill {
			fmt.Println("Error reading response body:", err)
			resp.Body.Close()
			time.Sleep(10 * time.Second)
			continue
		}

		fmt.Println("Data received:", string(body))
		resp.Body.Close()

		time.Sleep(10 * time.Second)
	}
}
