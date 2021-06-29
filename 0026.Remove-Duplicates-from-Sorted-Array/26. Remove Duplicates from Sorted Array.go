package leetcode

//解法一：双指针

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


//解法二：重复值个数计算法

func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dupCt := 0
	for finder := 1; finder < len(nums); finder++ {
		if nums[finder] == nums[finder - 1] {
			dupCt++
		} else {
			nums[finder - dupCt] = nums[finder]
		}
	}
	return len(nums) - dupCt
}
