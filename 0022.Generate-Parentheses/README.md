# [22. 括号生成](https://leetcode-cn.com/problems/generate-parentheses/)


## 题目

数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。


**Example 1:**

```
输入：n = 3

输出：
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
``` 

**Example 2:**

```
输入：n = 1

输出：["()"]
```

## 题目大意

给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。


## 解题思路

- 这道题乍一看需要判断括号是否匹配的问题，如果真的判断了，那时间复杂度就到 O(n * 2^n)了，虽然也可以 AC，但是时间复杂度巨高。
- 这道题实际上不需要判断括号是否匹配的问题。因为在 DFS 回溯的过程中，会让 `(` 和 `)` 成对的匹配上的。
- open, closed 代表剩余的左右括号
- 如果在某次递归时，剩余的左括号的个数大于剩余的右括号的个数，说明此时生成的字符串中右括号的个数大于左括号的个数，即会出现 ')(' 这样的非法串，所以这种情况直接返回，不继续处理。如果 open 和 closed 都为0，则说明此时生成的字符串已有3个左括号和3个右括号，且字符串合法，则存入结果中后返回。如果以上两种情况都不满足，若此时 open 大于0，则调用递归函数，注意参数的更新，若 closed 大于0，则调用递归函数，同样要更新参数


## 代码


**Java版**

```java
class Solution {
    public List<String> generateParenthesis(int n) {
        List<String> res = new ArrayList<>();

        generateParenthesisDFS(res, "", n, n);

        return res;
    }

    private void generateParenthesisDFS(List<String> res, String s, int open, int closed) {

        if (open > closed) {
            return;
        }

        if (open == 0 && closed == 0) {
            res.add(s);

        } else {
            if (open > 0) {
                generateParenthesisDFS(res, s + '(', open - 1, closed);
            }

            if (closed > 0) {
                generateParenthesisDFS(res, s + ')', open, closed - 1);
            }
        }

    }
}
```

**Golang版**

```go
package leetcode

func generateParenthesis(n int) []string {

	res := []string{}
	generateParenthesisDFS(&res, "", n, n)
	return res
}

func generateParenthesisDFS(res *[]string, s string, open int, closed int) {
	if open > closed {
		return
	}

	if open == 0 && closed == 0 {
		*res = append(*res, s)
	} else {
		if open > 0 {
			generateParenthesisDFS(res, s+"(", open-1, closed)
		}

		if closed > 0 {
			generateParenthesisDFS(res, s+")", open, closed-1)
		}
	}
}
```