package leetcode_122

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/description/
 *
 * algorithms
 * Easy (63.27%)
 * Likes:    868
 * Dislikes: 0
 * Total Accepted:    222.9K
 * Total Submissions: 352.3K
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
 *
 * 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
 *
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [7,1,5,3,6,4]
 * 输出: 7
 * 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
 * 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
 *
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
 * 示例 3:
 *
 * 输入: [7,6,4,3,1]
 * 输出: 0
 * 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 3 * 10 ^ 4
 * 0 <= prices[i] <= 10 ^ 4
 *
 *
 */

// @lc code=start
func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 || l == 1 {
		return 0
	}

	// 不使用单调栈(减小空间复杂度), 直接一次for循环, 低买高卖
	profit := 0
	current := 0
	for i := 1; i < l; i++ {
		if prices[i] < prices[i-1] {
			if i-current > 1 {
				profit += prices[i-1] - prices[current]
			}
			current = i
		}
	}
	if current < l -1 {
		profit += prices[l-1] - prices[current]
	}
	// 直接current也优化没掉, 相当于把121中涨跌的slice中所有正数加起来
	// input: [7,1,5,3,6,4] 涨幅: [0,-6,4,-2,3,-2] 最大利润就是4+3
	// 但是这种方式计算操作更多, 会相对上面记录current的方式慢一点
	// 贪心的味道
	// for i := 1; i < l; i++ {
	// 	if prices[i] > prices[i-1] {
	// 		profit += prices[i] - prices[i-1]
	// 	}
	// }

	return profit

}

// @lc code=end

func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 || l == 1 {
		return 0
	}
	// 单调栈, 非严格单调递增
	stack := make([]int, 0)

	profit := 0

	for i := 0; i < l; i++ {
		//价格下跌  则前一天就应该卖出, 价格下跌的这天要买入
		if len(stack) > 0 && prices[i] < stack[len(stack)-1] {
			if len(stack) > 1 {
				profit = profit + stack[len(stack)-1] - stack[0]
			}
			stack = nil
			// 买入下跌当天的股票
			stack = append(stack, prices[i])
		}
		stack = append(stack, prices[i])
	}
	if len(stack) > 1 {
		profit = profit + stack[len(stack)-1] - stack[0]
	}
	return profit

}
