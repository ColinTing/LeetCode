package leetcode

func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, n)
	res := 0
	for i := n - 2; i >= 0; i-- {
		if s[i] == '(' {
			j := i + dp[i+1] + 1
			if j < n && s[j] == ')' {
				dp[i] += dp[i+1] + 2
				if j+1 < n {
					dp[i] += dp[j+1]
				}
			}
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
