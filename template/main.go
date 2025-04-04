package main

import (
	"html/template"
	"os"
)

type Fuga struct {
	Hoges []string
}

func First(slice []string) string {
	if len(slice) == 0 {
		var zero string
		return zero
	}
	return slice[0]
}

func main() {
	fuga := Fuga{[]string{"foo"}}

	var customFunctions = template.FuncMap{
		"first": First,
	}

	tmpl, err := template.New("test").Funcs(customFunctions).Parse("{{first Hoges}}")

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, fuga)
	if err != nil {
		panic(err)
	}
}

var customFunctions = template.FuncMap{
	"comma":     format.Comma,
	"date":      jst.DisplayDate,
	"dateTime":  jst.DisplayDateTime,
	"yearMonth": jst.DisplayYearMonth,
	"first":     First,
}

func First[T any](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	return slice[0]
}
