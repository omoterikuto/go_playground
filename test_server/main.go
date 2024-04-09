package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type WebhookReq struct {
	ID        uint64    `json:"id"`
	Address   string    `json:"address"`
	Fields    []string  `json:"fields"`
	Topic     string    `json:"topic"`
	CreatedAt time.Time `json:"created_at"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	fmt.Printf("%s\n", body)
	var req WebhookReq

	if err := json.Unmarshal(body, &req); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("body: %+v\n", req)

	w.Write([]byte("success!!"))
}

func main() {
	http.HandleFunc("/test", helloHandler)
	http.ListenAndServe(":8081", nil)
}
