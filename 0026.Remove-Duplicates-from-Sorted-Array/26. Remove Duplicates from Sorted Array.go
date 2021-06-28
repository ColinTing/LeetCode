package leetcode

//解法一

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last := 0
	for finder := 1; finder < len(nums); finder++ {
		if nums[finder] != nums[last] {
			last++
			nums[last] = nums[finder]
		}
	}
	return last + 1
}
