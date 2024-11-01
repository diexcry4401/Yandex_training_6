package main

import "fmt"

func solution(nums []int) {
	sum := make([]int, len(nums)+1)
	for i := 1; i < len(nums)+1; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	fmt.Println(sum)
}

func countprefixsums(nums []int) {
	prefixsumbyvalue := make(map[int]int)
	prefixsumbyvalue[0] = 1
	nowsum := 0
	for _, val := range nums {
		nowsum += val

		_, ok := prefixsumbyvalue[nowsum]
		if !ok {
			prefixsumbyvalue[nowsum] = 0
		}
		prefixsumbyvalue[nowsum] += 1
	}
	countneededsumranges(prefixsumbyvalue)
}

func countneededsumranges(prefixsumbyvalue map[int]int) {
	cntranges := 0
	for i := range prefixsumbyvalue {
		cntsum := prefixsumbyvalue[i]
		cntranges += cntsum * (cntsum - 1) / 2
	}
	fmt.Println(cntranges)
}

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
	// fmt.Println(nK)
	// fmt.Println(nums)
	countprefixsums(nums)
}
