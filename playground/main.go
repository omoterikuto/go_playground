package main

import (
	"fmt"
	"reflect"
)

type ComplexityRoot struct {
	Account struct {
		Belong   func(childComplexity int) int
		Birthday func(childComplexity int) int
	}
}

func main() {
	cr := ComplexityRoot{}

	cr.Account.Belong = func(childComplexity int) int {
		return childComplexity
	}
	//cr.Account.Birthday = func(childComplexity int) int {
	//	return childComplexity
	//}

	crValue := reflect.ValueOf(cr) // Account{}

	for i := 0; i < crValue.NumField(); i++ {
		resolver := crValue.Field(i) // [Belong, Birthday]
		for j := 0; j < resolver.NumField(); j++ {
			if resolver.Field(j).IsNil() {
				fmt.Println("not implement")
			}
		}
	}
}
