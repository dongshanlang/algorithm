/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/3/17 10:28 上午
 * @Desc:
 */

package tlist

import (
	"fmt"
	"testing"
)

//描述
//将给出的链表中的节点每 k 个一组翻转，返回翻转后的链表
//如果链表中的节点数不是 k 的倍数，将最后剩下的节点保持原样
//你不能更改节点中的值，只能更改节点本身。
//
//数据范围：  0≤n≤2000 ，链表中每个元素都满足0≤val≤1000
//要求空间复杂度 O(1)，时间复杂度O(n)
//例如：
//给定的链表是 1→2→3→4→5
//对于 k=2 , 你应该返回 2→1→4→3→5
//对于 k=3 , 你应该返回3→2→1→4→5
//
//示例1
//输入：
//{1,2,3,4,5},2
//返回值：
//{2,1,4,3,5}
//示例2
//输入：
//{},1
//返回值：
//{}

/**
 *
 * @param head ListNode类
 * @param k int整型
 * @return ListNode类
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	var partitionHead, partitionTail, beforePartitionHead, afterPartitionTail *ListNode
	ListHead := &ListNode{
		Val:  0,
		Next: head,
	}
	beforePartitionHead = ListHead
	afterPartitionTail = ListHead
	partitionTail = ListHead
loop:
	for {
		//找到分组的尾节点
		for i := 0; i < k; i++ {
			partitionTail = partitionTail.Next
			if partitionTail == nil {
				break loop
				//结束循环
			}
		}
		//标记尾节点的下一个
		afterPartitionTail = partitionTail.Next
		//已经有反转头partitionHead，并且分组k个节点没有到nil
		partitionHead, partitionTail = reverseK(beforePartitionHead.Next, k-1)
		//组合链表
		beforePartitionHead.Next = partitionHead
		partitionTail.Next = afterPartitionTail
		beforePartitionHead = partitionTail
	}
	return ListHead.Next
}

func TestReverseKGroup(t *testing.T) {
	head := initList()
	newHead := reverseKGroup(head, 2)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}

func reverseK(node *ListNode, k int) (head, tail *ListNode) {
	if node == nil || node.Next == nil {
		return head, head
	}
	pre := node
	cur := pre.Next
	var next *ListNode
	for i := 0; cur != nil && i < k; i++ {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	node.Next = nil
	return pre, node
}
func TestReverseK(t *testing.T) {
	head := initList()
	newHead, newTail := reverseK(head, 1)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
	fmt.Println(newTail.Val)
}
