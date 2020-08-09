/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode-cn.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (61.89%)
 * Likes:    686
 * Dislikes: 0
 * Total Accepted:    187.7K
 * Total Submissions: 303.2K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * 示例:
 *
 * 输入: [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 * 说明:
 *
 *
 * 必须在原数组上操作，不能拷贝额外的数组。
 * 尽量减少操作次数。
 *
 *
 */

// @lc code=start
func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}
	j := 0
	for c := 0; c < len(nums); c++ {
		if nums[c] != 0 {
			nums[j] = nums[c]
			j++
		}
	}
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}

// @lc code=end

