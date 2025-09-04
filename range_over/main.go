package main

import (
	"fmt"
	"time"
)

func main() {
	for x := range f {
		fmt.Printf("loop=%d, begin\n", x)
		time.Sleep(3 * time.Second)
		fmt.Printf("loop=%d, end\n", x)
	}
}

func f(yield func(int) bool) {
	fmt.Println("start...")
	yield(1)
	fmt.Println("midway...")
	yield(2)
	fmt.Println("finish!")
}
