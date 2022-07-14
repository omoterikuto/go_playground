package main

import (
	"fmt"
)

func main() {
	a := []int{2, 1, 3, 5, 4}
	fmt.Println(Solution(a))
}

func Solution(A []int) int {
	states := make([]int, len(A))
	count := 0
	for _, a := range A {
		aa := a - 1
		states[aa] = 1
		on := true
		for i := 0; i < aa; i++ {
			if states[i] != 1 {
				on = false
				break
			}
		}
		if on {
			count++
		}
	}
	return count
}
