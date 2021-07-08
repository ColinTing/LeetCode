package leetcode

func jump(nums []int) int {
	maxJumpIdx, curEndIdx, jumps := 0, 0, 0
	for i, v := range nums[:len(nums)-1] {
		if i+v >= len(nums)-1 {
			return jumps + 1
		} else {
			maxJumpIdx = max(maxJumpIdx, i+v)
			if curEndIdx == i {
				jumps++
				curEndIdx = maxJumpIdx
			}
		}
	}
	return jumps
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
