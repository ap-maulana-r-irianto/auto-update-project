package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func getStatus() Data {
	// Read the JSON file
	fileData, err := ioutil.ReadFile("status.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Unmarshal the JSON data
	var data Data
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	return data
}

func determineStatus(value int, thresholdLow int, thresholdHigh int) string {
	switch {
	case value < thresholdLow:
		return "aman"
	case value >= thresholdLow && value <= thresholdHigh:
		return "siaga"
	default:
		return "bahaya"
	}
}

func displayStatus() {
	for {
		data := getStatus()

		// Determine water and wind status
		waterStatus := determineStatus(data.Status.Water, 5, 8)
		windStatus := determineStatus(data.Status.Wind, 6, 15)

		fmt.Printf("Water Status: %d meter (%s), Wind Status: %d meter per second (%s)\n", data.Status.Water, waterStatus, data.Status.Wind, windStatus)

		// Refresh every 15 seconds
		time.Sleep(15 * time.Second)
	}
}
