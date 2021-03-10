package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func infixToPostfix(exp string) string {
	stack := ItemStack{}
	postfix := ""
	expLen := len(exp)

	for i := 0; i < expLen; i++ {
		char := string(exp[i])
		switch char {
		case " ":
			continue
		case "(":
			stack.Push("(")
		case ")":
			for !stack.IsEmpty() {
				preChar := stack.Top()
				if preChar == "(" {
					stack.Pop()
					break
				}
				postfix += preChar
				stack.Pop()
			}
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			j := i
			digit := ""
			for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ {
				digit += string(exp[j])
			}
			postfix += digit
			i = j - 1
		default:
			for !stack.IsEmpty() {
				top := stack.Top()
				if top == "(" || isLower(top, char) {
					break
				}
				postfix += top
				stack.Pop()
			}
			stack.Push(char)
		}
	}
	for !stack.IsEmpty() {
		postfix += stack.Pop()
	}
	return postfix
}

func calculate(postfix string) int {
	stack := ItemStack{}
	fixLen := len(postfix)

	for i := 0; i < fixLen; i++ {
		nextChar := string(postfix[i])
		if unicode.IsDigit(rune(postfix[i])) {
			stack.Push(nextChar)
		} else {
			num1, _ := strconv.Atoi(stack.Pop())
			num2, _ := strconv.Atoi(stack.Pop())
			switch nextChar {
			case "+":
				stack.Push(strconv.Itoa(num2 + num1))
			case "-":
				stack.Push(strconv.Itoa(num2 - num1))
			case "*":
				stack.Push(strconv.Itoa(num2 * num1))
			case "/":
				stack.Push(strconv.Itoa(num2 / num1))
			}
		}
	}
	ans, _ := strconv.Atoi(stack.Top())
	return ans
}

func isLower(top string, newTop string) bool {
	switch top {
	case "+", "-":
		if newTop == "*" || newTop == "/" {
			return true
		}
	case "(":
		return true
	}
	return false
}

func main() {
	s := "2+4*(7-3)/2"
	s = infixToPostfix(s)
	fmt.Println(calculate(s))
}
