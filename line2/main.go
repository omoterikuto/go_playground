package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type orderInfo struct {
	ID       string
	menuName string
	size     string
	time     time.Time
	complete string
}

//var C int

func main() {
	file := "data.txt"
	fp, err := os.Open(file)
	if err != nil {
		fmt.Printf("err:%s", err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	scanner.Scan()
	//cc, err := strconv.ParseInt(scanner.Text(), 10, 64)
	//C = int(cc)

	scanner.Scan()
	line2 := strings.Split(scanner.Text(), " ")
	//nor, err := strconv.ParseInt(line2[0], 10, 64)
	//if err != nil {
	//	fmt.Printf("parse err:%s", err)
	//}
	//numberOfRecipe := int(nor)

	nom, err := strconv.ParseInt(line2[1], 10, 64)
	if err != nil {
		fmt.Printf("parse err:%s", err)
	}
	numberOfMachine := int(nom)
	//
	capa, err := strconv.ParseInt(line2[2], 10, 64)
	if err != nil {
		fmt.Printf("parse err:%s", err)
	}
	capacity := int(capa)
	//
	//rt, err := strconv.ParseInt(line2[3], 10, 64)
	//if err != nil {
	//	fmt.Printf("parse err:%s", err)
	//}
	//requiredTime := int(rt)

	var menus []string
	scanner.Scan()
	line3 := strings.Split(scanner.Text(), " ")
	for _, recipe := range line3 {
		menus = append(menus, recipe)
	}

	var orderInfos []orderInfo
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		if len(args) == 4 {
			t, err := time.Parse("2006-01-02 15:04:05 MST", fmt.Sprintf("2022-05-22 %s JST", args[3]))
			if err != nil {
				fmt.Printf("parse err:%s", err)
			}

			info := orderInfo{
				ID:       args[0],
				menuName: args[1],
				size:     args[2],
				time:     t,
			}
			orderInfos = append(orderInfos, info)
		} else {
			info := orderInfo{
				complete: args[1],
			}
			orderInfos = append(orderInfos, info)
		}
	}

	//for i, info := range orderInfos {
	//	fmt.Println(i, info)
	//	fmt.Println(numberOfRecipe, numberOfMachine, capacity, requiredTime)
	//}

	//var usingTime time.Time
	var orderList []orderInfo
	var cocking []string
	nextOrderIndex := 0
	for _, info := range orderInfos {
		if info.complete == "" {
			canAdd := true
			for _, c := range orderList {
				if c.ID == info.ID {
					fmt.Println("ERROR: reference number duplicates.")
					canAdd = false
					break
				}
			}
			if canAdd {
				orderList = append(orderList, info)
				//fmt.Println("append", orderList)
				//fmt.Println("cocking", cocking)
				if len(cocking) <= numberOfMachine {
					cocking = append(cocking, info.ID)
					nextOrderIndex++
					fmt.Printf("make: %d %s %d\n", numberOfMachine, info.menuName, convSizeToML(info.size))
				}
			}
		} else {
			if len(cocking) != 0 {
				fmt.Printf("serve to: ")
				for _, c := range cocking {
					fmt.Printf("%s ", c)
				}
				fmt.Printf("\n")
				fmt.Println("cocking", cocking)
				cocking = cocking[len(cocking):]
			}
			if len(orderList) > nextOrderIndex {
				var addList []string
				addList = append(addList, orderList[nextOrderIndex].ID)
				for _, order := range orderList {
					if order.ID == orderList[nextOrderIndex].ID {
						continue
					}
					if orderList[nextOrderIndex].menuName != order.menuName {
						continue
					}
					if convSizeToML(orderList[nextOrderIndex].size)+convSizeToML(order.size) <= capacity {
						addList = append(addList, order.ID)
						break
					}
					cocking = append(cocking, addList...)
				}
				fmt.Printf("make: %d %s %d\n", numberOfMachine, orderList[nextOrderIndex].menuName, convSizeToML(orderList[nextOrderIndex].size))
			}
		}
	}
}

func convSizeToML(size string) int {
	var intSize int
	switch size {
	case "S":
		intSize = 300
	case "M":
		intSize = 500
	case "L":
		intSize = 700
	}
	return intSize
}
