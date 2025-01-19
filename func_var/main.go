package main

import "fmt"

func main() {
	hh := h()
	fmt.Println(hh())
	fmt.Println(hh())
	fmt.Println(hh())
	fmt.Println(hh())
}

func h() func() bool {
	var hoge bool
	return func() bool {
		hoge = !hoge
		return hoge
	}
}
