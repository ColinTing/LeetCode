# [42. 接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)

## 题目

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。


**Example**:

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rainwatertrap.png)


```
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]

输出：6

解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
```

## 题目大意

从 x 轴开始，给出一个数组，数组里面的数字代表从 (0,0) 点开始，宽度为 1 个单位，高度为数组元素的值。如果下雨了，问这样一个容器能装多少单位的水？

## 解题思路

- 记录左右两边（首尾开始计算）的最高高度，移动高度更低那边的指针。计算当前高度与此边最高高度的差值，如果当前高度更低，累加差值进结果中，反之则更新此边最高高度。下一次循环开始，比较两边最高高度，移动更低边指针...

## 代码

**Java版**

```java
class Solution {
    public int trap(int[] height) {
        if (height.length == 0) {
            return 0;
        }
        int i = 0;
        int j = height.length - 1;
        int maxLeft = height[i];
        int maxRight = height[j];
        int res = 0;
        while (i < j) {
            if (maxLeft < maxRight) {
                if (height[i + 1] < maxLeft) {
                    res += maxLeft - height[i + 1];
                } else {
                    maxLeft = height[i + 1];
                }
                i++;
            } else {
                if (height[j - 1] < maxRight) {
                    res += maxRight - height[j - 1];
                } else {
                    maxRight = height[j - 1];
                }
                j--;
            }
        }
        return res;
    }
}
```

**Golang版**

```go
package leetcode

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	i, j := 0, len(height)-1
	maxLeft, maxRight := height[i], height[j]
	res := 0
	for i < j {
		if maxLeft < maxRight {
			if maxLeft > height[i+1] {
				res += maxLeft - height[i+1]
			} else {
				maxLeft = height[i+1]
			}
			i++
		} else {
			if maxRight > height[j-1] {
				res += maxRight - height[j-1]
			} else {
				maxRight = height[j-1]
			}
			j--
		}
	}
	return res
}
```