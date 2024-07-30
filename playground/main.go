package main

import (
	"cloud.google.com/go/civil"
	"fmt"
	"time"
)

func main() {
	a := civil.Date{
		Year:  2024,
		Month: 1,
		Day:   1,
	}
	fmt.Println(a.In(time.FixedZone("JST", 9*60*60)).Format("20060102"))
	fmt.Println(a.In(time.UTC).Format("20060102"))
}
