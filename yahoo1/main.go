package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var N int
var K int64

type Input struct {
	name  string
	value int64
}

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

	K, err = strconv.ParseInt(line1[1], 10, 64)

	var inputs []Input
	for i := 0; i < N; i++ {
		scanner.Scan()
		line2 := strings.Split(scanner.Text(), " ")
		val, err := strconv.ParseInt(line2[1], 10, 64)
		if err != nil {
			fmt.Println("error")
			return
		}
		inputs = append(inputs, Input{
			name:  line2[0],
			value: val,
		})
	}
	for _, input := range inputs {
		if input.value >= K {
			fmt.Println(input.name)
		}
	}
}
