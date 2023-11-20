package hoge

import "fmt"

type hoge struct {
	h string
}

func NewHoge() hoge {
	return hoge{
		h: "aa",
	}
}

func (h hoge) Name() {
	fmt.Println(h.h)
}
