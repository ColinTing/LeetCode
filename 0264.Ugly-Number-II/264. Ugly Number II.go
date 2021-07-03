package leetcode

func nthUglyNumber(n int) int {
	dp, p2, p3, p5 := make([]int, n), 0, 0, 0
	dp[0] = 1
	for i := 1; i < n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if x2 == dp[i] {
			p2++
		}

		if x3 == dp[i] {
			p3++
		}

		if x5 == dp[i] {
			p5++
		}
	}
	return dp[n-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
