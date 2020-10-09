package leetcode_189

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 *
 * https://leetcode-cn.com/problems/rotate-array/description/
 *
 * algorithms
 * Medium (43.22%)
 * Likes:    710
 * Dislikes: 0
 * Total Accepted:    170.2K
 * Total Submissions: 393.8K
 * Testcase Example:  '[1,2,3,4,5,6,7]\n3'
 *
 * 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
 *
 * 示例 1:
 *
 * 输入: [1,2,3,4,5,6,7] 和 k = 3
 * 输出: [5,6,7,1,2,3,4]
 * 解释:
 * 向右旋转 1 步: [7,1,2,3,4,5,6]
 * 向右旋转 2 步: [6,7,1,2,3,4,5]
 * 向右旋转 3 步: [5,6,7,1,2,3,4]
 *
 *
 * 示例 2:
 *
 * 输入: [-1,-100,3,99] 和 k = 2
 * 输出: [3,99,-1,-100]
 * 解释:
 * 向右旋转 1 步: [99,-1,-100,3]
 * 向右旋转 2 步: [3,99,-1,-100]
 *
 * 说明:
 *
 *
 * 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
 * 要求使用空间复杂度为 O(1) 的 原地 算法。
 *
 *
 */

// @lc code=start
// 反转法
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	// [1,2,3,4,5,6,7] -> [7,6,5,4,3,2,1]
	// 反转[0:k], 反转[k:]
	// [5,6,7] [1,2,3,4]
	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k:])
}
func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// @lc code=end

func rotate(nums []int, k int) {
	l := len(nums)
	for i := 0; i < k; i++ {
		last := nums[l-1]
		copy(nums[1:], nums[:l-1])
		nums[0] = last
	}

}
