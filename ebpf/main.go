package main

import (
	"flag"
	"fmt"
)

func fib(n uint) uint {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func Fib(n uint) uint {
	return fib(n)
}

func main() {
	n := flag.Uint("n", 10, "The Fibonacci number to calculate")
	flag.Parse()

	fmt.Println(*n, "=>", Fib(*n))
}
