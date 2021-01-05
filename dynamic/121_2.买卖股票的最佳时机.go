package leetcode121t2

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (54.95%)
 * Likes:    1206
 * Dislikes: 0
 * Total Accepted:    289.6K
 * Total Submissions: 527.1K
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
 *
 * 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
 *
 * 注意：你不能在买入股票前卖出股票。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [7,1,5,3,6,4]
 * 输出: 5
 * 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 * ⁠    注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
 *
 *
 * 示例 2:
 *
 * 输入: [7,6,4,3,1]
 * 输出: 0
 * 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 *
 *
 */
// 1. 单调队列
// 2. 利用最大子序和
// 3. 动态规划
// 3.1 前i天的最大利润 = max{前i-1天的最大利润,  i天卖出的最大利润}
//  i天利润 = i天价格 - 前i-1天的最低价

// 3.2 第i天持有/未持有时的收益
// @lc code=start

// 第i天持有/未持有时的收益  压缩空间
func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	// 第一天持有股票, 利润
	dp_00 := -prices[0]
	// 第一天未持有,利润
	dp_01 := 0

	for i := 1; i < len(prices); i++ {
		// 第二天未持有股票, 最大利润
		// max{第一天持有 第二天卖出, 第一天未持有}
		dp_01 = maxInt(dp_00+prices[i], dp_01)
		// 第二天持有股票, 最大利润
		// max{第二天买入, 第一天持有}
		// 只有一次交易 不能用dp[i-1][1] - prices[i] dp[i-1][1]中包含1次成功交易的信息
		dp_00 = maxInt(-prices[i], dp_00)
	}
	return dp_01

}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

// 单调队列
func maxProfit2(prices []int) int {
	profit := 0
	if len(prices) == 0 || len(prices) == 1 {
		return profit
	}
	queue := make([]int, 1, len(prices))
	queue[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > queue[len(queue)-1] {
			queue = append(queue, prices[i])
		} else if prices[i] < queue[0] {
			if len(queue) > 1 {
				profit = maxInt(profit, queue[len(queue)-1]-queue[0])
			}
			queue = queue[:0]
			queue = append(queue, prices[i])
		}
	}
	if len(queue) > 1 {
		profit = maxInt(profit, queue[len(queue)-1]-queue[0])
	}
	return profit
}

// 第i天持有/未持有时的收益  未压缩空间
func maxProfit3(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = []int{0, 0}
	}

	// 第一天持有股票, 利润
	dp[0][0] = -prices[0]
	// 第一天未持有,利润
	dp[0][1] = 0

	for i := 1; i < len(prices); i++ {
		// 第二天持有股票, 最大利润
		// max{第二天买入, 第一天持有}
		// 只有一次交易 不能用dp[i-1][1] - prices[i] dp[i-1][1]中包含1次成功交易的信息
		dp[i][0] = maxInt(-prices[i], dp[i-1][0])
		// 第二天未持有股票, 最大利润
		// max{第一天持有 第二天卖出, 第一天未持有}
		dp[i][1] = maxInt(dp[i-1][0]+prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][1]

}
