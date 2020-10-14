package leetcode_641
/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 *
 * https://leetcode-cn.com/problems/design-circular-deque/description/
 *
 * algorithms
 * Medium (53.29%)
 * Likes:    60
 * Dislikes: 0
 * Total Accepted:    12.2K
 * Total Submissions: 22.8K
 * Testcase Example:  '["MyCircularDeque","insertLast","insertLast","insertFront","insertFront","getRear","isFull","deleteLast","insertFront","getFront"]\n' +
  '[[3],[1],[2],[3],[4],[],[],[],[4],[]]'
 *
 * 设计实现双端队列。
 * 你的实现需要支持以下操作：
 * 
 * 
 * MyCircularDeque(k)：构造函数,双端队列的大小为k。
 * insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
 * insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
 * deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
 * deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
 * getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
 * getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
 * isEmpty()：检查双端队列是否为空。
 * isFull()：检查双端队列是否满了。
 * 
 * 
 * 示例：
 * 
 * MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
 * circularDeque.insertLast(1);                    // 返回 true
 * circularDeque.insertLast(2);                    // 返回 true
 * circularDeque.insertFront(3);                    // 返回 true
 * circularDeque.insertFront(4);                    // 已经满了，返回 false
 * circularDeque.getRear();                  // 返回 2
 * circularDeque.isFull();                        // 返回 true
 * circularDeque.deleteLast();                    // 返回 true
 * circularDeque.insertFront(4);                    // 返回 true
 * circularDeque.getFront();                // 返回 4
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 所有值的范围为 [1, 1000]
 * 操作次数的范围为 [1, 1000]
 * 请不要使用内置的双端队列库。
 * 
 * 
 */

// 使用双向链表解决
// 也可以使用slice, head, tail分别指向头部和尾部
// head tail的值随着增删变化, 可以避免unshift操作O(n)的复杂度
// @lc code=start
type node struct {
  val int
  next *node
  prev *node
}
type MyCircularDeque struct {
  cap int
  length int
  head *node
  last *node
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
  return MyCircularDeque{k, 0, nil, nil}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
  if this.IsFull() {
    return false
  }
  if this.length == 0 {
    this.head  = &node{value, nil, nil}
    this.last = this.head
  } else {
    tmp := this.head
    this.head = &node{value, tmp,nil}
    tmp.prev = this.head
  }
  this.length++
  return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
  if this.IsFull() {
    return false
  }
  if this.length == 0 {
    this.head  = &node{value, nil, nil}
    this.last = this.head
  } else {
    tmp := this.last
    this.last = &node{value, nil, tmp}
    tmp.next = this.last
  }
  this.length++
  return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
  if this.IsEmpty() {
    return false
  }
  if this.head.next!= nil {
    this.head = this.head.next
    this.head.prev = nil
  } else {
    this.head = nil
    this.last = nil
  }
  this.length--
  return true

}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
  if this.IsEmpty() {
    return false
  }
  if this.length == 1 {
    this.last = nil
    this.head = nil
  } else {
    this.last = this.last.prev
    this.last.next = nil
  }
  this.length--
  return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
  if this.IsEmpty() {
    return -1
  }
  return this.head.val
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
  if this.IsEmpty() {
    return -1
  }
  return this.last.val
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
  return this.length == 0

}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
  return this.length == this.cap
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end

