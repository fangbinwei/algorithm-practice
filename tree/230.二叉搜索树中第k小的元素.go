package leetcode230

/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
 *
 * https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/description/
 *
 * algorithms
 * Medium (73.15%)
 * Likes:    370
 * Dislikes: 0
 * Total Accepted:    98.7K
 * Total Submissions: 135K
 * Testcase Example:  '[3,1,4,null,2]\n1'
 *
 * 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,1,4,null,2], k = 1
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [5,3,6,2,4,null,null,1], k = 3
 * 输出：3
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数为 n 。
 * 1
 * 0
 *
 *
 *
 *
 * 进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？
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
//中序遍历 左-》 根 -》 右
func kthSmallest(root *TreeNode, k int) int {
	ans := inorder(root)
	return ans[k-1]
}
func inorder (root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	ans = append(ans, inorder(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, inorder(root.Right)...)
	return ans
}

// @lc code=end

//中序遍历 左-》 根 -》 右
func kthSmallest2(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	ans := make([]int, 0)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top == nil {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, top.Val)
		} else {
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
	return ans[k-1]
}