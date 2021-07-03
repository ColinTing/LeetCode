# [11. 盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/)

## 题目

给你 `n` 个非负整数 `a1，a2，...，an`，每个数代表坐标中的一个点 `(i, ai)` 。在坐标内画 `n` 条垂直线，垂直线 `i` 的两个端点分别为 `(i, ai)` 和 `(i, 0)` 。找出其中的两条线，使得它们与 `x` 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器。

**Example 1:**

![](https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/07/25/question_11.jpg)


```
输入：[1,8,6,2,5,4,8,3,7]

输出：49 

解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
```

**Example 2:**

```
输入：height = [1,1]

输出：1
```

**Example 3:**

```
输入：height = [4,3,2,1,4]

输出：16
```

**Example 4:**

```
输入：height = [1,2,1]

输出：2
```

## 题目大意

给出一个非负整数数组 a1，a2，a3，…… an，每个整数标识一个竖立在坐标轴 x 位置的一堵高度为 ai 的墙，选择两堵墙，和 x 轴构成的容器可以容纳最多的水。

## 解题思路


这一题也是对撞指针的思路。首尾分别 2 个指针，每次移动以后都分别判断长宽的乘积是否最大。


## 代码

**Java版**

```java
class Solution {
    public int maxArea(int[] height) {
        int left = 0, right = height.length - 1;
        int res = 0;
        while (left < right) {
            if (height[left] < height[right]) {
                res = Math.max(res, height[left] * (right - left));
                left++;
            } else {
                res = Math.max(res, height[right] * (right - left));
                right--;
            }
        }
        return res;
    }
}
```

**Golang版**

```go
package leetcode

func maxArea(height []int) int {

	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}
```