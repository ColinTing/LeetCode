package leetcode

// 解法一：双指针法 （不断的用 i，j 标记 0 和非 0 的元素，然后相互交换，最终到达题目的目的）
func moveZeroes(nums []int) {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] == 0 {
			if i != j {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					i++
				}
			}
		} else {
			i++
		}
	}
}

// 解法二：零值个数计算法（当前替换位可通过index减去当前累计零值个数获得）
func moveZeroes1(nums []int) {
	zeroCt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCt++
		} else {
			nums[i-zeroCt], nums[i] = nums[i], nums[i-zeroCt]
		}
	}
}
