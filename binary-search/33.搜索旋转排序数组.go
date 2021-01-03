package leetcode33

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (40.03%)
 * Likes:    1122
 * Dislikes: 0
 * Total Accepted:    205.1K
 * Total Submissions: 512.3K
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为 [4,5,6,7,0,1,2]
 * ）。
 *
 * 请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,5,6,7,0,1,2], target = 0
 * 输出：4
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,5,6,7,0,1,2], target = 3
 * 输出：-1
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1], target = 0
 * 输出：-1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10^4
 * nums 中的每个值都 独一无二
 * nums 肯定会在某个点上旋转
 * -10^4
 *
 *
 */

// @lc code=start
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2

		// [left, mid] 为升序
		if nums[left] < nums[mid] {
			if nums[left] == target {
				return left
			}
			if nums[mid] == target {
				return mid
			}
			// target 在[left, mid]
			if nums[left] < target && target < nums[mid] {
				right = mid
			} else {
				// target 不在[left, mid]
				left = mid + 1
			}

			// [mid+1, right] 为升序
		} else {
			if nums[mid+1] == target {
				return mid + 1
			}
			if nums[right] == target {
				return right
			}
			// [mid+1, right]
			if nums[mid+1] < target && target < nums[right] {
				left = mid + 1

			} else {
				right = mid
			}
		}
	}
	// left == right
	if nums[left] == target {
		return left
	}
	return -1

}

// @lc code=end
