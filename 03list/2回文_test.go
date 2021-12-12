/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/10 4:24 下午
 * @Desc:
 */

package _3list

import (
	"fmt"
	"testing"
)

//给定一个单链表的头节点head，请判断该链表是否为回文结构
//1）栈方法简单
//2）改原链表的方法需要注意边界

//1.中点或上中点
//2。逆序
//3比对
//4调整回原序列
func checkForPalindromicString(head *Node) bool {
	if head == nil || head.Next == nil {
		return true
	}
	//找到中点或者上中点
	midPointer := head
	fastPointer := head
	for fastPointer.Next != nil && fastPointer.Next.Next != nil {
		midPointer = midPointer.Next
		fastPointer = fastPointer.Next.Next
	}

	//中点之后的链表反转
	backListHead := reverseList(midPointer.Next)
	range2Pointer := backListHead
	palindromicString := false
	for range1Pointer := head; range1Pointer != nil; range1Pointer = range1Pointer.Next {
		if range2Pointer == nil {
			palindromicString = true
			break
		}
		if range1Pointer.Value != range2Pointer.Value {
			palindromicString = false
			break
		}
		range2Pointer = range2Pointer.Next
	}
	//恢复链表
	_ = reverseList(backListHead)

	return palindromicString
}
func TestCheckForPalindrome(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 13}
	n4 := &Node{Value: 18}
	n5 := &Node{Value: 13}
	n6 := &Node{Value: 2}
	n7 := &Node{Value: 2}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7
	fmt.Println(checkForPalindromicString(n1))
}
func reverseList(node *Node) *Node {
	if node == nil || node.Next == nil {
		return node
	}
	var (
		pre, cur, next *Node
	)
	pre = node
	cur = pre.Next
	if node.Next.Next == nil {
		cur.Next = pre
		pre.Next = nil
		return cur
	}
	next = cur.Next
	pre.Next = nil
	for next != nil {
		cur.Next = pre
		pre = cur
		cur = next
		next = next.Next
	}
	cur.Next = pre
	return cur
}
func TestReverseList(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	//n3 := &Node{Value: 3}
	//n4 := &Node{Value: 4}
	//n5 := &Node{Value: 5}
	//n6 := &Node{Value: 6}
	//n7 := &Node{Value: 7}

	n1.Next = n2
	//n2.Next = n3
	//n3.Next = n4
	//n4.Next = n5
	//n5.Next = n6
	//n6.Next = n7
	reverseHead := reverseList(n1)
	fmt.Println(reverseHead)
}
