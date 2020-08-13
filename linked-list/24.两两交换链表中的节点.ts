/*
 * @lc app=leetcode.cn id=24 lang=typescript
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

class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val===undefined ? 0 : val)
        this.next = (next===undefined ? null : next)
    }
}
// @lc code=start
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function swapPairs(head: ListNode | null): ListNode | null {
  if (head == null || head.next == null) return head
  let secondNode = head.next
  head.next = swapPairs(secondNode.next)
  secondNode.next = head
  return secondNode
};
// @lc code=end


function swapPairs(head: ListNode | null): ListNode | null {
  if (head == null || head.next == null) return head

  let current: ListNode | null = head
  let newFirst: ListNode | null = head.next
  while(current && current.next) {
    let secondNode: ListNode = current.next
    current.next = 
      secondNode.next 
      // 需要看下一对是否完整,
      // 1-> 2 -> 3->4  1要指向4, 如果没有4, 要指向3
        ? (secondNode.next.next || secondNode.next)
        : null

    let tmp2: ListNode| null = secondNode.next
    secondNode.next = current
    current =  tmp2
  }
  return newFirst
};