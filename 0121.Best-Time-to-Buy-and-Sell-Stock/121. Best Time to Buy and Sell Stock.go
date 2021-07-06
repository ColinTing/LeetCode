package leetcode

func maxProfit(prices []int) int {

	res, dp := 0, make([]int, len(prices))

	for i := 1; i < len(prices); i++ {
		dp[i] = prices[i] - prices[i-1]
	}

	for i := 1; i < len(dp); i++ {
		if dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}