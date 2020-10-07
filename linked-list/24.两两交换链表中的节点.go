package leetcode_24

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (66.32%)
 * Likes:    583
 * Dislikes: 0
 * Total Accepted:    134.9K
 * Total Submissions: 203.5K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例:
 *
 * 给定 1->2->3->4, 你应该返回 2->1->4->3.
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
// 遍历, 利用哨兵
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{-1, head}

	prev := dummy
	current := head

	for current != nil && current.Next != nil {
		second := current.Next
		nextGroup := second.Next

		// 第一次执行的时候, 会将dummy.Next 执向new head
		prev.Next = second
		second.Next = current
		current.Next = nextGroup

		// prev 指向group中的第二个 , current已经不再是第一个
		prev = current
		current = nextGroup
	}

	return dummy.Next
}

// @lc code=end

// 遍历无哨兵
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h := head.Next

	current := head
	for current != nil && current.Next != nil {
		tmp := current.Next.Next
		current.Next.Next = current
		// 后一组满两个
		if tmp != nil && tmp.Next != nil {
			current.Next = tmp.Next
			// 后一组不满两个, 指向第一个(*ListNode),tmp可能为nil
		} else {
			current.Next = tmp
		}
		current = tmp
	}
	return h
}

// 递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next

	remian := swapPairs(head.Next.Next)
	head.Next.Next = head
	head.Next = remian

	return newHead

}
