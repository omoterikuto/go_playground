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
	case "3":
		case3(scanner)
	}
}

type menuInfo struct {
	id    int64
	stock int64
	value int64
}

type orderInfo struct {
	seat       int64
	menuID     int64
	orderCount int64
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
			id:    id,
			stock: stock,
			value: value,
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
			seat:       seat,
			menuID:     menuID,
			orderCount: orderCount,
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
					for i := 0; i < int(order.orderCount); i++ {
						fmt.Printf("received order %d %d\n", order.seat, order.menuID)
					}
				}
			}
		}
	}
}

type receiveInfo struct {
	state  string
	menuID int64
	seat   int64
}

func case2(scanner *bufio.Scanner) {
	scanner.Scan()
	lines2 := strings.Split(scanner.Text(), " ")
	m, err := strconv.ParseInt(lines2[0], 10, 64)
	if err != nil {
		panic(err)
	}
	kk, err := strconv.ParseInt(lines2[1], 10, 64)
	if err != nil {
		panic(err)
	}
	k := int(kk)

	var menuInfos []*menuInfo
	for i := 0; i < int(m); i++ {
		scanner.Scan()
		args := strings.Split(scanner.Text(), " ")
		id, _ := strconv.ParseInt(args[0], 10, 64)
		stock, _ := strconv.ParseInt(args[1], 10, 64)
		value, _ := strconv.ParseInt(args[2], 10, 64)

		info := &menuInfo{
			id:    id,
			stock: stock,
			value: value,
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
				menuID: menuID,
				seat:   seat,
			}
		} else {
			menuID, _ := strconv.ParseInt(args[1], 10, 64)
			info = receiveInfo{
				state:  args[0],
				menuID: menuID,
			}
		}
		receiveInfos = append(receiveInfos, info)
	}

	type orderStock struct {
		state  string
		menuID int64
		seat   int64
		done   bool
	}
	var orderStocks []*orderStock

	for _, info := range receiveInfos {
		if info.state == "received" {
			o := orderStock{
				state:  info.state,
				menuID: info.menuID,
				seat:   info.seat,
				done:   false,
			}
			orderStocks = append(orderStocks, &o)
			cockingCount := 0
			for _, order := range orderStocks {
				if !order.done {
					cockingCount++
				}
			}
			if cockingCount > k {
				fmt.Println("wait")
			} else {
				fmt.Println(info.menuID)
			}
		} else {
			isInvalid := true
			for _, v := range orderStocks {
				if v.menuID == info.menuID {
					v.done = true
					isInvalid = false
					break
				}
			}
			if isInvalid {
				fmt.Println("unexpected input")
				continue
			}
			nextIndex := 0
			other := true
			for _, v := range orderStocks {
				if !v.done {
					nextIndex++
					if nextIndex == k {
						fmt.Printf("ok %d\n", v.menuID)
						other = false
						break
					}
				}
			}
			if other {
				fmt.Println("ok")
			}
		}
	}
}

func case3(scanner *bufio.Scanner) {
	scanner.Scan()
	m, err := strconv.ParseInt(scanner.Text(), 10, 64)
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
			id:    id,
			stock: stock,
			value: value,
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
				menuID: menuID,
				seat:   seat,
			}
		} else {
			menuID, _ := strconv.ParseInt(args[1], 10, 64)
			info = receiveInfo{
				state:  args[0],
				menuID: menuID,
			}
		}
		receiveInfos = append(receiveInfos, info)
	}

	type o struct {
		menuID int64
		seat   int64
	}
	var stack []o
	var head int
	for _, info := range receiveInfos {
		if info.state == "received" {
			a := o{
				menuID: info.menuID,
				seat:   info.seat,
			}
			stack = append(stack, a)
		} else {
			for i, s := range stack {
				if i < head {
					continue
				}
				if s.menuID == info.menuID {
					fmt.Printf("ready %d %d\n", s.seat, s.menuID)
					head++
				}
			}
		}
	}
}
