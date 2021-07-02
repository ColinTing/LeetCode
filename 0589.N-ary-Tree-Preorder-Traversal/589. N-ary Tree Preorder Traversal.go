package leetcode

import (
	"github.com/ColinTing/LeetCode/structures"
)

type Node = structures.Node


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func preorder(root *Node) []int {
	res := []int{}
	preorderDfs(root, &res)
	return res
}

func preorderDfs(root *Node, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	for _, child := range root.Children {
		preorderDfs(child, res)
	}
}
