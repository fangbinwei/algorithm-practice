package leetcode_42

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (52.78%)
 * Likes:    1745
 * Dislikes: 0
 * Total Accepted:    156.2K
 * Total Submissions: 296K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 *
 * 示例:
 *
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 *
 */

// @lc code=start
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	ans := 0
	// 单调栈
	stack := make([]int, 0, len(height))

	for i, v := range height {
		for len(stack) > 0 && v > height[stack[len(stack)-1]] {
			// 可以形成凹槽
			if len(stack) > 1 {
				// 长 * 宽
				ans += (minInt(v, height[stack[len(stack)-2]]) - height[stack[len(stack)-1]]) * (i - stack[len(stack)-2] - 1)
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return ans
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
