package leetcode_53

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (52.37%)
 * Likes:    2438
 * Dislikes: 0
 * Total Accepted:    328.6K
 * Total Submissions: 627.5K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 *
 * 示例:
 *
 * 输入: [-2,1,-3,4,-1,2,1,-5,4]
 * 输出: 6
 * 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 *
 *
 * 进阶:
 *
 * 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 *
 */

// 算法导论分治经典题
// 动态规划, dp[i] = max(dp[i-1] + C_i, C_i)
// i = 1...len(nums)
// @lc code=start
// 动态规划, 压缩空间, 空间复杂度O(1)
func maxSubArray(nums []int) int {
	pre:= nums[0]
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		pre = maxInt(pre+nums[i], nums[i])
		max = maxInt(pre, max)
	}
	return max

}
func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

// 动态规划, 无压缩空间, 空间复杂度O(n)
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums)+1, len(nums)+1)

	for i := 1; i <= len(nums); i++ {
		dp[i] = maxInt(dp[i-1]+nums[i-1], nums[i-1])
	}
	max := dp[1]
	for i:=2; i <=len(nums); i++ {
		max = maxInt(max, dp[i])
	}
	return max

}