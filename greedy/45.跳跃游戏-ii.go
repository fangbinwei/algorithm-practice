package leetcode45

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 *
 * https://leetcode-cn.com/problems/jump-game-ii/description/
 *
 * algorithms
 * Hard (38.24%)
 * Likes:    789
 * Dislikes: 0
 * Total Accepted:    97.4K
 * Total Submissions: 254.6K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组，你最初位于数组的第一个位置。
 *
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 *
 * 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
 *
 * 示例:
 *
 * 输入: [2,3,1,1,4]
 * 输出: 2
 * 解释: 跳到最后一个位置的最小跳跃数是 2。
 * 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
 *
 *
 * 说明:
 *
 * 假设你总是可以到达数组的最后一个位置。
 *
 */

// @lc code=start
func jump(nums []int) int {
	maxIndex := 0
	count := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		maxIndex = maxInt(maxIndex, i+nums[i])

		// 输入: [2,3,1,1,4]
		// 第一次最远可以到下标2, 但是跳到下标2的过程中发现
		// 其实我们可以最远跳到下标4
		// 那么最优解是: 第一步就是跳到下标0~下标2之间的一步(具体哪一步不重要)
		// 第二步跳到更远的下标4
		if i == end {
			end = maxIndex
			count++
		}
		if maxIndex >= len(nums)-1 {
			return count
		}
	}
	return count
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

// 如果有多个位置通过跳跃都能够到达最后一个位置，那么我们应该如何进行选择呢？
// 直观上来看，我们可以「贪心」地选择距离最后一个位置最远的那个位置
func jump2(nums []int) int {
	maxIndex := 0
	count := 0

	targetIndex := len(nums) - 1
	for targetIndex != 0 {
		for i := 0; i < targetIndex; i++ {
			maxIndex = i + nums[i]
			if maxIndex >= targetIndex {
				count++
				targetIndex = i
				break
			}
		}
	}
	return count
}
