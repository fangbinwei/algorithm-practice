package leetcode236

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (68.44%)
 * Likes:    735
 * Dislikes: 0
 * Total Accepted:    125.1K
 * Total Submissions: 182.8K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
// 迭代
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solution/cong-qian-xu-yu-zhong-xu-bian-li-xu-lie-gou-zao-9/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	stack := make([]*TreeNode, 0, 10)
	stack = append(stack, root)
	var inorderIndex int

	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		// 和栈顶的元素比较, 如果相同,
		top := stack[len(stack)-1]
		if top.Val != inorder[inorderIndex] {
			top.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, top.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				top = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			top.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, top.Right)
		}
	}

	return root
}

// @lc code=end

// 递归
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}

	var k int
	// 可以用hashmap优化这里的搜索 map[*TreeNode]int
	for i, v := range inorder {
		if v == preorder[0] {
			k = i
			break
		}
	}
	root.Left = buildTree(preorder[1:k+1], inorder[:k])
	root.Right = buildTree(preorder[k+1:], inorder[k+1:])
	return root
}
