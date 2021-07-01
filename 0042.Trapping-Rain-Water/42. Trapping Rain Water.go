package leetcode

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	i, j := 0, len(height)-1
	maxLeft, maxRight := height[i], height[j]
	res := 0
	for i < j {
		if maxLeft < maxRight {
			if maxLeft > height[i+1] {
				res += maxLeft - height[i+1]
			} else {
				maxLeft = height[i+1]
			}
			i++
		} else {
			if maxRight > height[j-1] {
				res += maxRight - height[j-1]
			} else {
				maxRight = height[j-1]
			}
			j--
		}
	}
	return res
}
