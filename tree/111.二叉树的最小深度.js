/*
 * @lc app=leetcode.cn id=111 lang=javascript
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (44.27%)
 * Likes:    354
 * Dislikes: 0
 * Total Accepted:    129.3K
 * Total Submissions: 292K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 * 
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例:
 * 
 * 给定二叉树 [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 返回它的最小深度  2.
 * 
 */

 function TreeNode(val) {
     this.val = val;
     this.left = this.right = null;
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
// dfs
var minDepth = function(root) {
  if (root == null) {
    return 0
  }
  if (!root.left && !root.right) {
    return 1
  }
  if (root.left == null) {
    return minDepth(root.right) +1
  }
  if (root.right == null) {
    return minDepth(root.left) +1
  }

  return  Math.min(minDepth(root.left), minDepth(root.right)) +1

};
// @lc code=end


// 层序遍历
var minDepth2 = function(root) {
  let ans = 0
  if (root == null) {
    return ans
  }
  const queue = [root]

  let flag = true
  while (queue.length > 0 && flag) {
    for (let i=0,l=queue.length; i< l; i++) {
      const head = queue.shift()
      if (!head.left && !head.right) {
        flag = false
        break
      }
      if (head.left!= null) {
        queue.push(head.left)
      }
      if (head.right != null) {
        queue.push(head.right)
      }
    }
    ans++
  }

  return ans

};