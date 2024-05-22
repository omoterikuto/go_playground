package main

import "fmt"

func main() {
	userId := []string{"1", "2", "3", "4"}
	i := 0
	for _, id := range userId {
		fmt.Println(id)
		i += 1
		if i == 2 {
			fmt.Println("hoge")
			userId = userId[i:]
		}
	}
}
