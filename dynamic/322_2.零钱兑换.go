package leetcode322t2

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (41.28%)
 * Likes:    804
 * Dislikes: 0
 * Total Accepted:    132K
 * Total Submissions: 319.9K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给定不同面额的硬币 coins 和一个总金额
 * amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
 *
 *
 *
 * 示例 1:
 *
 * 输入: coins = [1, 2, 5], amount = 11
 * 输出: 3
 * 解释: 11 = 5 + 5 + 1
 *
 * 示例 2:
 *
 * 输入: coins = [2], amount = 3
 * 输出: -1
 *
 *
 *
 * 说明:
 * 你可以认为每种硬币的数量是无限的。
 *
 */

// Count(i, a) i种硬币, a总金额
// Count(i, a) = min{Count(i-1, a), Count(i, a-coins_i)+1}
// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		// 初始化+∞ 
		dp[i] = amount +1
	}

	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				dp[j] = minInt(dp[j], dp[j-coins[i-1]]+1)
			}
		}
	}
	if dp[amount] == amount +1 {
		return -1
	}
	return dp[amount]

}
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
