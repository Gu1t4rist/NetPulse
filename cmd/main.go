package main

import (
	"fmt"
	"netpulse/internal/speedtest"
	"netpulse/internal/database"
)

func main() {
	result, err := speedtest.RunSpeedtest()
	if err != nil {
		fmt.Printf("Speed test failed:", err)
		return
	}

	go api.StartServer()

	err = database.LogResult(result)
	if err != nil {
		fmt.Println("Failed to log result:", err)
		return
	}

	fmt.Println("Speed test result logged successfully!")
}