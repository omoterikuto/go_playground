package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var N int
var W int64
var A []int64

func main() {
	file := "data.txt"
	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	scanner.Scan()
	switch scanner.Text() {
	case "1":
		case1(scanner)
	case "2":
		case2(scanner)
	}
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

func case1(scanner *bufio.Scanner) {
	scanner.Scan()
	mm := scanner.Text()
	m, err := strconv.ParseInt(mm, 10, 64)
	if err != nil {
		panic(err)
	}
	var menuInfos []*menuInfo
	for i := 0; i < int(m); i++ {
		scanner.Scan()
		args := strings.Split(scanner.Text(), " ")
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
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
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

	for _, v := range menuInfos {
		fmt.Println(v.id, v.stock, v.value)
	}
	for _, v := range orderInfos {
		fmt.Println(v.seat, v.menuID, v.orderCount)
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

type receiveInfo struct {
	state  string
	menuID int
	seat   int
}

func case2(scanner *bufio.Scanner) {
	scanner.Scan()
	lines2 := strings.Split(scanner.Text(), " ")
	m, err := strconv.ParseInt(lines2[0], 10, 64)
	if err != nil {
		panic(err)
	}
	k, err := strconv.ParseInt(lines2[1], 10, 64)
	if err != nil {
		panic(err)
	}

	var menuInfos []*menuInfo
	for i := 0; i < int(m); i++ {
		scanner.Scan()
		args := strings.Split(scanner.Text(), " ")
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

	var receiveInfos []receiveInfo
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		var info receiveInfo
		if args[0] == "received" {
			seat, _ := strconv.ParseInt(args[2], 10, 64)
			menuID, _ := strconv.ParseInt(args[3], 10, 64)
			info = receiveInfo{
				state:  args[0],
				menuID: int(menuID),
				seat:   int(seat),
			}
		} else {
			menuID, _ := strconv.ParseInt(args[1], 10, 64)
			info = receiveInfo{
				state:  args[0],
				menuID: int(menuID),
			}
		}
		receiveInfos = append(receiveInfos, info)
	}

	doing := make([]*receiveInfo, k)

	for i, info := range receiveInfos {
		if info.state == "received" {
			isAddInfo := false
			for j := 0; j < int(k); j++ {
				if doing[j] == nil {
					doing[j] = &receiveInfos[i]
					isAddInfo = true
					fmt.Println(info.menuID)
					break
				}
			}
			if !isAddInfo {
				fmt.Println("wait")
			}
		} else {
			for _, v := range doing {
				if v == nil {
					continue
				}
				if v.menuID == info.menuID {
					v = nil
					fmt.Printf("ok %d\n", info.menuID)
					break
				}
			}
			fmt.Println("unexpected input")
		}
	}
}
