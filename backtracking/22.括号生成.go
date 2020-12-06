package leetcode22t3

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (76.03%)
 * Likes:    1254
 * Dislikes: 0
 * Total Accepted:    167.4K
 * Total Submissions: 220.1K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例：
 *
 * 输入：n = 3
 * 输出：[
 * ⁠      "((()))",
 * ⁠      "(()())",
 * ⁠      "(())()",
 * ⁠      "()(())",
 * ⁠      "()()()"
 * ⁠    ]
 *
 *
 */

// @lc code=start
func generateParenthesis(n int) []string {
	ans := make([]string, 0, 10)

	s := make([]byte, 0, 2*n)

	var backtracking func(int, int, int)
	backtracking = func(total, left, right int) {
		if total == 2*n {
			ans = append(ans, string(s))
		}
		if left < n {
			s = append(s, '(')
			backtracking(total+1, left+1, right)

			s = s[:len(s)-1]
		}

		if left > right {
			s = append(s, ')')
			backtracking(total+1, left, right+1)
			s = s[:len(s)-1]
		}

	}
	backtracking(0, 0, 0)
	return ans
}

// @lc code=end
