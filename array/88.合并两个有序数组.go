package leetcode_88

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 *
 * https://leetcode-cn.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (48.74%)
 * Likes:    648
 * Dislikes: 0
 * Total Accepted:    216.8K
 * Total Submissions: 444.8K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
 *
 *
 *
 * 说明:
 *
 *
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
 * 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 *
 *
 *
 *
 * 示例:
 *
 * 输入:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * 输出: [1,2,2,3,5,6]
 *
 */

// @lc code=start
// 双指针从后
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1

	p := len(nums1) - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] <= nums2[p2] {
			nums1[p] = nums2[p2]
			p2--
		} else {
			nums1[p] = nums1[p1]
			p1--
		}
		p--
	}
	// 只需要复制num2
	// 只有2中情况
	// 1. num2有剩余
	// 2. num2没有剩余
	copy(nums1[:p+1], nums2[:p2+1])
}

// @lc code=end

// 双指针从前往后
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := 0
	nums1_copy := make([]int, m)
	copy(nums1_copy, nums1)

	p2 := 0

	p := 0

	for p1 < m && p2 < n {
		if nums1_copy[p1] <= nums2[p2] {
			nums1[p] = nums1_copy[p1]
			p1++
		} else {
			nums1[p] = nums2[p2]
			p2++
		}
		p++
	}
	if p1 < m {
		copy(nums1[p:], nums1_copy[p1:])
	} else {
		copy(nums1[p:], nums2[p2:])
	}
}

// 暴力复制法
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	count := m
	for j, l := 0, n; j < l; {
		if i < count && nums2[j] <= nums1[i] {
			copy(nums1[i+1:], nums1[i:])
			nums1[i] = nums2[j]
			j++

			count++
		} else if i >= count {
			// 这些比nums1中所有的数都大
			nums1[i] = nums2[j]
			j++
		}
		i++
	}
}
