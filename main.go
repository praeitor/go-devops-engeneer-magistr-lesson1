package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {
	url := "http://srv.msk01.gigacorp.local/_stats"
	client := &http.Client{}
	errorCount := 0

	for {
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("Error fetching data:", err)
			errorCount++
			if errorCount >= 3 {
				fmt.Println(("Unaable to fetch server stat"))
			}
			time.Sleep(10 * time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Non-200 response received:", resp.Status)
			errorCount++
			if errorCount >= 3 {
				fmt.Println("Unable to fetch server statistic")
			}
			resp.Body.Close()
			time.Sleep(10 * time.Second)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			resp.Body.Close()
			time.Sleep(10 * time.Second)
			continue
		}

		data := strings.Split(string(body), ",")
		if len(data) < 6 {
			fmt.Println("Invalid format data")
			errorCount++
			if errorCount >= 3 {
				fmt.Println("Unable to fetch server stat")
			}
			resp.Body.Close()
			time.Sleep(10 * time.Second)
			continue
		}

		resp.Body.Close()

		time.Sleep(10 * time.Second)
	}
}
