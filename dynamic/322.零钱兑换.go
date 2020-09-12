package leetcode_322


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

// @lc code=start
// 完全背包, 压缩空间
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for index := range dp {
		if index == 0 {
			continue
		}
		dp[index] = amount + 1
	}

	for i, l := 1, len(coins); i <= l; i++ {
		for j := coins[i-1]; j <= amount; j++ {
			dp[j] = minInt(dp[j], dp[j-coins[i-1]]+1)
		}
	}

	if dp[amount] > amount {
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
func coinChange(coins []int, amount int) int {
	// 使用slice 比map性能好, leetcode 216 ms	 -> 28ms
	memo := make([]int, amount)

	return dp(coins, amount, memo)
}
func dp(coins []int, amount int, memo []int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := -1
	if v := memo[amount-1]; v != 0 {
		return v
	}

	for _, coin := range coins {
		count := dp(coins, amount-coin, memo) + 1
		if count == 0 {
			continue
		}
		if count < res {
			res = count
		}

		if res == -1 {
			res = count
		}
	}
	memo[amount-1] = res

	return res
}
