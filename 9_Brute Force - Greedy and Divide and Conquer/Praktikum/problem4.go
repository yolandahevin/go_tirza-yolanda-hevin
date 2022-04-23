package main

import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BinarySearch(array []int, x int) {
	prefixSum := make([]int, x+1)
	for i := 1; i <= x; i++ {
		prefixSum[i] = prefixSum[i-1] + array[i-1]
	}
	dp := make([][]int, x+2)
	for i := range dp {
		dp[i] = make([]int, x+2)
	}
	for i := 2; i <= x; i++ {
		for j := 1; j+i-1 <= x; j++ {
			s, e, lowest := j, j+i-1, math.MaxInt32
			for k := s; k <= e; k++ {
				lowest = min(lowest, dp[s][k-1]+dp[k+1][e]+prefixSum[e]-prefixSum[s-1]-array[k-1])
			}
			dp[s][e] = lowest
		}
	}
	return
}

func main() {
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)
}
