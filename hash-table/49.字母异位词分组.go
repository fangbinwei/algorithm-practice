package main

/*
 * @lc app=leetcode.cn id=49 lang=golang
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
import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0, 200)
	m := make(map[string][]string)
	for _, s := range strs {
		key := convert(s)
		if len(m[key]) == 0 {
			m[key] = make([]string, 0, 50)
		}
		m[key] = append(m[key], s)
	}
	for _, value := range m {
		result = append(result, value)
	}
	return result
}

func convert(o string) string {
	s := strings.Split(o, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// @lc code=end
