/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/10 4:49 下午
 * @Desc:
 */

package tlist

import (
	"fmt"
	"testing"
)

//一种特殊的单链表节点描述如下
type SpecialNode struct {
	Value int
	Next  *SpecialNode
	Rand  *SpecialNode
}

//rand指针是单链表节点结构中新增的指针，rand可能指向链表中的任意一个节点，也可能指向null
//给定一个由Node节点类型组成的无环单链表的头节点head，请实现一个函数完成这个链表的复制，并返回复制的新链表的头节点。
//要求：
//时间复杂度是O(N)，额外空间复杂度是O(1)

func (sn *SpecialNode) New(value int) {
	sn.Value = value
}
func (sn *SpecialNode) Copy() *SpecialNode {
	return &SpecialNode{
		Value: sn.Value,
		Next:  sn.Next,
		Rand:  sn.Rand,
	}
}
func TestSNode(t *testing.T) {
	sn1 := SpecialNode{}
	fmt.Println(sn1)
}
func copySpecialList(head *SpecialNode) *SpecialNode {
	if head == nil {
		return head
	}
	//1.range for copy next
	rangePointer := head
	for rangePointer != nil {
		next := rangePointer.Next
		sn := rangePointer.Copy()
		rangePointer.Next = sn
		sn.Next = next
		rangePointer = next
	}
	//2.range for copy Rand pointer
	rangePointer = head
	for rangePointer != nil {
		sn := rangePointer.Next
		if rangePointer.Rand != nil {
			sn.Rand = rangePointer.Rand.Next
		}
		rangePointer = rangePointer.Next.Next
	}
	//3.range to reduction
	var retHead, newList, oldList *SpecialNode
	retHead = head.Next
	newList = head.Next
	oldList = head
	rangePointer = head.Next.Next
	for rangePointer != nil {
		newList.Next = rangePointer.Next
		newList = newList.Next
		oldList.Next = rangePointer
		oldList = oldList.Next
		rangePointer = rangePointer.Next.Next
	}
	return retHead
}
func TestCopySpecialNode(t *testing.T) {
	n1 := &SpecialNode{Rand: nil, Value: 1}
	n2 := &SpecialNode{Rand: n1, Value: 2}
	n3 := &SpecialNode{Rand: n1, Value: 3}
	n4 := &SpecialNode{Rand: n3, Value: 4}
	n5 := &SpecialNode{Rand: n2, Value: 5}
	n6 := &SpecialNode{Rand: n1, Value: 6}
	n7 := &SpecialNode{Rand: n4, Value: 7}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7
	newList := copySpecialList(n1)
	fmt.Println(newList)
}
