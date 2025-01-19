package main

import (
	"fmt"
	"time"
)

func main() {
	//target := version.Must(version.NewVersion("0.64.133"))
	//fmt.Println(target)
	now := time.Now()
	now = now.In(time.FixedZone("UTC", 0))
	fmt.Println(now.Unix(), now)
	now2 := time.Now()
	now2 = now2.In(time.FixedZone("JST", int(60*60*9)))
	fmt.Println(now2.Unix(), now2)
}
