package leetcode55

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 *
 * https://leetcode-cn.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (41.48%)
 * Likes:    982
 * Dislikes: 0
 * Total Accepted:    182.4K
 * Total Submissions: 439.8K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组，你最初位于数组的第一个位置。
 *
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 *
 * 判断你是否能够到达最后一个位置。
 *
 * 示例 1:
 *
 * 输入: [2,3,1,1,4]
 * 输出: true
 * 解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
 *
 *
 * 示例 2:
 *
 * 输入: [3,2,1,0,4]
 * 输出: false
 * 解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
 *
 *
 */

// @lc code=start
func canJump(nums []int) bool {
	// 只有一个位置, 该位置就是终点
	if len(nums) == 1 {
		return true
	}

	curIndex := 0
	// 当前能跳的步数
	curNum := nums[curIndex]

	// 满足条件说明能继续往前跳, curNum+curIndex 为最远能到达的位置
	for curIndex < curNum+curIndex {
		// 往前跳一步
		curIndex++
		// 跳到了终点
		if curIndex == len(nums)-1 {
			return true
		}
		// 贪婪的思想
		// 跳了一步之后, 当前能跳的步数为max{前一格能跳步数-1, 当前格能跳的步数}
		curNum = maxInt(nums[curIndex], curNum-1)
	}
	// 没跳到终点
	return false
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
