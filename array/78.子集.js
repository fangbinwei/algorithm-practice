/*
 * @lc app=leetcode.cn id=78 lang=javascript
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (77.92%)
 * Likes:    738
 * Dislikes: 0
 * Total Accepted:    128.3K
 * Total Submissions: 164.6K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: nums = [1,2,3]
 * 输出:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
// backtracking
var subsets = function (nums) {
  let ans = []
  let set = []
  backtracking(0)
  return ans

  function backtracking(index) {
    if (index === nums.length) {
      ans.push([...set])
      return
    }
    set.push(nums[index])
    backtracking(index+1)
    set.pop()
    backtracking(index+1)
  }
}
// @lc code=end

var subsets3 = function (nums) {
  let ans = []
  backtracking(0, [])
  return ans
  function backtracking(index, collection) {
    let tmp = [...collection]
     ans.push(tmp)

    for (let i=index;i<nums.length;i++) {
      tmp.push(nums[i])
      ++index
      backtracking(index, tmp)
      tmp.pop()
    }
  }
}
var subsets2 = function (nums) {
  let ans = [[]]
  for (let i=0;i<nums.length;i++) {
    for(let j=0,l=ans.length; j<l;j++) {
      let tmp = ans[j].slice(0)
      tmp.push(nums[i])
      ans.push(tmp)
    }
  }
  return ans
}