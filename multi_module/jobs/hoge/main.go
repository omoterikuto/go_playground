package main

import (
	"fmt"
	"github.com/oklog/ulid/v2"
	"multi_module/shared"
)

func main() {
	fmt.Println("hoge")
	fmt.Println(ulid.Make())
	shared.SHoge()
}
