/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/3/15 10:17 上午
 * @Desc:
 */

//描述
//将一个节点数为 size 链表 m 位置到 n 位置之间的区间反转，要求时间复杂度 O(n)O(n)，空间复杂度 O(1)O(1)。
//例如：
//给出的链表为 1\to 2 \to 3 \to 4 \to 5 \to NULL1→2→3→4→5→NULL, m=2,n=4m=2,n=4,
//返回 1\to 4\to 3\to 2\to 5\to NULL1→4→3→2→5→NULL.
//
//数据范围： 链表长度 0 < size <= 10000, 0 < m <= n <= size，链表中每个节点的值满足 |val| <= 1000
//要求：时间复杂度 O(n)O(n) ，空间复杂度 O(n)O(n)
//进阶：时间复杂度 O(n)O(n)，空间复杂度 O(1)O(1)
//示例1
//输入：
//{1,2,3,4,5},2,4
//复制
//返回值：
//{1,4,3,2,5}
//复制
//示例2
//输入：
//{5},1,1
//复制
//返回值：
//{5}

package tlist

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestReverseBetween(t *testing.T) {
	head := initList()
	newHead := reverseBetween(head, 1, 6)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}

/**
 *
 * @param head ListNode类
 * @param m int整型
 * @param n int整型
 * @return ListNode类
 */

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	beforeMidHead := head
	for i := 2; i < m; i++ {
		beforeMidHead = beforeMidHead.Next
	}
	afterMidTail := head
	for i := 0; i < n; i++ {
		afterMidTail = afterMidTail.Next
	}
	var midHead, midTail *ListNode
	if m == 1 {
		midHead, midTail = reverseNList(beforeMidHead, n-m)
	} else {
		midHead, midTail = reverseNList(beforeMidHead.Next, n-m)
	}
	midTail.Next = afterMidTail
	if m != 1 {
		beforeMidHead.Next = midHead
		return head
	} else {
		return midHead
	}
}

func reverseNList(head *ListNode, count int) (newStart, newStop *ListNode) {
	if head == nil || head.Next == nil || count == 0 {
		return head, head
	}
	preStep := head
	curStep := preStep.Next
	nextStep := curStep.Next
	for i := 0; curStep != nil && i < count; i++ {
		nextStep = curStep.Next
		curStep.Next = preStep
		preStep = curStep
		curStep = nextStep
	}
	head.Next = nil

	return preStep, head
}
func TestReverseListN(t *testing.T) {
	head := initList()
	newStart, newStop := reverseNList(head, 2)
	for newStart != nil {
		fmt.Println(newStart.Val)
		newStart = newStart.Next
	}
	fmt.Println(newStop.Val)
}
func initList() *ListNode {
	n1 := ListNode{
		Val:  1,
		Next: nil,
	}
	n2 := ListNode{
		Val:  2,
		Next: nil,
	}
	n3 := ListNode{
		Val:  3,
		Next: nil,
	}
	n4 := ListNode{
		Val:  4,
		Next: nil,
	}
	n5 := ListNode{
		Val:  5,
		Next: nil,
	}
	//n6 := ListNode{
	//	Val:  6,
	//	Next: nil,
	//}

	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5
	//n5.Next = &n6
	return &n1
}
