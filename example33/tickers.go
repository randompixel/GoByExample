package main

import (
	"fmt"
	"time"
)

func main() {

	// If timers are like setTimeout, tickers are like setInterval

	ticker := time.NewTicker(time.Millisecond * 500)

	// Keep ticking every 1/2 second
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
		}
	}()

	// Sleep for 1600ms - enough for 3 ticks ^^
	time.Sleep(time.Millisecond * 1600)
	// Stop the original ticker
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
