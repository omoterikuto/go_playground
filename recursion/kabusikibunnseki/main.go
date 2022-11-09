package main

import "fmt"

func main() {
	fmt.Println(stockSpan([]int32{34, 640, 100, 234, 56, 34, 25, 200, 1020, 160}))
}

func stockSpan(stocks []int32) []int32 {
	ret := make([]int32, len(stocks))

	for i, num := range stocks {
		count := int32(1)
		for j := i - 1; j >= 0; j-- {
			if num <= stocks[j] {
				break
			}
			count++
		}
		ret[i] = count
	}

	return ret
}
