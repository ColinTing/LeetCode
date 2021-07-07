# [55. 跳跃游戏](https://leetcode-cn.com/problems/jump-game/)


## 题目

给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。

**Example 1**:

```
输入：nums = [2,3,1,1,4]

输出：true

解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
```

**Example 2**:

```
输入：nums = [3,2,1,0,4]

输出：false

解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
```

## 题目大意

给定一个非负整数数组，最初位于数组的第一个位置。数组中的每个元素代表在该位置可以跳跃的最大长度。判断是否能够到达最后一个位置。

## 解题思路

- 给出一个非负数组，要求判断从数组 0 下标开始，能否到达数组最后一个位置。
- 这一题比较简单。如果某一个作为 `起跳点` 的格子可以跳跃的距离是 `n`，那么表示后面 `n` 个格子都可以作为 `起跳点`。可以对每一个能作为 `起跳点` 的格子都尝试跳一次，把 `能跳到最远的距离maxJump` 不断更新。如果可以一直跳到最后，就成功了。如果中间有一个点比 `maxJump` 还要大，说明在这个点和 maxJump 中间连不上了，有些点不能到达最后一个位置。

## 代码

**Java版**

```java
class Solution {
    public boolean canJump(int[] nums) {

        int maxJumpIndex = 0;
        for (int i = 0; i < nums.length; i++) {
            if (i > maxJumpIndex) {
                break;
            }

            if (i + nums[i] >= nums.length - 1) {
                return true;
            } else {
                maxJumpIndex = Math.max(maxJumpIndex, i + nums[i]);
            }
        }
        return false;
    }
}
```

**Golang版**

```go
package leetcode

func canJump(nums []int) bool {
	maxJumpIdx := 0

	for i, v := range nums {
		if i > maxJumpIdx {
			break
		}

		if i+v >= len(nums)-1 {
			return true
		} else {
			maxJumpIdx = max(maxJumpIdx, i+v)
		}
	}
	return false
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
```