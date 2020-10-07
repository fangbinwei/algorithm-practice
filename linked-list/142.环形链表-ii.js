/*
 * @lc app=leetcode.cn id=142 lang=javascript
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

 function ListNode(val) {
    this.val = val;
    this.next = null;
 }

// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var detectCycle = function(head) {
  if (head == null || head.next == null) {
    return null
  }
  let runner1 = head
  let runner2 = head
  let firstMeet = null
  while (true) {
    if (runner1 === null) {
      return null
    }

    if (runner2 === null) {
      return null
    }

    if (runner2.next === null) {
      return null
    }
    runner1 = runner1.next
    runner2 = runner2.next.next

    if (runner1 === runner2) {
      // 第一次相遇
      // 起始点到环起点的距离为l
      // 环的周长为S, 第一次相遇距离环起点x
      // l + x = v*t
      // l + S + x = 2v*t
      // 所以 S = l + x
      firstMeet = runner1      
      break
    }
  }

  // 在第一次相遇的点, 用v速度走完剩下圈的路程S-x = l,
  // 同时用速度v在head出发, 刚好在环的起点相遇
  runner1 = head
  runner2 = firstMeet
  while (true) {
    if (runner1 === runner2) {
      return runner1
    }
    runner1 = runner1.next
    runner2 = runner2.next
  }
};
// @lc code=end

