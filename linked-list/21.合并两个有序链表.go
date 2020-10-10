package leetcode_21

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (64.49%)
 * Likes:    1307
 * Dislikes: 0
 * Total Accepted:    385K
 * Total Submissions: 597K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 *
 *
 * 示例：
 *
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 *
 *
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	pre := &ListNode{-1, nil}

	if l1 != nil && l2 != nil {

		// pre指向较小的值
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre.Next.Next = mergeTwoLists(l1,l2)
	}

	return pre.Next

}

// @lc code=end

// 遍历
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	c1, c2 := l1, l2

	pre := &ListNode{-1, nil}
	head := pre

	for c1 != nil && c2 != nil {

		// pre指向较小的值
		if c1.Val <= c2.Val {
			pre.Next = c1
			c1 = c1.Next
		} else {
			pre.Next = c2

			c2 = c2.Next
		}
		// 更新pre
		pre = pre.Next
	}
	// 处理剩余的一个节点
	if c1 != nil {
		pre.Next = c1
	}

	if c2 != nil {
		pre.Next = c2
	}
	return head.Next

}