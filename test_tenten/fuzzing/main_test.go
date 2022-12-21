package main

import (
	"fmt"
	"testing"
)

func hoge(s string) string {
	return s
}

func FuzzCalc(f *testing.F) {
	fmt.Println(Calc("1+1"))
	f.Add("1+1")
	f.Fuzz(func(t *testing.T, queryStr string) {
		hoge(queryStr)
	})
}
