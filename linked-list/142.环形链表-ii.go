/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 *
 * https://leetcode-cn.com/problems/linked-list-cycle-ii/description/
 *
 * algorithms
 * Medium (51.33%)
 * Likes:    575
 * Dislikes: 0
 * Total Accepted:    103.6K
 * Total Submissions: 201.9K
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
 *
 * 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
 *
 * 说明：不允许修改给定的链表。
 *
 *
 *
 * 示例 1：
 *
 * 输入：head = [3,2,0,-4], pos = 1
 * 输出：tail connects to node index 1
 * 解释：链表中有一个环，其尾部连接到第二个节点。
 *
 *
 *
 *
 * 示例 2：
 *
 * 输入：head = [1,2], pos = 0
 * 输出：tail connects to node index 0
 * 解释：链表中有一个环，其尾部连接到第一个节点。
 *
 *
 *
 *
 * 示例 3：
 *
 * 输入：head = [1], pos = -1
 * 输出：no cycle
 * 解释：链表中没有环。
 *
 *
 *
 *
 *
 *
 * 进阶：
 * 你是否可以不用额外空间解决此题？
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	// 快慢指针, 如果有环, 它们会相遇,
	// 慢指针速度为1, 快指针速度为2
	// 起始点到入环点的距离为A
	// 假设相遇后慢指针路程 A + B,
	// 快指针路程 A + B + N
	// 快指针 因为速度为慢指针的两倍, 所以路程也是两倍2A + 2B
	// 快指针超了慢指针1圈, 所以周长N = A + B, 当前在B位置, 所以继续走 到入环的第一个节点距离为A
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head

	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			for fast != head {
				fast = fast.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}

// @lc code=end

// 利用hash表记录走过的节点, 入环的节点为第一个重复的点
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})

	if head == nil || head.Next == nil {
		return nil
	}
	m[head] = struct{}{}
	for current := head.Next; current != nil; current = current.Next {

		if _, ok := m[current]; ok {
			return current
		}
		m[current] = struct{}{}
	}
	return nil
}

// 空间O(1), 不考虑空间, 可以用hash table保存
func detectCycle(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}
	guard := head
	for current := head; current != nil; current = current.Next {
		if current.Next == nil {
			return nil
		}
		if current.Next == current {
			return current
		}
		// 不断和前面的节点比对
		for ; guard != current; guard = guard.Next {
			if current.Next == guard {
				return guard
			}
		}
		guard = head
	}
	return nil
}