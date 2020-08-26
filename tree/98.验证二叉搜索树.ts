/*
 * @lc app=leetcode.cn id=98 lang=typescript
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (32.19%)
 * Likes:    733
 * Dislikes: 0
 * Total Accepted:    160.9K
 * Total Submissions: 500K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 * 
 * 假设一个二叉搜索树具有如下特征：
 * 
 * 
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 * 
 * 
 * 示例 1:
 * 
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 * 根节点的值为 5 ，但是其右子节点值为 4 。
 * 
 * 
 */

 class TreeNode {
     val: number
     left: TreeNode | null
     right: TreeNode | null
     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
         this.val = (val===undefined ? 0 : val)
         this.left = (left===undefined ? null : left)
         this.right = (right===undefined ? null : right)
     }
 }
// @lc code=start
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

 // 中序遍历
function isValidBST(root: TreeNode | null): boolean {
  if (root == null) return true
  let stack: (TreeNode|null)[] = [root]
  let res: number[] = []

  while(stack.length) {
    let cur: TreeNode = stack.pop()
    if (cur !== null) {
      cur.right && stack.push(cur.right)
      stack.push(cur)
      stack.push(null)
      cur.left && stack.push(cur.left)
    } else {
      let value = stack.pop().val
      if (res.length) {
        if (value <= res[res.length - 1]) return false
      }
      res.push(value)
    }
  }
  return true

};
// @lc code=end


// https://leetcode-cn.com/problems/validate-binary-search-tree/solution/yi-zhang-tu-rang-ni-ming-bai-shang-xia-jie-zui-da-/
function isValidBST(root: TreeNode | null): boolean {
  return helper(root, Number.MAX_SAFE_INTEGER, Number.MIN_SAFE_INTEGER)

};
function helper(root: TreeNode | null, max: number, min: number ): boolean {
  if (root == null) return true
  if (root.val >= max ) return false
  if (root.val <= min) return false

  if (!helper(root.left, root.val, min)) return false
  if (!helper(root.right, max, root.val)) return false
  return true

}