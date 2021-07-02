# [589. N 叉树的前序遍历](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/)

## 题目

给定一个 N 叉树，返回其节点值的 **前序遍历** 。

N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 `null` 分隔（请参见示例）。


**进阶：**

递归法很简单，你可以使用迭代法完成此题吗?


**Example 1:**

![](https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png)

```
输入：root = [1,null,3,2,4,null,5,6]

输出：[1,3,5,6,2,4]
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png)

```
输入：root = 
[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]

输出：[1,2,3,6,7,11,14,4,8,12,5,9,13,10]
```

**提示：**

- N 叉树的高度小于或等于 1000
- 节点总数在范围 [0, 10^4] 内


## 题目大意

给定一个 N 叉树，返回其节点值的 **前序遍历** 。N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 `null` 分隔（请参见示例）。

## 解题思路

- 递归解法非常简单，前序遍历就是指根节点在左右节点之前遍历，中序遍历就是指根节点在中间遍历，后序遍历就是指根节点在最后遍历。

## 代码

**Java版**

```java
class Solution {

    public List<Integer> list = new ArrayList<>();

    public List<Integer> preorder(Node root) {
        if (root == null) {
            return list;
        }

        list.add(root.val);

        for (Node child : root.children) {
            preorder(child);
        }

        return list;
    }
}
```

**Golang版**

```go
package leetcode

type Node struct {
	Val      int
	Children []*Node
}

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
```