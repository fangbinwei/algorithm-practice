package main

/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/description/
 *
 * algorithms
 * Hard (45.74%)
 * Likes:    516
 * Dislikes: 0
 * Total Accepted:    58.2K
 * Total Submissions: 127.2K
 * Testcase Example:  '[3,3,5,0,0,3,1,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
 *
 * 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
 *
 * 注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 * 示例 1:
 *
 * 输入: [3,3,5,0,0,3,1,4]
 * 输出: 6
 * 解释: 在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
 * 随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
 *
 * 示例 2:
 *
 * 输入: [1,2,3,4,5]
 * 输出: 4
 * 解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4
 * 。
 * 注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
 * 因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
 *
 *
 * 示例 3:
 *
 * 输入: [7,6,4,3,1]
 * 输出: 0
 * 解释: 在这个情况下, 没有交易完成, 所以最大利润为 0。
 *
 */
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/discuss/108870/Most-consistent-ways-of-dealing-with-the-series-of-stock-problems/111002

// @lc code=start
func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	// 买入就算一笔交易

	// 第一天 最多完成1笔交易, 手里有股票
	dp_i11 := -prices[0]
	// 第一天手里没有股票, 收益0
	dp_i10 := 0
	// 第一天 最多完成2笔交易
	dp_i21 := -prices[0]
	// 第一天手里没有股票
	dp_i20 := 0

	// 从第二天开始
	for i := 1; i < l; i++ {

		dp_i20 = maxInt(dp_i20, dp_i21+prices[i])
		dp_i21 = maxInt(dp_i21, dp_i10-prices[i])

		dp_i10 = maxInt(dp_i10, dp_i11+prices[i])
		dp_i11 = maxInt(dp_i11, -prices[i])
		// 看到有网友避免顺序问题的一种方法
		// 下面这种思路, 可以保证max的时候, 使用的都是旧值
		// 	dp20, dp21, dp10, dp11 = max(dp20, dp21+prices[i]), max(dp21, dp10-prices[i]), max(dp10, dp11+prices[i]), max(dp11, -prices[i])

		// for k := 1; k <= 2; k++ {
		// 	// max{前一天就是卖出 今天不操作, 前一天是买入, 今天卖出, k+1}
		// 	dp[k][0] = maxInt(dp[k][0], dp[k][1]+prices[i])
		// 	// max{前一天就是买入 今天不操作, 前一天是卖出, 今天买入}
		// 	dp[k][1] = maxInt(dp[k][1], dp[k-1][0]-prices[i])
		// }
	}
	return dp_i20
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

// 压缩部分空间
func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	dp := make([][]int, 3)
	for k := 0; k <= 2; k++ {
		// []int{sell, buy}
		dp[k] = []int{0, 0}
	}
	// 买入就算一笔交易

	// 第一天
	dp[0][0] = 0
	dp[1][0] = 0

	// 第一天 最多完成0笔交易买入股票
	dp[0][1] = 0
	// 第一天 最多完成1笔交易
	dp[1][1] = -prices[0]
	// 第一天 最多完成2笔交易
	dp[2][1] = -prices[0]

	// 从第二天开始
	for i := 1; i < l; i++ {
		// 可以把这边循环展开, 继续压缩空间
		for k := 1; k <= 2; k++ {
			// max{前一天就是卖出 今天不操作, 前一天是买入, 今天卖出, k+1}
			dp[k][0] = maxInt(dp[k][0], dp[k][1]+prices[i])
			// max{前一天就是买入 今天不操作, 前一天是卖出, 今天买入}
			dp[k][1] = maxInt(dp[k][1], dp[k-1][0]-prices[i])

		}
	}
	return dp[2][0]
}

// 不压缩空间
func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	dp := make([][][]int, l)
	for index := range dp {
		kdp := make([][]int, 3)
		for k := 0; k <= 2; k++ {
			// []int{sell, buy}
			kdp[k] = []int{0, 0}
		}
		dp[index] = kdp
	}
	// 买入就算一笔交易
	// Note that the maximum number of allowable transactions remains the same,
	// due to the fact that a transaction consists of two actions coming as a pair -- buy and sell.
	//  Only action buy will change the maximum number of transactions allowed
	// (well, there is actually an alternative interpretation, see my comment below).

	// 第一天
	dp[0][0][0] = 0
	dp[0][1][0] = 0

	// 第一天 最多完成0笔交易买入股票
	dp[0][0][1] = 0
	// 第一天 最多完成1笔交易
	dp[0][1][1] = -prices[0]
	// 第一天 最多完成2笔交易
	dp[0][2][1] = -prices[0]

	// 从第二天开始
	for i := 1; i < l; i++ {
		for k := 1; k < 3; k++ {
			// 0代表手里没有股票, 可能是前一天卖出了, 或者今天卖出了
			// max{前一天就是卖出 今天不操作, 前一天是买入, 今天卖出, k+1}
			dp[i][k][0] = maxInt(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			// 1代表手里有股票, 可能是前一天买入了, 或者今天买入
			// max{前一天就是买入 今天不操作, 前一天是卖出, 今天买入}
			dp[i][k][1] = maxInt(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])

		}
	}
	return dp[l-1][2][0]
}
