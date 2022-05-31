package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var A int
var B int
var N int

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

	aa, err := strconv.ParseInt(line1[0], 10, 64)
	if err != nil {
		fmt.Println("error")
		return
	}
	A = int(aa)

	bb, err := strconv.ParseInt(line1[1], 10, 64)
	if err != nil {
		fmt.Println("error")
		return
	}
	B = int(bb)

	nn, err := strconv.ParseInt(line1[2], 10, 64)
	if err != nil {
		fmt.Println("error")
		return
	}
	N = int(nn)

	var day int
	for N > 0 {
		strDay := strconv.Itoa(day)
		dayChars := strings.Split(strDay, "")
		point := 0
		addDay := 0
		for i, char := range dayChars {
			if char == "7" {
				if i == 1 && N > 1000 {
					point = 10 * (A + B)
					addDay = 10
					break
				}
				addDay = 1
				point = A + B
			} else {
				addDay = 1
				point = A
			}
		}

		N -= point
		if N <= 0 {
			fmt.Println(day)
		}
		day += addDay
	}
}
