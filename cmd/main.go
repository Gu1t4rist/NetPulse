package main

import (
	"fmt"
	"github.com/Gu1t4rist/NetPulse/internal/speedtest"
	"github.com/Gu1t4rist/NetPulse/internal/database"
	"github.com/Gu1t4rist/NetPulse/internal/api"
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