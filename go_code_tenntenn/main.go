package go_code_tenntenn

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("aaa")
	err1 := errors.New("err1")
	err2 := errors.New("err2")
	err := errors.New("err2")
	fmt.Println(err)
	if errors.Is(err, err1) {
		fmt.Println("err is err1")
	}
	if errors.Is(err, err2) {
		fmt.Println("err is err2")
	}

}
func hoge(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
