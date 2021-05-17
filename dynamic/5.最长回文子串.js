/*
 * @lc app=leetcode.cn id=5 lang=javascript
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (34.22%)
 * Likes:    3645
 * Dislikes: 0
 * Total Accepted:    584.2K
 * Total Submissions: 1.7M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "cbbd"
 * 输出："bb"
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "a"
 * 输出："a"
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：s = "ac"
 * 输出："a"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * s 仅由数字和英文字母（大写和/或小写）组成
 * 
 * 
 */

// @lc code=start
/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function(s) {
  let l = s.length
  if (l< 2) {
    return s
  }

  let dp = new Array(l)
  for (let i=0; i< l; i++) {
    dp[i] = new Array(l).fill(true)
  }
  let maxLen = 1
  let begin = 0
  // 枚举字符串的长度
  for (let L = 2; L <=l; L++) {
    // i 字符串的左边界
    for (let i = 0; i< l; i++) {
      // j 字符串右边界
      let j = i + L -1
      if (j >= l) break
      if (s[i]!== s[j]) {
        dp[i][j] = false
      } else  {
        if (j - i<3) {
          dp[i][j] = true
        } else {
          // dp[i+1][j-1] 的长度肯定比当前L小, 说明是已经计算过的
          dp[i][j] = dp[i+1][j-1]
        }
      }
      if (dp[i][j]&& L > maxLen) {
        maxLen = L
        begin = i
      }
    }
  }
  return s.substr(begin, maxLen)

};
// @lc code=end

