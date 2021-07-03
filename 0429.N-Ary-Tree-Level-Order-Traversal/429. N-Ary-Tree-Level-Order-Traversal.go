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
func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	var result [][]int
	for l := len(queue); len(queue) != 0; l = len(queue) {
		var level []int
		for i := 0; i < l; i++ {
			level = append(level, queue[i].Val)
			queue = append(queue, queue[i].Children...)
		}
		queue = queue[l:]
		result = append(result, level)
	}
	return result
}
