/*
 * @lc app=leetcode.cn id=49 lang=typescript
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (63.06%)
 * Likes:    430
 * Dislikes: 0
 * Total Accepted:    96.8K
 * Total Submissions: 153.6K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 * 
 * 示例:
 * 
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
 * 输出:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 * 
 * 说明：
 * 
 * 
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 * 
 * 
 */

// @lc code=start
function groupAnagrams(strs: string[]): string[][] {
  let ans: string[][] = []

  let m = new Map<string, number>()

  for (let i=0; i < strs.length; i++) {
    let normalize = strs[i].split('').sort().join('')
    let index = m.get(normalize)
    if (index === undefined) {
      m.set(normalize, ans.length)
      ans.push([strs[i]])
    } else {
      ans[index].push(strs[i])
    }
  }

  return ans
};

// @lc code=end

