/*
 * @lc app=leetcode.cn id=15 lang=typescript
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (28.93%)
 * Likes:    2462
 * Dislikes: 0
 * Total Accepted:    296.2K
 * Total Submissions: 1M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？请你找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例：
 *
 * 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */

// @lc code=start
function threeSum(nums: number[]): number[][] {
  let result = []
  nums.sort((a, b) => a - b)
  if (nums.length < 3) return []
  for (let i = 0; i < nums.length - 2; i++) {
    if (i > 0 && nums[i] === nums[i - 1]) continue
    if (nums[i] > 0) break
    let target = -nums[i]
    let left = i + 1
    let right = nums.length - 1
    while (left < right) {
      if (nums[left] + nums[right] === target) {
        result.push([nums[i], nums[left], nums[right]])
        left++
        right--
        while (nums[left] === nums[left - 1]) {
          left++
        }
        while (nums[right] === nums[right + 1]) {
          right--
        }
        continue
      }
      if (nums[left] + nums[right] < target) {
        left++
      } else {
        right--
      }
    }
  }
  return result
}
// @lc code=end
