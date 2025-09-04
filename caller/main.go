package main

import "caller/hoge"

func A() {
	B()
}

func B() {
	hoge.Hoge()
}

func main() {
	A()
}
