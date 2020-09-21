package leetcode_152

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 *
 * https://leetcode-cn.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (40.30%)
 * Likes:    764
 * Dislikes: 0
 * Total Accepted:    94.4K
 * Total Submissions: 234.3K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 *
 *
 * 示例 2:
 *
 * 输入: [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 *
 */
// @lc code=start
// 动态规划, 压缩空间
func maxProduct(nums []int) int {
	dp := []int{1, 1}
	max := math.MinInt64

	for i := 1; i <= len(nums); i++ {
		current := nums[i-1]
		preMin := dp[1]
		preMax := dp[0]
		if current > 0 {
			// current * max
			dp[0] = maxInt(preMax*current, current)
			// current * min
			dp[1] = minInt(preMin*current, current)
		}
		if current <= 0 {
			// current * min
			dp[0] = maxInt(preMin*current, current)
			// current * max
			dp[1] = minInt(preMax*current, current)
		}
		max = maxInt(max, dp[0])
	}
	return max

}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

// 动态规划, 无压缩空间
func maxProduct(nums []int) int {
	// dp[i][0] max
	// dp[i][1] min

	dp := make([][]int, len(nums)+1)
	for index := range dp {
		dp[index] = []int{1, 1}
	}
	max := math.MinInt64

	for i := 1; i <= len(nums); i++ {
		current := nums[i-1]
		if current > 0 {
			// current * max
			dp[i][0] = maxInt(dp[i-1][0]*current, current)
			// current * min
			dp[i][1] = minInt(dp[i-1][1]*current, current)
		}
		if current <= 0 {
			// current * min
			dp[i][0] = maxInt(dp[i-1][1]*current, current)
			// current * max
			dp[i][1] = minInt(dp[i-1][0]*current, current)
		}
		max = maxInt(max, dp[i][0])
	}
	return max

}
