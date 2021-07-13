# [106. 从中序与后序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)

## 题目

根据一棵树的中序遍历与后序遍历构造二叉树。

**注意:**
你可以假设树中没有重复的元素。

例如，给出
```
中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
```
返回如下的二叉树：

    	3
       / \
      9  20
        /  \
       15   7


## 题目大意

根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。


## 解题思路

- 给出 2 个数组，根据 inorder 和 postorder 数组构造一颗树。
- 利用递归思想，从 postorder 可以得到根节点，从 inorder 中得到左子树和右子树。postorder在inorder的当前递归范围内最后一个节点即为根节点。不断的递归直到所有的树都生成完成。画个图就明了了


## 代码

**Java版**

```java
class Solution {
    public TreeNode buildTree(int[] inorder, int[] postorder) {

        return buildTreeDFS(inorder, 0, inorder.length - 1, postorder, 0);
    }

    private TreeNode buildTreeDFS(int[] inorder, int inStart, int inEnd, int[] postorder, int postStart) {
        if (postStart > postorder.length - 1 || inStart > inEnd) {
            return null;
        }

        TreeNode root = new TreeNode(postorder[postStart + inEnd - inStart]);

        int inIdx = 0;
        for (int i = inStart; i <= inEnd; i++) {
            if (inorder[i] == postorder[postStart + inEnd - inStart]) {
                inIdx = i;
                break;
            }
        }


        root.left = buildTreeDFS(inorder, inStart, inIdx - 1, postorder, postStart);
        root.right = buildTreeDFS(inorder, inIdx + 1, inEnd, postorder, postStart + inIdx - inStart);
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
```