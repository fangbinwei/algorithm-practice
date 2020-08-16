/*
 * @lc app=leetcode.cn id=242 lang=typescript
 *
 * [242] 有效的字母异位词
 *
 * https://leetcode-cn.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (60.95%)
 * Likes:    235
 * Dislikes: 0
 * Total Accepted:    126.5K
 * Total Submissions: 207.5K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
 * 
 * 示例 1:
 * 
 * 输入: s = "anagram", t = "nagaram"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: s = "rat", t = "car"
 * 输出: false
 * 
 * 说明:
 * 你可以假设字符串只包含小写字母。
 * 
 * 进阶:
 * 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
 * 
 */

// @lc code=start
function isAnagram(s: string, t: string): boolean {
  if (s.length !== t.length) return false
  return s.split('').sort().join('') === t.split('').sort().join('')

};
// @lc code=end


function isAnagram(s: string, t: string): boolean {
  if (s.length !== t.length) return false
  let m = new Map<string, number>()
  for (let i=0; i < s.length; i++) {
    m.set(s[i], m.has(s[i]) ? m.get(s[i]) as number + 1 : 1 )
  }

  for (let i=0; i < t.length; i++) {
    if (m.has(t[i])) {
      m.set(t[i], m.get(t[i]) as number - 1)
    }
  }

  for (let v of m.values()) {
    if (v) return false
  }
  return true
};