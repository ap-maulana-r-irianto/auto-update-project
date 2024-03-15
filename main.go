package main

import "time"

func main() {
	go updateStatus()
	time.Sleep(1 * time.Second) // Delay to ensure status.json is created before reading
	displayStatus()
}
