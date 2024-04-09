package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	hogeMap := map[string]map[int]float64{
		"key1": {1: 1.1, 2: 2.2},
	}
	by, _ := json.Marshal(hogeMap)
	fmt.Println(string(by))
}
