package leetcode

import "github.com/ColinTing/LeetCode/structures"

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
func buildTree(inorder []int, postorder []int) *TreeNode {
	return buildTreeDFS(inorder, 0, len(inorder)-1, postorder, 0)
}

func buildTreeDFS(inorder []int, inStart int, inEnd int, postorder []int, postStart int) *TreeNode {
	if postStart > len(postorder)-1 || inStart > inEnd {
		return nil
	}

	root := &TreeNode{Val: postorder[postStart+inEnd-inStart]}

	inIdx := 0
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == postorder[postStart+inEnd-inStart] {
			inIdx = i
			break
		}
	}

	root.Left = buildTreeDFS(inorder, inStart, inIdx-1, postorder, postStart)
	root.Right = buildTreeDFS(inorder, inIdx+1, inEnd, postorder, postStart+inIdx-inStart)
	return root
}
