package main

import (
	"fmt"
	"time"
)

func main() {
	for _, v := range []int{1, 2, 3} {
		go func() {
			//fmt.Println(v)
			fmt.Println(&v)
		}()
	}
	// ちょっと待つ
	time.Sleep(1 * time.Second)
}
