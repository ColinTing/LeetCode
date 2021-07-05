# [77. 组合](https://leetcode-cn.com/problems/combinations/)


## 题目

给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

**Example:**

```
输入: n = 4, k = 2

输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
```


## 题目大意

给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

## 解题思路

- 计算排列组合中的组合，用 DFS 深搜即可，注意剪枝


## 代码

**Java版**

```java
class Solution {
    public List<List<Integer>> combine(int n, int k) {
        List<List<Integer>> res = new ArrayList<>();
        List<Integer> currCombine = new ArrayList<>();
        combineDFS(res, currCombine, 1, n, k);
        return res;
    }

    private void combineDFS(List<List<Integer>> res, List<Integer> currCombine, int start, int n, int k) {

        if (k == 0) {
            res.add(new ArrayList<>(currCombine));
            return;
        }
        for (int i = start; i <= n; i++) {

            currCombine.add(i);
            combineDFS(res, currCombine, i + 1, n, k - 1);
            currCombine.remove(currCombine.size() - 1);
        }
        return;
    }
}
```

**Golang版**

```go
package leetcode

func combine(n int, k int) [][]int {
	res, c :=  [][]int{}, []int{}

	combineDFS(&res, c, 1, n, k)
	return res
}

func combineDFS(res *[][]int, c []int, start int, n int, k int) {
	if k == 0 {
		temp := make([]int, len(c))
		copy(temp, c)
		*res = append(*res, temp)
		return
	}

	for i := start; i <= n; i++ {
		c = append(c, i)
		combineDFS(res, c, i + 1, n, k - 1)
		c = c[:len(c) - 1]
	}
	return
 }
```