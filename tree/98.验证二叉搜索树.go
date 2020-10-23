package leetcode_98

import "math"

/*
 * @lc app=leetcode.cn id=98 lang=golang
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
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 
func isValidBST(root *TreeNode) bool {
	return _isValidBST(root, math.MinInt64, math.MaxInt64)
}

func _isValidBST(root *TreeNode, lower, upper int) bool {
	// 子树的所有值 要求在范围(lower, upper)
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
		if root.Left.Val <= lower {
			return false
		}
	} 
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
		if root.Right.Val >= upper {
			return false
		}
	}


	// lower< 左子树 的值要求<root.Val
	if left := _isValidBST(root.Left, lower, root.Val); !left {
		return false
	}
	// root.Val<右子树的值 < upper
	if right := _isValidBST(root.Right, root.Val, upper); !right {
		return false
	}

	return true
}

// @lc code=end

// 中序遍历得到的数, 肯定是从小到大的
func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := []*TreeNode{root}
	min := math.MinInt64

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top == nil {
			cur := stack[len(stack)-1]

			// 不满足从小到大的增长, 不是二叉搜索树
			if cur.Val <= min {
				return false
			}

			min = cur.Val
			stack = stack[:len(stack)-1]
		} else {
			// 中序遍历是左 -》 根 -》 右
			// stack中就右 -》 根 -》 左的顺序
			if top.Right != nil {
				stack = append(stack, top.Right)
			}
			stack = append(stack, top)
			stack = append(stack, nil)
			if top.Left != nil {
				stack = append(stack, top.Left)
			}
		}
	}
	return true
}