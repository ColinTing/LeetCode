package leetcode

func maxProduct(nums []int) int {
	maxNum := nums[0]
	minNum := nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxNum, minNum = minNum, maxNum
		}
		maxNum = max(nums[i], nums[i]*maxNum)
		minNum = min(nums[i], nums[i]*minNum)
		res = max(res, maxNum)
	}
	return res
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
