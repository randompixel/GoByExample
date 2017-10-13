package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg // place msg into channel pings
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings // read message from pings into msg
	pongs <- msg   // send message to pongs
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
