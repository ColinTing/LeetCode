package leetcode

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

func rotate1(nums []int, k int) {
	newNums := make([]int, len(nums))
	k %= len(nums)
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}

	copy(nums, newNums)
}
