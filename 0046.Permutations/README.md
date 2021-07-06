# [46. 全排列](https://leetcode-cn.com/problems/permutations/)


## 题目

给定一个不含重复数字的数组 `nums` ，返回其 **所有可能的全排列** 。你可以 **按任意顺序** 返回答案。

**Example:**

```
输入: [1,2,3]

输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
```

## 题目大意

给定一个没有重复数字的序列，返回其所有可能的全排列。


## 解题思路

- 求出一个数组的排列组合中的所有排列，用 DFS 深搜即可。
- 类似题目 `77. 组合`， 只不过这里for循环不用设定变动的起始值，且标记重复访问项


## 代码

**Java版**

```java
class Solution {
    public List<List<Integer>> permute(int[] nums) {

        List<List<Integer>> res = new ArrayList<>();
        List<Integer> currPermute = new ArrayList<>();
        int[] visited = new int[nums.length];
        premuteDFS(res, currPermute, nums, visited, nums.length);

        return res;
    }

    private void premuteDFS(List<List<Integer>> res, List<Integer> currPermute, int[] nums, int[] visited, int n) {
        if (n == 0) {
            res.add(new ArrayList<>(currPermute));
            return;
        }
        for (int i = 0; i < nums.length; i++) {
            if (visited[i] == 1) {
                continue;
            }
            currPermute.add(nums[i]);
            visited[i] = 1;
            premuteDFS(res, currPermute, nums,  visited,n - 1);
            visited[i] = 0;
            currPermute.remove(currPermute.size() - 1);
        }
        return;
    }
}
```

**Golang版**

```go
package leetcode

func permute(nums []int) [][]int {
	res, p, used := [][]int{}, []int{}, make([]bool, len(nums))

	permuteDFS(&res, p, nums, &used, len(nums))
	return res
}

func permuteDFS(res *[][]int, p []int, nums []int, used *[]bool, n int) {
	if n == 0 {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if (*used)[i] == true {
			continue
		}
		p = append(p, nums[i])
		(*used)[i] = true
		permuteDFS(res, p, nums, used, n-1)
		(*used)[i] = false
		p = (p)[:len(p)-1]
	}
	return
}
```