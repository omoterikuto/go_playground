package main

import "fmt"

func main() {
	hoge()
}

func hoge() string {
	defer a()
	return b()
}

func a() {
	fmt.Println("a")
}

func b() string {
	fmt.Println("b")
	return "b"
}
