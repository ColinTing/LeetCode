# [113. 路径总和 II](https://leetcode-cn.com/problems/path-sum-ii/)


## 题目

给你二叉树的根节点 `root` 和一个整数目标和 `targetSum` ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

**叶子节点** 是指没有子节点的节点。

**Example:**

```

          5
         / \
        4   8
       /   / \
      11  13  4
     /  \    / \
    7    2  5   1

输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22

输出：
    [
       [5,4,11,2],
       [5,8,4,5]
    ]
```
## 题目大意

给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。说明: 叶子节点是指没有子节点的节点。

## 解题思路

- 这一题是第 `112 题` 的增强版

## 代码

**Java版**

```java
class Solution {
    public List<List<Integer>> pathSum(TreeNode root, int targetSum) {

        List<List<Integer>> res = new ArrayList<>();
        dfs(res, new ArrayList<Integer>(), root, targetSum);
        return res;
    }

    public void dfs(List<List<Integer>> res, List<Integer> list, TreeNode root, int targetSum) {
        if (root == null) {
            return;
        }
        list.add(root.val);
        if (root.left == null && root.right == null && root.val == targetSum) {
            res.add(new ArrayList<>(list));
        }
        dfs(res, list, root.left, targetSum - root.val);
        dfs(res, list, root.right, targetSum - root.val);
        list.remove(list.size() - 1);
    }
}
```

**Golang版**

```go
package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

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
```