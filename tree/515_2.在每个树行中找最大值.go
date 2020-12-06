package leetcode515t2

/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
 *
 * https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/description/
 *
 * algorithms
 * Medium (61.79%)
 * Likes:    90
 * Dislikes: 0
 * Total Accepted:    17.4K
 * Total Submissions: 28.2K
 * Testcase Example:  '[1,3,2,5,3,null,9]'
 *
 * 您需要在二叉树的每一行中找到最大的值。
 *
 * 示例：
 *
 *
 * 输入:
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       3   2
 * ⁠      / \   \
 * ⁠     5   3   9
 *
 * 输出: [1, 3, 9]
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

func largestValues(root *TreeNode) []int {
	ans := make([]int, 0, 10)
	if root == nil {
		return ans
	}

	queue := make([]*TreeNode, 0, 10)
	queue = append(queue, root)

	for len(queue) > 0 {
		l := len(queue)

		max := queue[0].Val
		for i := 0; i < l; i++ {
			head := queue[0]
			queue = queue[1:]

			if head.Val > max {
				max = head.Val
			}

			if head.Left != nil {
				queue = append(queue, head.Left)
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}
		ans = append(ans, max)

	}
	return ans

}

// @lc code=end
