package main

/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 *
 * https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/description/
 *
 * algorithms
 * Hard (51.81%)
 * Likes:    344
 * Dislikes: 0
 * Total Accepted:    47.5K
 * Total Submissions: 91.6K
 * Testcase Example:  '[1,2,3,null,null,4,5]'
 *
 *
 * 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
 *
 * 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 /
 * 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
 *
 * 示例:
 *
 * 你可以将以下二叉树：
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * ⁠    / \
 * ⁠   4   5
 *
 * 序列化为 "[1,2,3,null,null,4,5]"
 *
 * 提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode
 * 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
 *
 * 说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。
 *
 */

// @lc code=start
import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	// 用来记录序列化后的值, 用于反序列化
	s []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := make([]string, 0, 10)

	// 前序遍历
	result = append(result, strconv.Itoa(root.Val))
	result = append(result, this.serialize(root.Left))
	result = append(result, this.serialize(root.Right))

	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.s = strings.Split(data, ",")
	// 也可以用一个函数, 然后把s的指针传递过去, 不用this.s的方式
	return this._deserialize2()

}

func (this *Codec) _deserialize() *TreeNode {
	if this.s[0] == "" {
		this.s = this.s[1:]
		return nil
	}

	v, _ := strconv.Atoi(this.s[0])
	this.s = this.s[1:]
	root := &TreeNode{Val: v}
	root.Left = this._deserialize()
	root.Right = this._deserialize()
	return root
}

// 使用遍历的方式
func (this *Codec) _deserialize2() *TreeNode {
	s := this.s
	if s[0] == "" && len(s) == 1 {
		return nil
	}

	generateNode := func() *TreeNode {
		s0 := s[0]
		s = s[1:]
		if s0 == "" {
			return nil
		}

		v, _ := strconv.Atoi(s0)
		return &TreeNode{Val: v}
	}

	v, _ := strconv.Atoi(s[0])
	s = s[1:]
	root := &TreeNode{Val: v}
	// 利用栈 进行前序遍历, 参考144
	stack := make([]*TreeNode, 0, 10)
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {

			stack = append(stack, cur)
			cur.Left = generateNode()
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur.Right = generateNode()
			cur = cur.Right
		}
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end
