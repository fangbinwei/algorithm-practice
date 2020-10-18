/*
 * @lc app=leetcode.cn id=429 lang=javascript
 *
 * [429] N叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (66.23%)
 * Likes:    106
 * Dislikes: 0
 * Total Accepted:    27.1K
 * Total Submissions: 40.9K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
 * 
 * 例如，给定一个 3叉树 :
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 返回其层序遍历:
 * 
 * [
 * ⁠    [1],
 * ⁠    [3,2,4],
 * ⁠    [5,6]
 * ]
 * 
 * 
 * 
 * 
 * 说明:
 * 
 * 
 * 树的深度不会超过 1000。
 * 树的节点总数不会超过 5000。
 * 
 */

// @lc code=start
/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node} root
 * @return {number[][]}
 */
var levelOrder = function(root) {
  let ans = []
  if (root == null) {
    return ans
  }
  let stack = [root]

  while (stack.length>0) {
    ans.push([])
    let l = stack.length
    for (let i=0; i< l; i++) {
      ans[ans.length-1].push(stack[i].val)
      if (stack[i].children) {
        for (let j=0; j< stack[i].children.length;j++) {
          stack.push(stack[i].children[j])
        }
      }
    }
    while(l>0) {
      stack.shift()
      l--
    }
  }
  return ans
    
};
// @lc code=end

