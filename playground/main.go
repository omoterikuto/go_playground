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
	wg.Add(4)
	b1t := transpose(b1)
	b2t := transpose(b2)
	b3t := transpose(b3)
	b4t := transpose(b4)
	log.Println("start")
	go calc(a1, b1t, &wg)
	go calc(a2, b2t, &wg)
	go calc(a3, b3t, &wg)
	go calc(a4, b4t, &wg)
	wg.Wait()
	log.Println("done")
}

const n = 2048

func calc(a1, b1 [n][n]int, wg *sync.WaitGroup) {
	var p1 [n][n]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
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

func transpose(p [n][n]int) [n][n]int {
	var r [n][n]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			r[i][j] = p[i][j]
		}
	}
	return r
}
