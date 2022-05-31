package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var N int
var M int
var K int

func main() {
	file := "data.txt"
	fp, err := os.Open(file)
	if err != nil {
		fmt.Println("error")
		return
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	//scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line1 := strings.Split(scanner.Text(), " ")

	nn, err := strconv.ParseInt(line1[0], 10, 64)
	if err != nil {
		fmt.Println("error")
		return
	}
	N = int(nn)

	mm, err := strconv.ParseInt(line1[1], 10, 64)
	if err != nil {
		fmt.Println("error")
		return
	}
	M = int(mm)

	kk, err := strconv.ParseInt(line1[2], 10, 64)
	if err != nil {
		fmt.Println("error")
		return
	}
	K = int(kk)

	var color []int64
	scanner.Scan()
	line2 := strings.Split(scanner.Text(), " ")
	for _, line := range line2 {
		l, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Println("error")
			return
		}
		color = append(color, l)
	}

	pass := make([][]int64, N)
	for i := 0; i < N; i++ {
		pass[i] = make([]int64, N)
	}
	for i := 0; i < M; i++ {
		scanner.Scan()
		line3 := strings.Split(scanner.Text(), " ")
		a, err := strconv.ParseInt(line3[0], 10, 64)
		if err != nil {
			fmt.Println("error")
			return
		}
		b, err := strconv.ParseInt(line3[1], 10, 64)
		if err != nil {
			fmt.Println("error")
			return
		}
		pass[a-1][b-1] = 1
		pass[b-1][a-1] = 1
	}

	for t := 0; t < K; t++ {
		afterColor := make([]int64, N)
		for i := 0; i < N; i++ {
			colorCount := make([]int64, N)
			for i2, p := range pass[i] {
				if p == 1 {
					colorCount[color[i2]]++
				}
			}
			fmt.Println(colorCount)
			var maxColorCount int64
			var maxColor int64
			for i2, i3 := range colorCount {
				if maxColorCount < i3 {
					maxColorCount = i3
					maxColor = int64(i2)
				}
			}
			afterColor[i] = maxColor
		}
		//fmt.Println(afterColor)
		color = afterColor
	}
	for _, i2 := range color {
		fmt.Println(i2)
	}
}
