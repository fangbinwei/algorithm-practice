/*
 * @lc app=leetcode.cn id=283 lang=typescript
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
/**
 Do not return anything, modify nums in-place instead.
 */
// 双指针
function moveZeroes(nums: number[]): void {
  if (nums.length < 2) return
  let j = 0
  for (let c = 0; c < nums.length; c++) {
    if (nums[c] !== 0) {
      nums[j++] = nums[c]
    }
  }
  while(j< nums.length) {
    nums[j++] = 0
  }
};
// @lc code=end


// 快慢指针
// function moveZeroes(nums: number[]): void {
//   if (nums.length < 2) return
//   let lastNoZeroIndex = 0
//   for (let c = 0; c < nums.length; c++) {
//     if (nums[c]!==0) {
//       const bak = nums[lastNoZeroIndex]
//       nums[lastNoZeroIndex++] = nums[c]
//       nums[c] = bak
//     }
//   }
// };


// 快慢指针
// function moveZeroes(nums: number[]): void {
  // if (nums.length < 2) return
  // let placeholder = 0
  // let current = 1 
  // while (current < nums.length) {
  //   if (nums[placeholder] === 0) {
  //     if (nums[current] === 0) {
  //       current++
  //       continue
  //     }
  //     nums[placeholder++] = nums[current]
  //     nums[current++] = 0
  //     continue
  //   }
  //   placeholder++
  //   current++
  // }
// };