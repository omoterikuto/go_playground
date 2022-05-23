package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

var result []string

func main() {
	file := "data.txt"
	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	var rows []string
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	fmt.Println(rows)

	resultCh := make(chan []string, len(rows))

	for i, v := range rows {
		go getSHA256Binary(i, v, resultCh)
	}

	for len(result) != len(rows) {
		fmt.Println("wait...")
	}
	for _, v := range result {
		fmt.Println(v)
	}
}

func getSHA256Binary(i int, v string, resultCh chan []string) {
	r := sha256.Sum256([]byte(v))
	h := hex.EncodeToString(r[:])
	resultCh[i] = h
	return
}
