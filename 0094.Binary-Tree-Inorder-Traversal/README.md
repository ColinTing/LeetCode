# [94. 二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)

## 题目


给定一个二叉树的根节点 `root` ，返回它的 **中序** 遍历。



**Example 1:**

![](https://assets.leetcode.com/uploads/2020/09/15/inorder_1.jpg)

```
输入：root = [1,null,2,3]

输出：[1,3,2]
```

**Example 2:**

```
输入：root = []

输出：[]
```

**Example 3:**

```
输入：root = [1]

输出：[1]
```

**Example 4:**

![](https://assets.leetcode.com/uploads/2020/09/15/inorder_4.jpg)

```
输入：root = [1,null,2]

输出：[1,2]
```

**Example 5:**


![](https://assets.leetcode.com/uploads/2020/09/15/inorder_5.jpg)

```
输入：root = [1,2]

输出：[2,1]
```

**提示：**

- 树中节点数目在范围 [0, 100] 内
- -100 <= Node.val <= 100


**进阶:** 

递归算法很简单，你可以通过迭代算法完成吗？



## 题目大意

中序遍历一颗树。

## 解题思路

递归的实现方法，见代码。


## 代码

**Java版**

```java
class Solution {
    
    List<Integer> list = new ArrayList<>();

    public List<Integer> inorderTraversal(TreeNode root) {
        if (root == null) {
            return list;
        }

        inorderTraversal(root.left);
        list.add(root.val);
        inorderTraversal(root.right);
        return list;
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

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	inorder(root, &res)
	return res
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}
```