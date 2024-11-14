package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution(expression string) int {
	stack := []int{}
	tokens := strings.Fields(expression)

	for _, token := range tokens {
		if value, err := strconv.Atoi(token); err == nil {
			stack = append(stack, value)
		} else {
			op2 := stack[len(stack)-1]
			op1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var result int
			switch token {
			case "+":
				result = op1 + op2
			case "-":
				result = op1 - op2
			case "*":
				result = op1 * op2
			default:
				panic("Неизвестный оператор: " + token)
			}
			stack = append(stack, result)
		}
	}
	return stack[0]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)
	result := solution(expression)
	fmt.Println(result)
}
