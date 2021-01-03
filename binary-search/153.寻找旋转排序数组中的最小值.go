package leetcode153

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 *
 * https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (52.09%)
 * Likes:    320
 * Dislikes: 0
 * Total Accepted:    95.6K
 * Total Submissions: 183.6K
 * Testcase Example:  '[3,4,5,1,2]'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。例如，数组 [0,1,2,4,5,6,7]  可能变为 [4,5,6,7,0,1,2] 。
 *
 * 请找出其中最小的元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [3,4,5,1,2]
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,5,6,7,0,1,2]
 * 输出：0
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -5000
 * nums 中的所有整数都是 唯一 的
 * nums 原来是一个升序排序的数组，但在预先未知的某个点上进行了旋转
 *
 *
 */

// @lc code=start
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		mid := left + (right-left)/2
		// [left, mid] 为升序
		if nums[left] <= nums[mid] {
			left = mid + 1
			// [mid+1, right] 为升序
		} else {
			right = mid
		}
	}
	return nums[left]
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

func findMin2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	min := 10000
	left := 0
	right := len(nums) - 1

	if nums[0] < nums[right] {
		return nums[0]
	}

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if mid-1 >= 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		// [left, mid] 为升序
		if nums[left] < nums[mid] {
			min = minInt(min, nums[left])
			left = mid + 1
			// [mid+1, right] 为升序
		} else {
			min = minInt(min, nums[mid+1])
			right = mid
		}
	}
	// left == right
	min = minInt(min, nums[left])
	return min
}
