package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.ParseInLocation("20060102", "", time.UTC))
}
