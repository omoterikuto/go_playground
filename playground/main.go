package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var N int
var A []int

func main() {
	file := "data.txt"
	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	scanner.Scan()

	nn, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	N = int(nn)

	scanner.Scan()
	line2 := strings.Split(scanner.Text(), " ")
	for _, v := range line2 {
		a, _ := strconv.ParseInt(v, 10, 64)
		A = append(A, int(a))
	}

	var count int
	for i := 0; i < N; i++ {
		v := make([]int, 3)
		v[0] = A[i]
		for j := i + 1; j < N; j++ {
			if v[0] == A[j] {
				continue
			}
			v[1] = A[j]
			for k := j + 1; k < N; k++ {
				if v[0] == A[k] || v[1] == A[k] {
					continue
				}
				count++
			}
		}
	}
	fmt.Println(count)
}
