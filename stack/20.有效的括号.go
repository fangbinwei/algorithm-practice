package main

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (42.77%)
 * Likes:    1792
 * Dislikes: 0
 * Total Accepted:    379.5K
 * Total Submissions: 887.2K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */

// @lc code=start
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := make([]byte, 0, 100)
	for i:=0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, ')')
		} else if s[i] == '{' {
			stack = append(stack, '}')
		} else if s[i] == '[' {
			stack = append(stack, ']')
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack) -1] == s[i] {
				stack = stack[:len(stack) -1]
			}
		}
	}
	return len(stack) == 0
}
// @lc code=end

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	const (
		a byte = '('
		b byte = ')'
		c byte = '{'
		d byte = '}'
		e byte = '['
		f byte = ']'
	)
	stack := make([]byte, 0, 100)
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		l := len(stack)
		// 优化方向, switch用map代替
		if l > 1 {
			switch stack[l-1] {
			case b:
				if stack[l-2] != a {
					return false
				} else {
					stack = stack[:l-2]
				}
			case d:
				if stack[l-2] != c {
					return false
				} else {
					stack = stack[:l-2]
				}
			case f:
				if stack[l-2] != e {
					return false
				} else {
					stack = stack[:l-2]
				}
			}
		}
	}
	return len(stack) == 0

}