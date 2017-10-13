package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {

	// os signals get sent down a channel
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// register channel to receive specific signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// blocks for signals, prints then notifys we can finish
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	// wait for finish signal from go routine
	<-done
	fmt.Println("exiting")
}
