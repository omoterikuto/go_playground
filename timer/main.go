package main

import "time"

func main() {
	watchDogTimer := time.NewTicker(time.Second).C
	go func() {
		for {
			select {
			case <-watchDogTimer:
				println("watch dog timer fired")
			}
		}
	}()
	time.Sleep(time.Second * 100)
}
