package stack

import (
	"fmt"
	"strconv"
	"strings"
)

type elem struct {
	typ   int
	value string
}

const (
	op int = iota + 1
	num
	left
	right
)

const (
	add int = iota + 1
	sub
	mul
	div
)

const (
	endExpr = "+-*/"
)

func calc(expr string) int {
	midExpr := toMidExpr(expr)
	suffixExpr := toSuffix(midExpr)
	num := calcStack(suffixExpr)
	return num
}

func calcStack(stac *ArrayStack) int {
	numStack := NewArrayStack()
	for stac.Len() != 0 {
		top := stac.Pop()
		switch top.(elem).typ {
		case num:
			va := top.(elem).value
			vaInt, _ := strconv.Atoi(va)
			numStack.Push(vaInt)
		case op:
			top1 := numStack.Pop().(int)
			top2 := numStack.Pop().(int)
			tmp := 0
			switch top.(elem).value {
			case "+":
				tmp = top1 + top2
			case "-":
				tmp = top1 - top2
			case "*":
				tmp = top1 * top2
			case "/":
				tmp = top2 / top1
			}
			numStack.Push(tmp)
		}
	}
	return numStack.Pop().(int)
}

func toSuffix(ac []elem) *ArrayStack {
	numStack := NewArrayStack()
	opStack := NewArrayStack()
	for _, c := range ac {
		switch c.typ {
		case num:
			numStack.Push(c)
		case left:
			opStack.Push(c)
		case right:
			for opStack.Peek() != nil {
				if opStack.Peek().(elem).typ == left {
					opStack.Pop()
					break
				}
				numStack.Push(opStack.Pop())
			}
		case op:
			if opStack.Len() == 0 {
				opStack.Push(c)
				continue
			}
			if getExprPrority(c.value) > getExprPrority(opStack.Peek().(elem).value) {
				opStack.Push(c)
				continue
			} else {
				for opStack.Peek() != nil && opStack.Peek().(elem).value != "(" && getExprPrority(opStack.Peek().(elem).value) >= getExprPrority(c.value) {
					numStack.Push(opStack.Pop())
				}
				opStack.Push(c)
			}

		default:
			panic("error1")
		}
	}
	for opStack.Len() != 0 {
		numStack.Push(opStack.Pop())
	}
	reStack := NewArrayStack()
	mm := ""
	for numStack.Peek() != nil {
		value := numStack.Pop()
		mm += value.(elem).value
		reStack.Push(value)
	}
	fmt.Println(mm)
	return reStack
}

func getExprPrority(expr string) int {
	re := rune(expr[0])
	if re == 42 || re == 47 {
		return 1
	}
	return 0
}

func toMidExpr(expr string) []elem {

	re := []elem{}
	s := ""
	for _, e := range expr {
		if isEnd(e) {
			if s != "" {
				m := elem{typ: num, value: string(s)}
				re = append(re, m)
				s = ""
			}
			tt := op
			if e == 40 {
				tt = left
			} else if e == 41 {
				tt = right
			}
			m := elem{typ: tt, value: string(e)}
			re = append(re, m)
		} else if isNum(e) {
			s += string(e)
		} else {
			panic("error")
		}
	}
	if s != "" {
		m := elem{typ: num, value: string(s)}
		re = append(re, m)
	}
	return re
}

func isEnd(s rune) bool {
	return strings.Contains(endExpr, string(s)) || s == 40 || s == 41
}

func isNum(s rune) bool {
	return s >= 48 && 57 >= s
}
