# [236. 二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)


## 题目

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

根据维基百科对[最近公共祖先](https://en.wikipedia.org/wiki/Lowest_common_ancestor)的定义:“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”


**Example 1:**

![](https://assets.leetcode.com/uploads/2018/12/14/binarytree.png)

```
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1

输出：3

解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2018/12/14/binarytree.png)

```
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4

输出：5

解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
```

**Example 3:**

```
输入：root = [1,2], p = 1, q = 2

输出：1
```

**Note:**

- 树中节点数目在范围 [2, 105] 内。
- -109 <= Node.val <= 109
- 所有 Node.val 互不相同 。
- p != q
- p 和 q 均存在于给定的二叉树中。

## 题目大意

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”


## 解题思路

- 这是一套经典的题目，寻找任意一个二叉树中两个结点的 LCA(Lowest common ancestor) 最近公共祖先，考察递归。
- 找到终止条件：即递归找到p、q点，然后回溯返回p或q，或者没有找到p、q，但root == null了，这时返回null。
- 找到返回条件：1.如果先找到了p，右侧是null（说明q是p的子节点），那么就会一直回溯p值，最后得到的就是p值。2.如果先找到了q，左侧是null（说明p是q的子节点），那么就会一直回溯q值，最后得到的就是q值。3.如果p和q都找到了，那么就会一直回溯他们父节点，最后得到的就是他们的父节点。


## 代码

**Java版**

```java
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {

        if (root == p || root == q || root == null) {
            return root;
        }

        TreeNode left = lowestCommonAncestor(root.left, p, q);
        TreeNode right = lowestCommonAncestor(root.right, p, q);
        if (left == null) {
            return right;
        }
        if (right == null) {
            return left;
        }
        return root;
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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == q || root == p || root == nil {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
```

- 因为测试用例编写太复杂，此题没有给单元测试文件