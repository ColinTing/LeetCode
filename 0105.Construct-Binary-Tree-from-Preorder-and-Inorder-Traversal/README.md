# [105. 从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)


## 题目

根据一棵树的前序遍历与中序遍历构造二叉树。

**注意:**

你可以假设树中没有重复的元素。

例如，给出

```
前序遍历 preorder = [3,9,20,15,7]

中序遍历 inorder = [9,3,15,20,7]
```

返回如下的二叉树：
```
    3
   / \
  9  20
    /  \
   15   7
```


## 题目大意

根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。


## 解题思路

- 给出 2 个数组，根据 preorder 和 inorder 数组构造一颗树。
- 利用递归思想，从 preorder 可以得到根节点，从 inorder 中得到左子树和右子树。只剩一个节点的时候即为根节点。不断的递归直到所有的树都生成完成，画个图就明了了。

## 代码

**Java版**

```java
class Solution {
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        return buildTreeDFS(preorder, inorder, 0, 0, inorder.length - 1);
    }

    private TreeNode buildTreeDFS(int[] preorder, int[] inorder, int preStart, int inStart, int inEnd) {

        if (preStart > preorder.length - 1 || inStart > inEnd) {
            return null;
        }

        TreeNode root = new TreeNode(preorder[preStart]);

        int inIdx = 0;
        for (int i = inStart; i <= inEnd; i++) {
            if (preorder[preStart] == inorder[i]) {
                inIdx = i;
                break;
            }
        }

        root.left = buildTreeDFS(preorder, inorder, preStart + 1, inStart, inIdx - 1);
        root.right = buildTreeDFS(preorder, inorder, preStart + inIdx - inStart + 1, inIdx + 1, inEnd);

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
```