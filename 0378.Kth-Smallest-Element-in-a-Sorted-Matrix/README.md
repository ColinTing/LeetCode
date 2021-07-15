# [378. 有序矩阵中第 K 小的元素](https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/)


## 题目

给你一个 `n x n` 矩阵 `matrix` ，其中每行和每列元素均按升序排序，找到矩阵中第 `k` 小的元素。

请注意，它是 **排序后** 的第 `k` 小元素，而不是第 `k` 个 **不同** 的元素。


**Example:**
```
输入：matrix = [
       [ 1,  5,  9],
       [10, 11, 13],
       [12, 13, 15]
    ],
    k = 8,
    
输出：return 13.

解释：矩阵中的元素为 [1,5,9,10,11,12,13,13,15]，第 8 小元素是 13
```

## 题目大意

给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。请注意，它是排序后的第 k 小元素，而不是第 k 个元素。


说明:
你可以假设 k 的值永远是有效的, 1 ≤ k ≤ n2 。


## 解题思路


- 给出一个行有序，列有序的矩阵(并非是按照下标有序的)，要求找出这个矩阵中第 K 小的元素。注意找的第 K 小元素指的不是 k 个不同的元素，可能存在相同的元素。
- 最容易想到的就解法是优先队列。依次把矩阵中的元素推入到优先队列中。维护一个最大堆，一旦优先队列里面的元素有 k 个了，就算找到结果了。

## 代码

**Java版**

```java
class Solution {
    public int kthSmallest(int[][] matrix, int k) {
        int m = matrix.length;
        int n = matrix[0].length;
        PriorityQueue<Integer> q = new PriorityQueue<>((a, b) -> b - a);
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                q.offer(matrix[i][j]);
                if (q.size() > k) {
                    q.poll();
                }
            }
        }
        return q.peek().intValue();
    }
}
```

- Golang创建最大堆太复杂，暂不提供Golang代码
