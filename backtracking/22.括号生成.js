/*
 * @lc app=leetcode.cn id=22 lang=javascript
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
/**
 * @param {number} n
 * @return {string[]}
 */
var generateParenthesis = function (n) {
  let ans = []
  if (n === 0) {
    return [""]
  }
  backtracking('', n, n*2, 0, 0,ans)
  return ans
}
/**
 *
 * @param {string} str
 * @param {number} n
 * @param {number} count
 * @param {number} leftCount
 * @param {number} rightCount
 * @param {string[]} ans
 */
function backtracking(str, n, count, leftCount, rightCount, ans) {
  if (leftCount < rightCount) return
  if (leftCount > n) return
  if (count == 0) {
    ans.push(str)
    return
  }
  count--
  backtracking(str + '(', n, count, leftCount + 1, rightCount, ans)
  backtracking(str + ')', n, count, leftCount, rightCount + 1, ans)
}
// @lc code=end
