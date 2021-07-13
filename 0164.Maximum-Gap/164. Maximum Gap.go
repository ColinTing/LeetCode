package leetcode

func maximumGap(nums []int) int {
	res := 0
	quickSort(nums, 0, len(nums)-1)
	for i := 1; i < len(nums); i++ {
		if res < nums[i]-nums[i-1] {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}

func quickSort(nums []int, start int, end int) {
	if start >= end {
		return
	}
	p := partition(nums, start, end)
	quickSort(nums, start, p-1)
	quickSort(nums, p+1, end)
}

func partition(nums []int, start int, end int) int {

	pivot := nums[start]
	i := start
	for j := i + 1; j <= end; j++ {
		if pivot > nums[j] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[start] = nums[start], nums[i]
	return i
}
