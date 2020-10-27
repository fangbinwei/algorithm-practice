/*
 * @lc app=leetcode.cn id=102 lang=javascript
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (63.33%)
 * Likes:    624
 * Dislikes: 0
 * Total Accepted:    191K
 * Total Submissions: 301.5K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
 * 
 * 
 * 
 * 示例：
 * 二叉树：[3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回其层次遍历结果：
 * 
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
 * 
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
 * @return {number[][]}
 */
// DFS
var levelOrder = function(root) {
  let ans = []
  dfs(root, 0, ans)

  return ans

};

/**
 * @param {TreeNode} root
 * @param {number} index
 * @param {number[][]} ans
 */
function dfs(root, index, ans) {
  if (root == null) return

  if (ans[index] == null) ans[index] = []
  ans[index].push(root.val)
  dfs(root.left, index+1, ans)
  dfs(root.right, index+1, ans)
}
// @lc code=end


// BFS
var levelOrder1 = function(root) {
  let ans = []
  if (root == null) {
    return ans
  }

  const queue = [root]
  while (queue.length) {
    let layer = []
    for (let i=0,l=queue.length; i< l; i++) {
      const head = queue.shift()
      if (head.left) {
        queue.push(head.left)
      }
      if (head.right) {
        queue.push(head.right)
      }
      layer.push(head.val)
    }
    ans.push(layer)
  }
  return ans

};