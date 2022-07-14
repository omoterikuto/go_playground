package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var num1 []string
var num2 []string
var num3 []string

func main() {
	file := "data.txt"
	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	scanner.Scan()

	n := scanner.Text()
	if n == "" {
		fmt.Println("-1")
		return
	}

	nn := strings.Split(n, ".")

	if nn[0][0:1] == "0" {
		fmt.Println("-1")
		return
	}

	number, err := strconv.Atoi(nn[0])
	if err != nil {
		fmt.Println("-1")
		return
	}

	var point int
	if len(nn) > 1 {
		point, err = strconv.Atoi(nn[1])
		if err != nil {
			panic(err)
		}
	}

	var answer string

	num1 = []string{
		"zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine",
		"ten", "eleven", "twelve", "thirteen", "fourteen",
		"fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty",
	}
	num2 = []string{
		"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
	}
	num3 = []string{
		"hundred",
		"thousand",
		"million",
		"billion",
	}

	answer = calc(number)

	if point != 0 {
		answer += " point "
		answer += calcPoint(point)
	}

	answer = strings.Replace(answer, answer[0:1], strings.ToUpper(answer[0:1]), 1)

	fmt.Println(answer)
}

func calc(number int) string {
	var answer string
	if number <= 20 {
		answer = num1[number]
	} else if number <= 99 {
		answer = calc10(number)
	} else if number <= 999 {
		answer += num1[int(math.Floor(float64(number/100)))]
		answer += " "
		answer += num3[0]
		answer += " "
		t := number % 100
		if t == 0 {
			return answer
		}
		answer += num2[int(math.Floor(float64(t/10)))]
		answer += " "
		if t%10 == 0 {
			return answer
		}
		answer += num1[t%10]
	} else if number <= 99999 {
		answer += calc10(int(math.Floor(float64(number / 1000))))
		answer += " "
		answer += num3[1]
		answer += " "
		s := number % 1000
		answer += num1[int(math.Floor(float64(s/100)))]
		answer += " "
		answer += num3[0]
		answer += " "
		t := number % 100
		answer += num2[int(math.Floor(float64(t/10)))]
		answer += " "
		answer += num1[t%10]
	}

	return answer
}

func calcPoint(number int) string {
	strNumber := strconv.Itoa(number)
	nums := strings.Split(strNumber, "")
	var answer string
	for _, num := range nums {
		intNum, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		answer += num1[intNum]
		answer += " "
	}
	return answer
}

func calc10(number int) string {
	var answer string
	answer += num2[int(math.Floor(float64(number/10)))]
	answer += " "
	answer += num1[number%10]
	return answer
}
