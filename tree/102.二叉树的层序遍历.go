package leetcode_102

/*
 * @lc app=leetcode.cn id=102 lang=golang
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
// 利用递归
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0, 10)
	if root == nil {
		return ans
	}
	ans = append(ans, []int{root.Val})

	leftAns := levelOrder(root.Left)
	rightAns := levelOrder(root.Right)

	var depth int
	if len(leftAns) > len(rightAns) {
		depth = len(leftAns)
	} else {
		depth = len(rightAns)
	}
	//也可以直接for range leftAns, for range rightAns,两个for循环, 代码简洁
	// 下面通过1次for循环
	for i := 0; i < depth; i++ {
		layer := make([]int, 0, 10)
		if i < len(leftAns) {
			layer = append(layer, leftAns[i]...)
		}
		if i < len(rightAns) {

			layer = append(layer, rightAns[i]...)
		}
		ans = append(ans, layer)
	}
	return ans
}

// @lc code=end

// 利用queue
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0, 10)
	queue := make([]*TreeNode, 0, 10)
	if root == nil {
		return ans
	}

	queue = append(queue, root)
	// 使用nil作为标记, nil前面为同一层的节点
	queue = append(queue, nil)

	layer := make([]int, 0, 10)

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if head == nil {
			ans = append(ans, layer)
			if len(queue) == 0 {
				break
			}

			layer = make([]int, 0, 10)
			queue = append(queue, nil)
			continue
		}
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
		layer = append(layer, head.Val)

	}

	return ans
}
