/*
 * @lc app=leetcode.cn id=70 lang=typescript
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (50.22%)
 * Likes:    1181
 * Dislikes: 0
 * Total Accepted:    259.2K
 * Total Submissions: 516K
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *
 * 注意：给定 n 是一个正整数。
 *
 * 示例 1：
 *
 * 输入： 2
 * 输出： 2
 * 解释： 有两种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶
 * 2.  2 阶
 *
 * 示例 2：
 *
 * 输入： 3
 * 输出： 3
 * 解释： 有三种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶 + 1 阶
 * 2.  1 阶 + 2 阶
 * 3.  2 阶 + 1 阶
 *
 *
 */

// @lc code=start
let memo = new Map<number, number>()
function climbStairs(n: number): number {
  if (n === 1) return 1
  if (n === 2) return 2
  let sum =
    (memo.has(n - 1) ? memo.get(n - 1) : climbStairs(n - 1)) +
    (memo.has(n - 2) ? memo.get(n - 2) : climbStairs(n - 2));
  if (!memo.has(n)) {
    memo.set(n, sum)
  }
  return sum
}
// @lc code=end

//循环
function climbStairs(n: number): number {
  if (n === 1) return 1
  if (n === 2) return 2
  let sum = 0
  let pre = 1
  let next = 2
  for (let i = 2; i < n; i++) {
    sum = pre + next
    pre = next
    next = sum
  }
  return sum
}
// 尾调用优化
function climbStairs(n: number): number {
  if (n === 1) return 1
  if (n === 2) return 2
  return fib(n, 0, 1)
  function fib(c1: number, c2: number, c3: number): number {
    if (c1 === 0) {
      return c2
    }
    return fib(c1 - 1, c3, c2 + c3)
  }
}
