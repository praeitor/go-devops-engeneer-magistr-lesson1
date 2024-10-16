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

		loadAvg, err := strconv.Atoi(data[0])
		if err == nil && loadAvg > 30 {
			fmt.Printf("Load Average is too high: %d\n", int(loadAvg))
		}

		totalMemory, err := strconv.Atoi(data[1])
		usedMemory, err2 := strconv.Atoi(data[2])
		if err == nil && err2 == nil {
			memoryUsage := (usedMemory / totalMemory) * 100
			if memoryUsage > 80 {
				fmt.Printf("Memory usage too high: %d%%\n", int(memoryUsage))
			}
		}

		totalDisk, err := strconv.Atoi(data[3])
		usedDisk, err2 := strconv.Atoi(data[4])
		if err == nil && err2 == nil {
			freeDiskSpace := (totalDisk - usedDisk) / (1024 * 1024)
			if usedDisk*100 > totalDisk*90 {
				fmt.Printf("Free disk space is to low: %d Mb left\n", int(freeDiskSpace))
			}
		}

		totalBandwidth, err := strconv.Atoi(data[5])
		usedBandwidth, err2 := strconv.Atoi(data[6])
		if err == nil && err2 == nil {
			if usedBandwidth*100 > totalBandwidth*90 {
				freeBandwidth := (totalBandwidth - usedBandwidth) / (1024 * 1024 / 8)
				fmt.Printf("Network bandwidth usage high: %d Mbit/s available\n", int(freeBandwidth))
			}
		}

		resp.Body.Close()
		time.Sleep(10 * time.Second)
	}
}
