# [47. 全排列 II](https://leetcode-cn.com/problems/permutations-ii/)


## 题目

给定一个可包含重复数字的序列 `nums` ，**按任意顺序** 返回所有不重复的全排列。

**Example:**

```
输入: [1,1,2]

输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
```

## 题目大意

给定一个可包含重复数字的序列，返回所有不重复的全排列。

## 解题思路

- 这一题是第 46 题的加强版，第 46 题中求数组的排列，数组中元素不重复，但是这一题中，数组元素会重复，所以最终排列出来的结果需要去重。
- 去重的方法，将数组排序以后，判断重复元素：如果遇到重复元素（且回溯处理后used[i - 1] == false），扼杀他的排列请求。
- 其他思路和第 46 题完全一致，DFS 深搜即可。


## 代码

**Java版**

```java
class Solution {
    public List<List<Integer>> permuteUnique(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        List<Integer> pu = new ArrayList<>();
        int[] used = new int[nums.length];
        Arrays.sort(nums);
        permuteUniqueDFS(res, pu, nums, used, nums.length);
        return res;
    }

    private void permuteUniqueDFS(List<List<Integer>> res, List<Integer> pu, int[] nums, int[] used, int n) {
        if (n == 0) {
            res.add(new ArrayList<>(pu));
            return;
        }
        for (int i = 0; i < nums.length; i++) {
            if (used[i] == 1) {
                continue;
            }
            if (i > 0 && nums[i - 1] == nums[i] && used[i - 1] == 0) {
                continue;
            }
            pu.add(nums[i]);
            used[i] = 1;
            permuteUniqueDFS(res, pu, nums, used, n - 1);
            used[i] = 0;
            pu.remove(pu.size() - 1);
        }
        return;
    }
}
```

**Golang版**

```go
package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	res, p, used := [][]int{}, []int{}, make([]bool, len(nums))
	sort.Ints(nums)
	permuteUniqueDFS(&res, p, nums, &used, len(nums))
	return res
}

func permuteUniqueDFS(res *[][]int, p []int, nums []int, used *[]bool, n int) {
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
		if i > 0 && nums[i] == nums[i - 1] && (*used)[i - 1] == false {
			continue
		}
		p = append(p, nums[i])
		(*used)[i] = true
		permuteUniqueDFS(res, p, nums, used, n - 1)
		(*used)[i] = false
		p = p[:len(p)-1]
	}
	return
}
```