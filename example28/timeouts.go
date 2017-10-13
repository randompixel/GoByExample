package main

import (
	"fmt"
	"time"
)

func main() {

	// Make a channel and set to sleep for 2 seconds
	// before returning a result
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// If result doesn't return in a second then
	// jump to timeout clause
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	// Make a new channel and set to sleep for 2 seconds
	// before returning a result
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	// As above but timeout doesn't occur until 3 seconds
	// so should return the normal result
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
