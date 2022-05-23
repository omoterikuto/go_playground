package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var N int
var W int64
var A []int64

func main() {
	file := "data.txt"
	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	scanner.Scan()
	line1 := strings.Split(scanner.Text(), " ")

	nn, err := strconv.ParseInt(line1[0], 10, 64)
	N = int(nn)

	ll, err := strconv.ParseInt(line1[1], 10, 64)
	W = ll

	scanner.Scan()
	line2 := strings.Split(scanner.Text(), " ")
	for _, v := range line2 {
		p, _ := strconv.ParseInt(v, 10, 64)
		A = append(A, p)
	}

	var result []int
	for len(result) != N {
		for i, _ := range A {
			if A[i] <= 0 {
				continue
			}
			A[i] -= W
			if A[i] <= 0 {
				result = append(result, i)
			}
		}
	}
	for _, v := range result {
		fmt.Println(v + 1)
	}
}
