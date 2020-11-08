package leetcode47

import "sort"

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (62.16%)
 * Likes:    518
 * Dislikes: 0
 * Total Accepted:    119K
 * Total Submissions: 191.5K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,2]
 * 输出：
 * [[1,1,2],
 * ⁠[1,2,1],
 * ⁠[2,1,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10
 *
 *
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0, 10)
	sort.Ints(nums)

	visited := make([]bool, len(nums))
	var backtracking func(int)

	result := make([]int, 0)

	backtracking = func(currentFillIndex int) {
		if currentFillIndex == len(nums) {
			ans = append(ans, append([]int{}, result...))
		}

		// 从nums中选择, 填充`currentFillIndex`位置
		for index, item := range nums {
			// num中已经使用过的跳过
			// 如果前一个数没使用过, 当前数和前一个数相同, 也跳过
			// 因为当前的数填充 和前一个数填充的效果相同, 重复
			if visited[index] || index != 0 && nums[index] == nums[index-1] && !visited[index-1] {
				continue
			}

			result = append(result, item)
			visited[index] = true
			backtracking(currentFillIndex + 1)
			visited[index] = false
			result = result[:len(result)-1]
		}

	}

	// 先填充0位置
	backtracking(0)

	return ans
}

// @lc code=end
