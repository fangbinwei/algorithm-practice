/*
 * @lc app=leetcode.cn id=20 lang=typescript
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
// 使用栈
function isValid(s: string): boolean {
  if (s === '') return true
  if (s.length % 2 === 1) return false
  const m: Record<string, string> = {
    ')': '(',
    '}': '{',
    ']': '[',
  }
  const right = [')', '}', ']']
  const stack: string[] = []
  for (let i = 0; i < s.length; i++) {
    if (right.includes(s[i])) {
      // 成对的消除
      if (stack.length > 0 && stack[stack.length - 1] === m[s[i]]) {
        stack.pop()
      } else {
        return false
      }
    } else {
      stack.push(s[i])
    }
    // 长度为奇数的已经排除, stack.length > s.length/2 肯定不能配对
    if (stack.length > s.length / 2) {
      return false
    }
  }
  return stack.length === 0
}
// @lc code=end
