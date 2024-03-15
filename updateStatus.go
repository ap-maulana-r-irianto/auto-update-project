package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func updateStatus() {
	for {
		// Generate random values for water and wind
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		// Update status
		data := Data{
			Status: Status{
				Water: water,
				Wind:  wind,
			},
		}

		// Marshal data to JSON
		fileData, err := json.MarshalIndent(data, "", " ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		// Write to file
		err = ioutil.WriteFile("status.json", fileData, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Println("Status updated")

		// Wait for 15 seconds
		time.Sleep(15 * time.Second)
	}
}
