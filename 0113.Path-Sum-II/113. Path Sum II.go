package leetcode

import (
	"github.com/ColinTing/LeetCode/structures"
)

type TreeNode = structures.TreeNode

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var list []int
	findPath(root, targetSum, &res, &list)
	return res
}

func findPath(root *TreeNode, targetSum int, res *[][]int, list *[]int) {
	if root == nil {
		return
	}

	*list = append(*list, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		temp := make([]int, len(*list))
		copy(temp, *list)
		*res = append(*res, temp)
	}
	findPath(root.Left, targetSum-root.Val, res, list)
	findPath(root.Right, targetSum-root.Val, res, list)
	*list = (*list)[:len(*list)-1]
	return
}
