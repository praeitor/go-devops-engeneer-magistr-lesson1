package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
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

		loadAvg, err := strconv.ParseFloat(data[0], 64)
		if err == nil && loadAvg > 30 {
			fmt.Printf("Load Average is too high: %d\n", int(loadAvg))
		}

		totalMemory, err := strconv.ParseFloat(data[1], 64)
		usedMemory, err2 := strconv.ParseFloat(data[2], 64)
		if err == nil && err2 == nil {
			memoryUsage := (usedMemory / totalMemory) * 100
			if memoryUsage > 80 {
				fmt.Printf("Memory usage too high: %d%%\n", int(memoryUsage))
			}
		}

		totalDisk, err := strconv.ParseFloat(data[3], 64)
		usedDisk, err2 := strconv.ParseFloat(data[4], 64)
		if err == nil && err2 == nil {
			freeDiskSpace := (usedDisk - usedDisk) / (1024 * 1024)
			if usedDisk/totalDisk > 0.9 {
				fmt.Printf("Free disk space is to low: %d Mb left\n", int(freeDiskSpace))
			}
		}

		totalBandwidth, err := strconv.ParseFloat(data[5], 64)
		usedBandwidth, err2 := strconv.ParseFloat(data[6], 64)
		if err == nil && err2 == nil {
			freeBandwidth := (totalBandwidth - usedBandwidth) / (1024 * 1024 / 8)
			if usedBandwidth/totalBandwidth > 0.9 {
				fmt.Printf("Network bandwidth usage high: %d Mbit/s available\n", int(freeBandwidth))
			}
		}

		resp.Body.Close()
		time.Sleep(10 * time.Second)
	}
}
