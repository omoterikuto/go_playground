package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx.Done()
		w.WriteHeader(http.StatusInternalServerError)
		//time.Sleep(time.Second * 2)
		//fmt.Fprint(w, "Hello, playground")
	})

	log.Println("Starting server...")
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Fatal(http.Serve(l, nil))
	}()

	log.Println("Sending request...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/hello", nil)
	//res, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		log.Fatal(err)
	}
	cli := http.Client{}
	resp, err := cli.Do(req)
	log.Println(err)
	log.Println(resp)
}
