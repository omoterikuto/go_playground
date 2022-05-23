package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

// この `pong` 関数は、1 つ目のチャネルを受信専用で (`pings`)、
// 2 つ目のチャネルを送信専用で (`pongs`) 受け取ります。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Hello")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
