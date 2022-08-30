/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/1/17 11:08 上午
 * @Desc:
 */

package tlist

import (
	"fmt"
	"testing"
)

func ReverseList(pHead *Node) *Node {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}
	cur := pHead
	next := cur.Next
	pHead.Next = nil
	for next != nil {
		tmp := next.Next
		next.Next = cur
		cur = next
		next = tmp
	}

	return cur
}
func Test(t *testing.T) {
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
	head := ReverseList(n1)
	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
}

func DoubleListReverse(node *DoubleNode) *DoubleNode {
	if node == nil {
		return node
	}
	var pre, next, cur *DoubleNode
	pre = node
	cur = pre.Next
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		cur.Pre = next
		pre = cur
		cur = next
	}
	node.Next = nil
	return pre
}
func TestDoubleListReverse(t *testing.T) {
	n1 := &DoubleNode{Value: 1}
	n2 := &DoubleNode{Value: 2}
	n3 := &DoubleNode{Value: 3}
	n4 := &DoubleNode{Value: 4}
	n5 := &DoubleNode{Value: 5}
	n6 := &DoubleNode{Value: 6}
	n7 := &DoubleNode{Value: 7}

	n1.Next = n2
	n1.Pre = nil
	n2.Next = n3
	n2.Pre = n1
	n3.Next = n4
	n3.Pre = n2
	n4.Next = n5
	n4.Pre = n3
	n5.Next = n6
	n5.Pre = n4
	n6.Next = n7
	n6.Pre = n5
	n7.Pre = n6
	n7.Next = nil
	head := DoubleListReverse(n1)
	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
}
