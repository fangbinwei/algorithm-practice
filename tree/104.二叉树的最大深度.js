/*
 * @lc app=leetcode.cn id=104 lang=javascript
 *
 * [104] 二叉树的最大深度
 *
 * https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (74.91%)
 * Likes:    682
 * Dislikes: 0
 * Total Accepted:    264K
 * Total Submissions: 352.4K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最大深度。
 *
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例：
 * 给定二叉树 [3,9,20,null,null,15,7]，
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最大深度 3 。
 *
 */

function TreeNode(val) {
  this.val = val
  this.left = this.right = null
}
// @lc code=start
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
var maxDepth = function (root) {
  let ans = 0
  if (root == null) {
    return ans
  }

  let queue = [root]

  while (queue.length > 0) {
    for (let i = 0, l = queue.length; i < l; i++) {
      let head = queue.shift()
      if (head.left != null) {
        queue.push(head.left)
      }
      if (head.right != null) {
        queue.push(head.right)
      }
    }
    ans++
  }
  return ans
}
// @lc code=end
