# [45. 跳跃游戏 II](https://leetcode-cn.com/problems/jump-game-ii/)


## 题目

给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

假设你总是可以到达数组的最后一个位置。

**Example 1:**

```
输入: [2,3,1,1,4]

输出: 2

解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
```

**Example 2:**

```
输入: [2,3,0,1,4]

输出: 2
```

**提示:**

- `1 <= nums.length <= 1000`
- `0 <= nums[i] <= 10^5`

## 题目大意

给定一个非负整数数组，你最初位于数组的第一个位置。数组中的每个元素代表你在该位置可以跳跃的最大长度。你的目标是使用最少的跳跃次数到达数组的最后一个位置。

## 解题思路

- `55. 跳跃游戏`升级版

- 要求找到最少跳跃次数，顺理成章的会想到用贪心算法解题。扫描步数数组，维护当前能够到达最大下标的位置(maxJumpIdx)，记为能到达的最远边界，维护一个本次跳跃可到达最远的下标位置（curEndIdx），在到达本次跳跃可到达最远的下标位置后，需要形成新的跳跃了，跳跃数+1，更新下次跳跃可到达最远的下标位置（curEndIdx）为当前能够到达最大下标的位置(maxJumpIdx)， 如果扫描过程中到达了最远边界，直接将跳跃次数 + 1返回。
- 扫描数组的时候，其实不需要扫描最后一个元素，因为在跳到最后一个元素之前，能到达的最远边界一定大于等于最后一个元素的位置，不然就跳不到最后一个元素，到达不了终点了；如果遍历到最后一个元素，说明边界正好为最后一个位置，最终跳跃次数直接 + 1 即可，也不需要访问最后一个元素。

## 代码


**Java版**

```java
class Solution {
    public int jump(int[] nums) {
        int maxJumpIdx = 0;
        int curEndIdx = 0;
        int jumps = 0;
        for (int i = 0; i < nums.length - 1; i++) {
            if (i + nums[i] >= nums.length - 1) {
                return jumps + 1;
            } else {
                maxJumpIdx = Math.max(maxJumpIdx, i + nums[i]);
                if (curEndIdx == i) {
                    jumps++;
                    curEndIdx = maxJumpIdx;
                }
            }

        }
        return jumps;
    }
}
```

**Golang版**

```go
package leetcode

func jump(nums []int) int {
	maxJumpIdx, curEndIdx, jumps := 0, 0, 0
	for i, v := range nums[:len(nums)-1] {
		if i+v >= len(nums)-1 {
			return jumps + 1
		} else {
			maxJumpIdx = max(maxJumpIdx, i+v)
			if curEndIdx == i {
				jumps++
				curEndIdx = maxJumpIdx
			}
		}
	}
	return jumps
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
```