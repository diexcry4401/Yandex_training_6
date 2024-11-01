package main

import "fmt"

// func solution(nums []int) {
// 	sum := make([]int, len(nums)+1)
// 	for i := 1; i < len(nums)+1; i++ {
// 		sum[i] = sum[i-1] + nums[i-1]
// 	}
// 	fmt.Println(sum[1:])
// }

func main() {
	nK := make([]int, 2)
	for i := 0; i < 2; i++ {
		fmt.Scan(&nK[i])
	}
	n := nK[0]
	nums := make([]int, n)
	for i := range n {
		fmt.Scan(&nums[i])
	}
	fmt.Println(nK)
	fmt.Println(nums)
}
