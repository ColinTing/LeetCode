package leetcode

//解法一:三次旋转法, 时间复杂度 O(n)，空间复杂度 O(1)
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

//解法二：迭代取模赋值法, 时间复杂度 O(n)，空间复杂度 O(n)
func rotate1(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}

	copy(nums, newNums)
}
