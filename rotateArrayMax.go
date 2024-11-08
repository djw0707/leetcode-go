package main

import "fmt"

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	ans := 1 - 1<<32
	for i := 0; i < n; i++ {
		fmt.Println("end index", i)
		dp = make([]int, n)
		for j := i + 1; j <= i+n; j++ {
			fmt.Println("current index", j)
			if j == i+1 {
				dp[j%n] = nums[j%n]
			} else {
				if dp[(j-1+n)%n] < 0 {
					dp[j%n] = nums[j%n]
				} else {
					dp[j%n] = nums[j%n] + dp[(j-1+n)%n]
				}
			}
			ans = max(ans, dp[j%n])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ans := maxSubarraySumCircular([]int{5, -2, 5})
	fmt.Println(ans)
}
