package main

/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/
 *
 * algorithms
 * Easy (73.83%)
 * Likes:    93
 * Dislikes: 0
 * Total Accepted:    39.7K
 * Total Submissions: 53.8K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的前序遍历。
 *
 * 例如，给定一个 3叉树 :
 *
 *
 *
 *
 *
 *
 *
 * 返回其前序遍历: [1,3,5,6,2,4]。
 *
 *
 *
 * 说明: 递归法很简单，你可以使用迭代法完成此题吗?
 */

type Node struct {
	Val      int
	Children []*Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	res := make([]int, 0, 10)
	if root == nil {
		return res
	}
	stack := make([]*Node, 0, 10)
	stack = append(stack, root)
	for len(stack) > 0 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if t != nil {
			for i := len(t.Children) - 1; i >= 0; i-- {
				if c := t.Children[i]; c != nil {
					stack = append(stack, c)
				}
			}
			stack = append(stack, t)
			stack = append(stack, nil)
		} else {
			res = append(res, stack[len(stack)-1].Val)
			stack = stack[:len(stack)-1]
		}
	}
	return res

}

// @lc code=end
