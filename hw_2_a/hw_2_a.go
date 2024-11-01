package main

import "fmt"

func solution(nums []int) {
	sum := make([]int, len(nums)+1)
	for i := 1; i < len(nums)+1; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	for j := 1; j < len(nums)+1; j++ {
		fmt.Printf("%d ", sum[j])
	}
}

func main() {
	var n int
	fmt.Scanln(&n)
	nums := make([]int, n)
	for i := range n {
		fmt.Scan(&nums[i])
	}
	solution(nums)
}
