package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go hoge(ctx, wg)
	wg.Wait()

}

func hoge(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
			wg.Done()
			return
		default:
		}
		fmt.Println("a")
		time.Sleep(time.Second)
	}
}
