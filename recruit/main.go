package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var seed string
var n string

type serverResponse struct {
	n        int64
	response int64
}

var responses []serverResponse

func Exit() {
	fmt.Println("error!")
	os.Exit(1)
}

func main() {
	args := os.Args
	if len(args) != 3 {
		Exit()
	}
	seed = args[1]
	n = args[2]

	intN, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		Exit()
	}
	ans := f(intN)
	fmt.Println(ans)
}

func f(nn int64) int {
	switch nn {
	case 0:
		return 1
	case 2:
		return 2
	}

	if nn%2 == 0 {
		return f(nn-1) + f(nn-2) + f(nn-3) + f(nn-4)
	} else {
		for _, response := range responses {
			if response.n == nn {
				return int(response.response)
			}
		}
		res := askServer(nn)
		r := serverResponse{
			n:        nn,
			response: int64(res),
		}
		responses = append(responses, r)
		return res
	}
}

type Data struct {
	Seed   string
	N      int
	Result int
}

func askServer(nn int64) int {
	url := "http://challenge-server.code-check.io/api/recursive/ask"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Exit()
	}

	params := request.URL.Query()
	params.Add("n", strconv.Itoa(int(nn)))
	params.Add("seed", seed)
	request.URL.RawQuery = params.Encode()

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		Exit()
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Exit()
	}

	var data Data

	if err := json.Unmarshal(body, &data); err != nil {
		Exit()
	}

	return data.Result
}
