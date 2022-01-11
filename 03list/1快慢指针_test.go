/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/10 3:27 下午
 * @Desc:
 */

package tlist

import (
	"fmt"
	"testing"
)

/**
 * 快慢指针
 * 1。输入链表头节点，奇数长度返回中点，偶数长度返回上中点
 * 2。输入链表头节点，奇数长度返回中点，偶数长度返回下中点
 * 3。输入链表头节点，奇数长度返回中点前一个，偶数长度返回上中点前一个
 * 4。输入链表头节点，奇数长度返回中点前一个，偶数长度返回下中点前一个
 */
func TestMidOrUpMid(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}
	//n6 := &Node{Value: 6}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	//n5.Next = n6

	fmt.Println(getMidOrUpMid(n1))
}
func getMidOrUpMid(node *Node) *Node {
	if node == nil || node.Next == nil {
		return node
	}
	slowPointer := node
	fastPointer := node
	for fastPointer.Next != nil && fastPointer.Next.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}
	return slowPointer
}

//2. mid or down mid node
func getMidOrDownMid(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	slowPointer := head
	fastPointer := head.Next
	for fastPointer.Next != nil && fastPointer.Next.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}
	return slowPointer.Next
}
func TestMidOrDownMid(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	//n3 := &Node{Value: 3}
	//n4 := &Node{Value: 4}
	//n5 := &Node{Value: 5}
	//n6 := &Node{Value: 6}

	n1.Next = n2
	//n2.Next = n3
	//n3.Next = n4
	//n4.Next = n5
	//n5.Next = n6

	fmt.Println(getMidOrDownMid(n1))
}

//3.输入链表头节点，奇数长度返回中点前一个，偶数长度返回上中点前一个
func getBeforeMidOrBeforeUpMid(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	slowPointer := head
	fastPointer := head.Next.Next

	for fastPointer.Next != nil && fastPointer.Next.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next
	}
	return slowPointer
}

func TestGetBeforeMidOrBeforeUpMid(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}
	n6 := &Node{Value: 6}
	n7 := &Node{Value: 7}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7

	fmt.Println(getBeforeMidOrBeforeUpMid(n1))
}

//4.输入链表头节点，奇数长度返回中点前一个，偶数长度返回下中点前一个
func getBeforeMidOrBeforeDownMid(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Next.Next == nil {
		return head
	}
	slowPointer := head
	fastPointer := head.Next
	for fastPointer.Next != nil && fastPointer.Next.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}
	return slowPointer
}
func TestGetBeforeMidOrBeforeDownMid(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}
	n6 := &Node{Value: 6}
	n7 := &Node{Value: 7}
	n8 := &Node{Value: 8}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7
	n7.Next = n8

	fmt.Println(getBeforeMidOrBeforeDownMid(n1))
}
