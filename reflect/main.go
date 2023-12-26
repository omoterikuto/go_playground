package main

import (
	"fmt"
	"reflect"
)

type hoge struct {
	a  int
	ap *int
	b  string
	bp *string
}

func main() {
	bp := "bp"
	v := &hoge{
		a:  1,
		ap: nil,
		bp: &bp,
	}
	//reflect.Type()
	t := reflect.TypeOf(v).Elem()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type, f.Type.Kind())
	}
}
