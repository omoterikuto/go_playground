package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var A []string

func main() {
	file := "data.txt"
	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	scanner.Scan()
	line := strings.Split(scanner.Text(), "")
	for _, v := range line {
		A = append(A, v)
	}

	bws := []string{"b", "w"}

	for i, v := range A {
		if i%2 == 0 {
			if v == "R" {
				allW := true
				bws = append(bws, "b")
				for i := 0; i < len(bws)-1; i++ {
					if bws[i] == "w" {
						continue
					} else {
						allW = false
						break
					}
				}
				if allW {
					continue
				}
				for i := len(bws) - 2; i >= 0; i-- {
					if bws[i] == "w" {
						bws[i] = "b"
					} else {
						break
					}
				}
			} else {
				allW := true
				bws = append([]string{"b"}, bws...)
				for i := 1; i < len(bws); i++ {
					if bws[i] == "w" {
						continue
					} else {
						allW = false
						break
					}
				}
				if allW {
					continue
				}
				for i := 1; i < len(bws); i++ {
					if bws[i] == "w" {
						bws[i] = "b"
					} else {
						break
					}
				}
			}
		} else {
			if v == "R" {
				bws = append(bws, "w")
				allB := true
				for i := 0; i < len(bws)-1; i++ {
					if bws[i] == "b" {
						continue
					} else {
						allB = false
						break
					}
				}
				if allB {
					continue
				}
				for i := len(bws) - 2; i >= 0; i-- {
					if bws[i] == "b" {
						bws[i] = "w"
					} else {
						break
					}
				}
			} else {
				bws = append([]string{"w"}, bws...)
				allB := true
				for i := 1; i < len(bws); i++ {
					if bws[i] == "b" {
						continue
					} else {
						allB = false
						break
					}
				}
				if allB {
					continue
				}
				for i := 1; i < len(bws); i++ {
					if bws[i] == "b" {
						bws[i] = "w"
					} else {
						break
					}
				}
			}
		}
	}
	blackCount := 0
	whiteCount := 0
	for _, bw := range bws {
		if bw == "b" {
			blackCount++
			continue
		}
		whiteCount++
	}
	fmt.Println(blackCount, whiteCount)
}
