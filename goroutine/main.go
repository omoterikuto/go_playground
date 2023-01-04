package main

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
)

func main() {
	log.SetFlags(log.Lmicroseconds)

	// 使用するCPUの最大数を4に
	runtime.GOMAXPROCS(4)

	// 二次元配列の組を4つ作成
	a1, b1, a2, b2, a3, b3, a4, b4 := makeMatrix(), makeMatrix(), makeMatrix(),
		makeMatrix(), makeMatrix(), makeMatrix(), makeMatrix(), makeMatrix()

	var wg sync.WaitGroup

	// WaitGroup counterを4加算
	wg.Add(4)

	log.Println("start")

	//　スレッドを新規に4つ作成してそれぞれのスレッドで行列計算
	go calc(a1, b1, &wg)
	go calc(a2, b2, &wg)
	go calc(a3, b3, &wg)
	go calc(a4, b4, &wg)

	// WaitGroup counterが0になるまで待機
	wg.Wait()

	log.Println("done")
}

const n = 2048

// 行列を計算する
func calc(a, b [n][n]int, wg *sync.WaitGroup) {
	var p [n][n]int
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ { // kループとjループを入れ替える
			for j := 0; j < n; j++ {
				p[i][j] += a[i][k] * b[k][j]

			}
		}
	}
	wg.Done() // WaitGroup counterを1減算
	return
}

// ランダムな二次元配列を作成
func makeMatrix() [n][n]int {
	var r [n][n]int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			r[i][j] = rand.Intn(100)
		}
	}
	return r
}
