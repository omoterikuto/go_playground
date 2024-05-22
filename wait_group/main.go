package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	hoge()
	//time.Sleep(5 * time.Second)
}

func hoge() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("goroutine 1")
	}()
	fmt.Println("main")
}
