package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	a := strings.Split(sc.Text(), " ")
	//n, _ := strconv.Atoi(a[0])
	leftNum, _ := strconv.Atoi(a[1])

	sc.Scan()
	strSubTotal := strings.Split(sc.Text(), " ")

	sub := make([]int, len(strSubTotal))
	for i, s := range strSubTotal {
		b, _ := strconv.Atoi(s)
		sub[i] = b
	}

	ans := make([]int, len(sub))
	ans[0] = leftNum

	for i := 0; i < len(sub)-1; i++ {
		ans[i+1] = sub[i] - ans[i]
		if i != 0 {
			ans[i+1] -= ans[i-1]
		}
	}
	for i, a := range ans {
		fmt.Printf("%d", a)
		if i != len(ans) {
			fmt.Printf(" ")
		}
	}
}
