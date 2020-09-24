package leetcode_213

/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 *
 * https://leetcode-cn.com/problems/house-robber-ii/description/
 *
 * algorithms
 * Medium (39.43%)
 * Likes:    380
 * Dislikes: 0
 * Total Accepted:    57.7K
 * Total Submissions: 146.4K
 * Testcase Example:  '[2,3,2]'
 *
 *
 * 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 *
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
 *
 * 示例 1:
 *
 * 输入: [2,3,2]
 * 输出: 3
 * 解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
 *
 *
 * 示例 2:
 *
 * 输入: [1,2,3,1]
 * 输出: 4
 * 解释: 你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 *
 */

// @lc code=start
// 可以直接利用198的结果, 传入nums[1:], nums[:len(nums)-1], 比较结果大小
func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}	
	if l == 1 {
		return nums[0]
	}	
	if l == 2 {
		return maxInt(nums[0], nums[1])
	}

	return maxInt(_rob(nums[1:]), _rob(nums[:l]))

}
func _rob(nums []int) int {
	l := len(nums)
	max := 0
	// 想象前面还有两间没钱的房, 可以避免讨论边界(哨兵)
	pprev := 0
	prev := 0

	var hijack int

	for i := 2; i < l+2; i++ {
		hijack = maxInt(prev, pprev+nums[i-2])
		pprev = prev
		prev = hijack
		max = maxInt(max, hijack)
	}

	return max

}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

// 压缩空间
func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return maxInt(nums[0], nums[1])
	}
	// 把第一家去掉, 直接不偷
	pprev1 := nums[1]
	prev1 := maxInt(nums[1], nums[2])
	//把最后一家去掉, 直接不偷
	pprev2 := nums[0]
	prev2 := maxInt(nums[0], nums[1])
	// 最优解肯定在两个之中

	max1 := maxInt(pprev1, prev1)
	max2 := maxInt(pprev2, prev2)

	for i := 2; i < l-1; i++ {
		hijack1 := maxInt(prev1, pprev1+nums[i+1])
		hijack2 := maxInt(prev2, pprev2+nums[i])
		max2 = maxInt(max2, hijack2)
		max1 = maxInt(max1, hijack1)

		pprev1 = prev1
		prev1 = hijack1
		pprev2 = prev2
		prev2 = hijack2
	}
	return maxInt(max1, max2)
}

// 无压缩空间
func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return maxInt(nums[0], nums[1])
	}
	dp1 := make([]int, l-1, l-1)
	dp2 := make([]int, l-1, l-1)
	// 把第一家去掉, 直接不偷
	dp1[0] = nums[1]
	dp1[1] = maxInt(nums[1], nums[2])
	//把最后一家去掉, 直接不偷
	dp2[0] = nums[0]
	dp2[1] = maxInt(nums[0], nums[1])
	// 最优解肯定在两个之中

	max1 := maxInt(dp1[0], dp1[1])
	max2 := maxInt(dp2[0], dp2[1])

	for i := 2; i < l-1; i++ {
		dp1[i] = maxInt(dp1[i-1], dp1[i-2]+nums[i+1])
		dp2[i] = maxInt(dp2[i-1], dp2[i-2]+nums[i])
		max2 = maxInt(max2, dp2[i])
		max1 = maxInt(max1, dp1[i])
	}
	return maxInt(max1, max2)
}
