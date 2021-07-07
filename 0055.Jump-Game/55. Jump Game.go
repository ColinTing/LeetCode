package leetcode

func canJump(nums []int) bool {
	maxJumpIdx := 0

	for i, v := range nums {
		if i > maxJumpIdx {
			break
		}

		if i+v >= len(nums)-1 {
			return true
		} else {
			maxJumpIdx = max(maxJumpIdx, i+v)
		}
	}
	return false
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
