package main

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (77.92%)
 * Likes:    738
 * Dislikes: 0
 * Total Accepted:    128.3K
 * Total Submissions: 164.6K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: nums = [1,2,3]
 * 输出:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */

// @lc code=start
// 回溯
// https://leetcode-cn.com/problems/subsets/solution/c-zong-jie-liao-hui-su-wen-ti-lei-xing-dai-ni-gao-/
func subsets(nums []int) [][]int {
	result := [][]int{}
	backtrack(&result, []int{}, nums, 0)
	return result
}
func backtrack(res *[][]int, sub []int, nums []int, start int) {
	*res = append(*res, sub)

	for i := start; i < len(nums); i++ {
		sub = append(sub[:len(sub):len(sub)], nums[i])
		backtrack(res, sub, nums, i + 1)
		sub = sub[:len(sub)-1]
	}
}

// @lc code=end

// 遍历
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	result := make([][]int, 0, 10)
	result = append(result, []int{})

	for i, l := 0, len(nums); i < l; i++ {
		for _, r := range result {
			// input [9,0,3,5,7]
			// 比如[9,0,3] cap:4, [9,0,3, 5], 底层是共用一个数组
			// 当往[9,0,3]append 7的时候, 会替换掉[9,0,3,5]中的5
			tmp := append(r, nums[i])
			result = append(result, append([]int{}, tmp...))
			// 或者
			// result = append(result, append(r[0:len(r):len(r)], nums[i]))
			// 错误写法
			// result = append(result, append(r, nums[i]))
		}
	}

	return result
}

// 递归
func subsets(nums []int) [][]int {

	// 初始化, 一个空slice作为初始值
	result := [][]int{}
	_subsets(nums, &result)

	return result
}
func _subsets(nums []int, result *[][]int) [][]int {
	if len(nums) == 0 {
		*result = append(*result, []int{})
		return [][]int{{}}
	}
	r := make([][]int, 0, 10)

	// 子集为 当前第一个值 + 当前第一个值拼上 剩余值的子集

	// e.g. [2,3]的子集为[],[2], [3],[2,3]
	for _, s := range _subsets(nums[1:], result) {
		// 记录[],[2],[3],[2,3], 因为[2,3]的子集仍为[1,2,3]的子集
		r = append(r, s)

		sub := []int{nums[0]}
		sub = append(sub, s...)
		// 记录[1,2],[1,3],[1,2,3]
		r = append(r, sub)

		*result = append(*result, sub)
	}
	return r
}
