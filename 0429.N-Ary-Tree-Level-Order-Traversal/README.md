# [429. N 叉树的层序遍历](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)

## 题目

给定一个 N 叉树，返回其节点值的 **层序遍历**。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。




**Example 1:**

![](https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png)

```
输入：root = [1,null,3,2,4,null,5,6]

输出：[[1],[3,2,4],[5,6]]
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png)

```
输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]

输出：[[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]
```

**提示：**

- 树的高度不会超过 1000
- 树的节点总数在 [0, 10^4] 之间


## 题目大意

给定一个 N 叉树，返回其节点值的 **层序遍历** 。N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 `null` 分隔（请参见示例）。

## 解题思路

- Queue(队列) + List，BFS，类似题目有 `127. 单词接龙`

## 代码

**Java版**

```java
class Solution {
    public List<List<Integer>> levelOrder(Node root) {
        List<List<Integer>> res = new ArrayList<>();
        if (root == null) {
            return res;
        }
        Queue<Node> queue = new LinkedList<>();
        queue.offer(root);
        while (!queue.isEmpty()) {
            List<Integer> currLevelList = new ArrayList<>();
            int levelNodeCt = queue.size();
            for (int i = 0; i < levelNodeCt; i++) {
                Node node = queue.poll();
                currLevelList.add(node.val);
                for (Node child : node.children) {
                    queue.offer(child);
                }
            }
            res.add(currLevelList);
        }
        return res;
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
```