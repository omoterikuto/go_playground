package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var N int
var P []int

func main() {
	file := "data.txt"
	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	scanner.Scan()

	nn, err := strconv.ParseInt(scanner.Text(), 10, 64)
	N = int(nn)

	scanner.Scan()
	line2 := strings.Split(scanner.Text(), " ")
	for _, v := range line2 {
		p, _ := strconv.ParseInt(v, 10, 64)
		P = append(P, int(p))
	}

	connection := make([][]int, N)
	for i := 0; i < N; i++ {
		connection[i] = make([]int, N)
	}
	for i, v := range P {
		connection[i+1][v-1] = 1
		connection[v-1][i+1] = 1
	}
	for _, bools := range connection {
		for _, b := range bools {
			fmt.Printf("%v ", b)
		}
		fmt.Println("")
	}

	calculateNumber(connection)
}

func calculateNumber(connection [][]int) {
	for i := 0; i < N; i++ {
		visitable := []int{i}
		var visited []int
		count := 1

		for j := 0; j < 3; j++ {
			//fmt.Println(j, "______time")
			for _, roomToCheck := range visitable {
				visitable = visitable[1:]
				//fmt.Println(visitable)
				//fmt.Println("connection", connection[roomToCheck])
				for room, connectRoom := range connection[roomToCheck] {
					if connectRoom == 1 {
						noExist := true
						for _, v := range visited {
							if v == room {
								noExist = false
								break
							}
						}
						if noExist {
							visitable = append(visitable, room)
							count++
						}
					}
				}
				//fmt.Println("add to visit", roomToCheck)
				visited = append(visited, roomToCheck)
			}
		}
		fmt.Println(count)
	}
}
