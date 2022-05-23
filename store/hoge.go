package store

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

package main

import (
"fmt"
"io/ioutil"
"os"
"strings"
"strconv"
)

func main() {
	// このコードは標準入力と標準出力を用いたサンプルコードです。
	// このコードは好きなように編集・削除してもらって構いません。
	// ---
	// This is a sample code to use stdin and stdout.
	// Edit and remove this code as you like.

	lines := getStdin()

	switch lines[0] {
	case "1":
		case1(lines)
	}
}

func getStdin() []string {
	stdin, _ := ioutil.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}

type menuInfo struct {
	id    int
	stock int
	value int
}

type orderInfo struct {
	seat       int
	menuID     int
	orderCount int
}

func case1(lines []string) {
	mm := lines[1]
	m, err := strconv.ParseInt(mm, 10, 64)
	if err != nil {
		panic(err)
	}
	var menuInfos []*menuInfo
	for i := 0; i < int(m); i++ {
		args := strings.Split(lines[i+2], " ")
		id, _ := strconv.ParseInt(args[0], 10, 64)
		stock, _ := strconv.ParseInt(args[1], 10, 64)
		value, _ := strconv.ParseInt(args[2], 10, 64)

		info := &menuInfo{
			id:    int(id),
			stock: int(stock),
			value: int(value),
		}
		menuInfos = append(menuInfos, info)
	}

	var orderInfos []orderInfo
	for  i := 0; i < len(lines) - int(m) - 2; i++ {
		args := strings.Split(lines[i+int(m)+2], " ")
		seat, _ := strconv.ParseInt(args[1], 10, 64)
		menuID, _ := strconv.ParseInt(args[2], 10, 64)
		orderCount, _ := strconv.ParseInt(args[3], 10, 64)
		order := orderInfo{
			seat:       int(seat),
			menuID:     int(menuID),
			orderCount: int(orderCount),
		}
		orderInfos = append(orderInfos, order)
	}

	for _, order := range orderInfos {
		for _, menu := range menuInfos {
			if order.menuID == menu.id {

				if menu.stock < order.orderCount {
					fmt.Printf("sold out %d\n", order.seat)
				} else {
					menu.stock -= order.orderCount
					for i := 0; i < order.orderCount; i++ {
						fmt.Printf("received order %d %d\n", order.seat, order.menuID)
					}
				}
			}
		}
	}
}

