package leetcode209

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode-cn.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (45.16%)
 * Likes:    571
 * Dislikes: 0
 * Total Accepted:    118.6K
 * Total Submissions: 262.6K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 target 。
 *
 * 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr]
 * ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：target = 7, nums = [2,3,1,2,4,3]
 * 输出：2
 * 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
 *
 *
 * 示例 2：
 *
 *
 * 输入：target = 4, nums = [1,4,4]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * 1
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
 *
 *
 */

// @lc code=start
//滑动窗口
func minSubArrayLen(target int, nums []int) int {
	size := len(nums)
	queue := make([]int, size)
	sum := 0
	minLen := size + 1

	for i := 0; i < size; i++ {
		queue = append(queue, nums[i])
		// 不断进队列, >=target时, 再出队列(滑动窗口)
		sum = sum + nums[i]
		for sum >= target && len(queue) > 0 {
			minLen = minInt(minLen, len(queue))
			sum = sum - queue[0]
			queue = queue[1:]
		}
	}
	if minLen == size+1 {
		return 0
	}

	return minLen

}

func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// @lc code=end
