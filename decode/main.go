package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	decoded, err := base64.RawURLEncoding.DecodeString("")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decoded))
}
