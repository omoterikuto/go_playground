package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello https"))
	})
	err := http.ListenAndServeTLS(":443", "test_server.crt", "test_server.key", nil)
	if err != nil {
		panic(err)
	}
}
