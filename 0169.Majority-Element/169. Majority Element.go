package leetcode

func majorityElement(nums []int) int {
	res, count := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			res = nums[i]
			count++
		} else if nums[i] == res {
			count++
		} else {
			count--
		}
	}
	return res
}
