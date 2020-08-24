/*
 * @lc app=leetcode.cn id=22 lang=typescript
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
function generateParenthesis(n: number): string[] {
  let result: string[] = []
  _generate(n * 2 - 1, '(', 1, 0, n, result)
  return result
};
function _generate(n: number, cur: string, leftCount: number, rightCount: number, maxCount: number, result: string[]): void {
  if (n === 0) {
    result.push(cur)
    return
  }
  // 裁剪掉一些冗余的
  if (leftCount < maxCount) {
    _generate(n-1, cur + '(', leftCount + 1, rightCount, maxCount, result)
  }
  if (rightCount < leftCount) {
    _generate(n-1, cur + ')', leftCount, rightCount + 1, maxCount, result)
  }
}
// @lc code=end

