package leetcode

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	maxSideLen := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
		}
		maxSideLen = max(maxSideLen, dp[i][0])
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
		}
		maxSideLen = max(maxSideLen, dp[0][j])
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			maxSideLen = max(maxSideLen, dp[i][j])
		}
	}
	return maxSideLen * maxSideLen
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
