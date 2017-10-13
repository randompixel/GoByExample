package main

import (
	"fmt"
	"time"
)

func main() {

	// These timers seem like JS's setTimeout, except with automagic channels

	timer1 := time.NewTimer(time.Second * 2)

	fmt.Println("Go")
	<-timer1.C // Block until timer1.C returns something
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

}
