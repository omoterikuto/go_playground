package main

import (
	"fmt"
	"unicode"
)

func Calc(expr string) int {
	c := &calculator{expr: []rune(expr)}
	return c.calc()
}

type calculator struct {
	expr []rune
	idx  int
}

func (c *calculator) calc() int {
	if c.idx >= len(c.expr) {
		return 0
	}

	x := c.term()
	if c.idx >= len(c.expr) {
		return x
	}

	op := c.expr[c.idx]
	if op != '+' && op != '-' {
		return x
	}
	c.idx++

	if c.idx >= len(c.expr) {
		panic("invalid expression")
	}

	y := c.term()
	switch op {
	case '+':
		return x + y
	case '-':
		return x - y
	}

	panic("invalid operator")
}

func (c *calculator) term() int {
	x := c.factor()
	if c.idx >= len(c.expr) {
		return x
	}

	op := c.expr[c.idx]
	if op != '*' && op != '/' {
		return x
	}
	c.idx++

	y := c.factor()
	switch op {
	case '*':
		return x * y
	case '/':
		return x / y
	}

	panic("invalid operator")
}

func (c *calculator) factor() int {
	switch c.expr[c.idx] {
	case '(':
		c.idx++
		x := c.calc()
		if c.expr[c.idx] != ')' {
			panic("no )")
		}
		c.idx++
		return x
	default:
		return c.number()
	}
}

func (c *calculator) number() int {
	var n int
	for c.idx < len(c.expr) {
		r := c.expr[c.idx]
		if !unicode.IsDigit(r) {
			break
		}
		c.idx++
		n = n*10 + int(r-'0')
	}
	return n
}

func main() {
	fmt.Println(Calc("1+2*3"))
}
