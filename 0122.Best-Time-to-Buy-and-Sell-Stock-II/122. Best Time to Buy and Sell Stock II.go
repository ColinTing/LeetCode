package leetcode

func maxProfit(prices []int) int {

	dp := make([]int, len(prices))
	res := 0
	for i := 1; i < len(dp); i++ {
		dp[i] = prices[i] - prices[i-1]
	}

	for i := 1; i < len(dp); i++ {
		if dp[i] > 0 {
			res += dp[i]
		}
	}
	return res
}
