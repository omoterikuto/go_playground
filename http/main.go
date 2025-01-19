package main

import (
	"io"
	"net/http"
)

func main() {
	req := http.NewRequest("GET", "http://example.com", nil)
	resp, err := http.DefaultClient.Do(req)
	resp.Body.Close()
	io.ReadAll(resp.Body)
	&http.Transport{

}
