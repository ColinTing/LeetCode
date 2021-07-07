# [529. 扫雷游戏](https://leetcode-cn.com/problems/minesweeper/)



## 题目

让我们一起来玩扫雷游戏！ ([Wikipedia](https://zh.wikipedia.org/wiki/%E8%B8%A9%E5%9C%B0%E9%9B%B7), [online game](http://minesweeperonline.com/))!

给定一个代表游戏板的二维字符矩阵。 **'M'** 代表一个**未挖出的**地雷，**'E'** 代表一个**未挖出的**空方块，**'B'** 代表没有相邻（上，下，左，右，和所有4个对角线）地雷的**已挖出的**空白方块，**数字**（'1' 到 '8'）表示有多少地雷与这块**已挖出的**方块相邻，**'X'** 则表示一个**已挖出的**地雷。

现在给出在所有**未挖出的**方块中（'M'或者'E'）的下一个点击位置（行和列索引），根据以下规则，返回相应位置被点击后对应的面板：

1. 如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 **'X'**。
2. 如果一个**没有相邻地雷的**空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的**未挖出**方块都应该被递归地揭露。
3. 如果一个**至少与一个地雷相邻**的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
4. 如果在此次点击中，若无更多方块可被揭露，则返回面板。



**Example 1**:

```
输入:  

[['E', 'E', 'E', 'E', 'E'],
 ['E', 'E', 'M', 'E', 'E'],
 ['E', 'E', 'E', 'E', 'E'],
 ['E', 'E', 'E', 'E', 'E']]

Click : [3,0]

输出: 

[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'M', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]

解释:
```

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/minesweeper_example_1.png)


**Example 2**:

```
输入: 

[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'M', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]

Click : [1,2]

解释:

[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'X', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]

Explanation:
```

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/minesweeper_example_2.png)

**注意：**

1. 输入矩阵的宽和高的范围为 [1,50]。
2. 点击的位置只能是未被挖出的方块 ('M' 或者 'E')，这也意味着面板至少包含一个可点击的方块。
3. 输入面板不会是游戏结束的状态（即有地雷已被挖出）。
4. 简单起见，未提及的规则在这个问题中可被忽略。例如，当游戏结束时你不需要挖出所有地雷，考虑所有你可能赢得游戏或标记方块的情况。


## 题目大意

题目描述很详细就不复述了


## 解题思路

- 依题意解即可
- 分两种情况，点击的点是雷和不是雷，不是雷又分两种情况，一个是周围有雷，一个是周围无雷，再对无雷的点的周边进行递归探索（DFS）即可


## 代码

**Java版**

```java
class Solution {
    public char[][] updateBoard(char[][] board, int[] click) {
        int m = board.length;
        int n = board[0].length;

        if (m == 0 || n == 0) {
            return new char[0][0];
        }

        int row = click[0];
        int col = click[1];
        int ct = 0;
        if (board[row][col] == 'M') {
            board[row][col] = 'X';
        } else {
            for (int i = -1; i < 2; i++) {
                for (int j = -1; j < 2; j++) {
                    int x = row + i;
                    int y = col + j;
                    if (x < 0 || x >= m || y < 0 || y >= n || (i == 0 && j == 0)) {
                        continue;
                    }
                    if (board[x][y] == 'M') {
                        ct++;
                    }
                }
            }

            if (ct > 0) {
                board[row][col] = (char) ('0' + ct);
            } else {
                board[row][col] = 'B';
                for (int i = -1; i < 2; i++) {
                    for (int j = -1; j < 2; j++) {
                        int x = row + i;
                        int y = col + j;
                        if (x < 0 || x >= m || y < 0 || y >= n || (i == 0 && j == 0)) {
                            continue;
                        }
                        if (board[x][y] == 'E') {
                            int[] newClick = new int[]{x, y};
                            updateBoard(board, newClick);
                        }
                    }
                }
            }
        }
        return board;
    }
}
```

**Golang版**

```go
package leetcode

func updateBoard(board [][]byte, click []int) [][]byte {
	m, n := len(board), len(board[0])

	if m == 0 || n == 0 {
		return nil
	}

	row, col := click[0], click[1]
	ct := 0
	if board[row][col] == 'M' {
		board[row][col] = 'X'
	} else {
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				x := row + i
				y := col + j
				if x < 0 || x >= m || y < 0 || y >= n ||
					(i == 0 && j == 0) {
					continue
				}
				if board[x][y] == 'M' {
					ct++
				}
			}
		}

		if ct > 0 {
			board[row][col] = byte('0' + ct)
		} else {
			board[row][col] = 'B'
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					x := row + i
					y := col + j
					if x < 0 || x >= m || y < 0 || y >= n ||
						(i == 0 && j == 0) {
						continue
					}
					if board[x][y] == 'E' {
						newClick := []int{x, y}
						updateBoard(board, newClick)
					}
				}
			}
		}
	}
	return board
}
```