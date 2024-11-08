package main

import "fmt"

func findPairsWithPrefixSum(arr []int, target int) {
	prefixSum := 0
	countPair := 0
	prefixMap := make(map[int][]int)
	prefixMap[0] = []int{-1}
	fmt.Println(prefixMap)

	for i, num := range arr {
		fmt.Println(prefixSum, num)
		prefixSum += num
		_, exists := prefixMap[prefixSum-target]
		if exists {
			fmt.Printf("prefixMap[%d-%d]\n", prefixSum, target)
			countPair += 1
		}
		prefixMap[prefixSum] = append(prefixMap[prefixSum], i)
		fmt.Println(prefixMap)
	}
	fmt.Println(countPair)
}

func main() {
	// nK := make([]int, 2)
	// for i := 0; i < 2; i++ {
	// 	fmt.Scan(&nK[i])
	// }
	// n := nK[0]
	// nums := make([]int, n)
	// for i := range n {
	// 	fmt.Scan(&nums[i])
	// }
	nK := [2]int{5, 17}
	nums := []int{17, 7, 10, 7, 10}
	findPairsWithPrefixSum(nums, nK[1])
}
