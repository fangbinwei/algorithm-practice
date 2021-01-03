package leetcode198t2

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 *
 * https://leetcode-cn.com/problems/house-robber/description/
 *
 * algorithms
 * Easy (46.74%)
 * Likes:    1084
 * Dislikes: 0
 * Total Accepted:    190.2K
 * Total Submissions: 407K
 * Testcase Example:  '[1,2,3,1]'
 *
 *
 * 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 *
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[1,2,3,1]
 * 输出：4
 * 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 *
 * 示例 2：
 *
 * 输入：[2,7,9,3,1]
 * 输出：12
 * 解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
 * 偷窃到的最高金额 = 2 + 9 + 1 = 12 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= nums.length <= 100
 * 0 <= nums[i] <= 400
 *
 *
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prevv := 0
	prev := nums[0]

	for i := 2; i < len(nums)+1; i++ {
		// 不偷第i家 dp[i-1]
		// 偷第i家 dp[i-2] + X_i
		current := maxInt(prevv+nums[i-1], prev)
		prevv = prev
		prev = current
	}

	return prev
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

// 未压缩空间
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]

	for i := 2; i < len(nums)+1; i++ {
		// 不偷第i家 dp[i-1]
		// 偷第i家 dp[i-2] + X_i
		dp[i] = maxInt(dp[i-2]+nums[i-1], dp[i-1])
	}

	return dp[len(nums)]
}