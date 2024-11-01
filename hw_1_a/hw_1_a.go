package main

import "fmt"

func main() {
	var x1, x2, y1, y2, x, y int
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)
	fmt.Scan(&x)
	fmt.Scan(&y)

	if x < x2 {
		if x > x1 {
			if y > y2 {
				fmt.Println("N")
			} else {
				fmt.Println("S")
			}
		} else {
			if y > y2 {
				fmt.Println("NW")
			} else {
				if y > y1 {
					fmt.Println("W")
				} else {
					fmt.Println("SW")
				}
			}
		}
	} else {
		if y > y2 {
			fmt.Println("NE")
		} else {
			if y > y1 {
				fmt.Println("E")
			} else {
				fmt.Println("SE")
			}
		}
	}
}
