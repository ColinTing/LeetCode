# [41. 缺失的第一个正数](https://leetcode-cn.com/problems/first-missing-positive/description/)

## 题目

给你一个未排序的整数数组 `nums` ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 `O(n)` 并且只使用**常数**级别额外空间的解决方案。

Example 1:  

```
输入：nums = [1,2,0]

输出：3  
```

Example 2:  

```
输入：nums = [3,4,-1,1]
  
输出：2
```

Example 3:  

```
输入：nums = [7,8,9,11,12]

输出：1 
```


## 题目大意

找到缺失的第一个正整数。

## 解题思路


- 要求时间复杂度O(n),空间复杂度O(1),那么只能原地替换值了
- 原地替换，nums[nums[i] - 1] = nums[i],直到长度范围内的值都匹配到了应该在的位置时或者匹配到了超出范围的值，这时才后移一位


## 代码

**Java版**

```java
class Solution {
    public int firstMissingPositive(int[] nums) {
        int n = nums.length;
        int i = 0;
        while (i < n) {

            if (nums[i] > 0 && nums[i] <= n && nums[nums[i] - 1] != nums[i]) {
                int temp = nums[nums[i] - 1];
                nums[nums[i] - 1] = nums[i];
                nums[i] = temp;
            } else {
                i++;
            }
        }

        for (int j = 0; j < n; j++) {
            if (nums[j] != j + 1) {
                return j + 1;
            }
        }
        return n + 1;
    }
}
```

**Golang版**

```go
package leetcode

func firstMissingPositive(nums []int) int {
	i := 0
	n := len(nums)
	for i < n {
		if nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}
	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}
	return n + 1
}
```