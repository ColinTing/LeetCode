package leetcode

func countSubstrings(s string) int {
	res := 0
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 2 || dp[i+1][j-1] {
					res++
					dp[i][j] = true
				}
			}
		}
	}
	return res
}
