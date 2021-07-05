package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	res, p, used := [][]int{}, []int{}, make([]bool, len(nums))
	sort.Ints(nums)
	permuteUniqueDFS(&res, p, nums, &used, len(nums))
	return res
}

func permuteUniqueDFS(res *[][]int, p []int, nums []int, used *[]bool, n int) {
	if n == 0 {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if (*used)[i] == true {
			continue
		}
		if i > 0 && nums[i] == nums[i - 1] && (*used)[i - 1] == false {
			continue
		}
		p = append(p, nums[i])
		(*used)[i] = true
		permuteUniqueDFS(res, p, nums, used, n - 1)
		(*used)[i] = false
		p = p[:len(p)-1]
	}
	return
}
