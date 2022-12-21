package main

import (
	"log"
	"math/rand"
	"sync"
)

func main() {
	log.SetFlags(log.Lmicroseconds)

	a1 := makeMatrix()
	b1 := makeMatrix()
	a2 := makeMatrix()
	b2 := makeMatrix()
	a3 := makeMatrix()
	b3 := makeMatrix()
	a4 := makeMatrix()
	b4 := makeMatrix()

	var wg sync.WaitGroup
	wg.Add(3)

	log.Println("start")
	go calc(a1, b1, &wg)
	go calc(a2, b2, &wg)
	go calc(a3, b3, &wg)
	calc(a4, b4, &wg)
	wg.Wait()
	log.Println("done")
}

const n = 2048

func calc(a1, b1 [n][n]int, wg *sync.WaitGroup) {
	var p1 [n][n]int
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			for j := 0; j < n; j++ {
				p1[i][j] += a1[i][k] * b1[k][j]
			}
		}
	}
	wg.Done()
	return
}

func makeMatrix() [n][n]int {
	var r [n][n]int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			r[i][j] = rand.Intn(100)
		}
	}
	return r
}
