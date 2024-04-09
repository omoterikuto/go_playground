package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "%s", "test")
}
