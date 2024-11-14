package main

import "fmt"

func solution(input string) string {
	stack := []string{}
	for _, i := range input {
		l := len(stack)
		if string(i) == "(" || string(i) == "{" || string(i) == "[" {
			stack = append(stack, string(i))
		} else {
			if string(i) == ")" {
				if len(stack) > 0 {
					if stack[l-1] == "(" {
						stack = stack[:l-1]
					} else {
						return "no"
					}
				} else {
					return "no"
				}
			} else {
				if string(i) == "}" {
					if len(stack) > 0 {
						if stack[l-1] == "{" {
							stack = stack[:l-1]
						} else {
							return "no"
						}
					} else {
						return "no"
					}
				} else {
					if string(i) == "]" {
						if len(stack) > 0 {
							if stack[l-1] == "[" {
								stack = stack[:l-1]
							} else {
								return "no"
							}
						} else {
							return "no"
						}
					}
				}
			}
		}
	}
	if len(stack) == 0 {
		return "yes"
	} else {
		return "no"
	}
}

func main() {
	var input string
	fmt.Scanln(&input)
	fmt.Print(solution(input))

}
