package leetcode32

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 *
 * https://leetcode-cn.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (34.31%)
 * Likes:    1125
 * Dislikes: 0
 * Total Accepted:    120.3K
 * Total Submissions: 350.6K
 * Testcase Example:  '"(()"'
 *
 * 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "(()"
 * 输出：2
 * 解释：最长有效括号子串是 "()"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = ")()())"
 * 输出：4
 * 解释：最长有效括号子串是 "()()"
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = ""
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * s[i] 为 '(' 或 ')'
 *
 *
 *
 *
 */

// @lc code=start
// dp_i 为以i最长有效括号字串长度
func longestValidParentheses(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 0
	if s[0] == '(' && s[1] == ')' {
		dp[1] = 2
	} else {
		dp[1] = 0
	}
	m := maxInt(dp[0], dp[1])
	l := len(s)
	for i := 2; i < l-1; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = dp[i-2] + 2
				m = maxInt(dp[i], m)
			} else {
				// s[i] == ')' && s[i-1] == ')
				// i 对应的括号index 为 i-dp[i-1] -1
				l := i - dp[i-1] - 1
				if l < 0 {
					// 默认就是0 不需要赋值
					// dp[i] = 0
					continue
				}
				if s[l] == '(' {
					dp[i] = dp[i-1] + 2
					if l-1 > 0 {
						dp[i] += dp[l-1]
					}
					m = maxInt(dp[i], m)
				}
			}

		}
	}
	return m

}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

// stack
func longestValidParentheses2(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}
	m := 0

	stack := make([]int, 0, len(s))

	stack = append(stack, -1)

	// 栈底始终是最后一个没有被匹配的右括号的下标
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		} else {

			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				m = maxInt(m, i-stack[len(stack)-1])
			} else {
				stack = append(stack, i)
			}
		}
	}

	return m
}
