package leetcode

import (
	"github.com/ColinTing/LeetCode/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeDFS(preorder, inorder, 0, 0, len(inorder)-1)
}

func buildTreeDFS(preorder []int, inorder []int, preStart int, inStart int, inEnd int) *TreeNode {

	if preStart > len(preorder)-1 || inStart > inEnd {
		return nil
	}

	root := &TreeNode{Val: preorder[preStart]}
	inorderIdx := 0
	for i, v := range inorder {
		if v == preorder[preStart] {
			inorderIdx = i
			break
		}
	}
	root.Left = buildTreeDFS(preorder, inorder, preStart+1, inStart, inorderIdx-1)
	root.Right = buildTreeDFS(preorder, inorder, preStart+inorderIdx-inStart+1, inorderIdx+1, inEnd)
	return root
}
