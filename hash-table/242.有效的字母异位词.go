package main

/*
 * @lc app=leetcode.cn id=242 lang=golang
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
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	//var charCt = make([]int, 26)
	m := make(map[byte]int)

	i, l := 0, len(s)
	for ; i < l; i++ {
		// cover rune to byte
		m[s[i]-'a'] += 1
	}
	i = 0

	for ; i < l; i++ {
		// cover rune to byte
		m[t[i]-'a'] -= 1
		if m[t[i]-'a'] < 0 {
			return false
		}
	}

	return true

}

// @lc code=end