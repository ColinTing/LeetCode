package leetcode

func isValid(s string) bool {
	n := len(s)
	dp := make([]int, n)
	maxValidLen := 0
	for i := n - 2; i >= 0; i-- {
		a := validFunc(&dp, s, i, '(', ')')
		b := validFunc(&dp, s, i, '[', ']')
		c := validFunc(&dp, s, i, '{', '}')
		maxValidLen = max(a, max(b, c))
	}
	return maxValidLen == len(s)
}

func validFunc(dp *[]int, s string, i int, openByte byte, closeByte byte) int {
	if s[i] == openByte {
		j := i + (*dp)[i+1] + 1
		if j < len(s) && s[j] == closeByte {
			(*dp)[i] = (*dp)[i+1] + 2
			if j+1 < len(s) {
				(*dp)[i] += (*dp)[j+1]
			}
		}
	}
	return (*dp)[i]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
