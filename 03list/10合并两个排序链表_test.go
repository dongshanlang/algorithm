/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/3/18 10:17 上午
 * @Desc:
 */

package tlist

import (
	"fmt"
	"testing"
)

//描述
//输入两个递增的链表，单个链表的长度为n，合并这两个链表并使新链表中的节点仍然是递增排序的。
//数据范围： 0≤n≤1000，−1000≤节点值≤1000
//要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
//
//如输入{1,3,5},{2,4,6}时，合并后的链表为{1,2,3,4,5,6}，所以对应的输出为{1,2,3,4,5,6}，转换过程如下图所示：
//或输入{-1,2,4},{1,3,4}时，合并后的链表为{-1,1,2,3,4,4}，所以对应的输出为{-1,1,2,3,4,4}，转换过程如下图所示：
//
//示例1
//输入：
//{1,3,5},{2,4,6}
//返回值：
//{1,2,3,4,5,6}
//示例2
//输入：
//{},{}
//返回值：
//{}
//示例3
//输入：
//{-1,2,4},{1,3,4}
//返回值：
//{-1,1,2,3,4,4}
/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil {
		return pHead2
	} else if pHead2 == nil {
		return pHead1
	}
	step1 := pHead1
	step2 := pHead2
	newHead := &ListNode{}
	newTail := newHead
	for step1 != nil && step2 != nil {
		if step1.Val < step2.Val {
			newTail.Next = step1
			step1 = step1.Next
		} else {
			newTail.Next = step2
			step2 = step2.Next
		}
		newTail = newTail.Next
	}
	if step1 == nil && step2 != nil {
		newTail.Next = step2
	} else if step2 == nil && step1 != nil {
		newTail.Next = step1
	}
	return newHead.Next
}
func TestMergeSortedList(t *testing.T) {
	head1 := initList1()
	head2 := initList2()
	newHead := Merge(head1, head2)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}

func initList2() *ListNode {
	n1 := ListNode{
		Val:  1,
		Next: nil,
	}
	n2 := ListNode{
		Val:  5,
		Next: nil,
	}
	n3 := ListNode{
		Val:  7,
		Next: nil,
	}

	n1.Next = &n2
	n2.Next = &n3
	return &n1
}

func initList1() *ListNode {
	n1 := ListNode{
		Val:  2,
		Next: nil,
	}
	n2 := ListNode{
		Val:  3,
		Next: nil,
	}
	n3 := ListNode{
		Val:  9,
		Next: nil,
	}

	n1.Next = &n2
	n2.Next = &n3
	return &n1
}
