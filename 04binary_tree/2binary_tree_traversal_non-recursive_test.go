/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/14 10:47 上午
 * @Desc:
 */

package _4binary_tree

import (
	"algorithm/utils/mstack"
	"fmt"
	"testing"
)

//中序
//左中右
func midNonRecursiveTraverse(head *Node) {
	if head == nil {
		return
	}
	s := mstack.New()
	var node = head
	for node != nil || s.Peek() != nil {
		for node != nil {
			//fmt.Println(node.Value)//先序遍历
			s.Push(node)
			node = node.Left
		}
		if s.Peek() != nil {
			node = s.Pop().(*Node)
			fmt.Println(node.Value) //中序遍历
			node = node.Right
		}
	}
}
func TestMidNonRecursiveTraverse(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}
	n6 := &Node{Value: 6}
	n7 := &Node{Value: 7}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	midNonRecursiveTraverse(n1)
}

//后序遍历
func postOrderNonRecursiveTraverse(head *Node) {
	if head == nil {
		return
	}
	var mark = head
	s := mstack.New()
	s.Push(head)
	for !s.IsEmpty() {
		currentNode := s.Peek().(*Node)
		if currentNode.Left != nil && mark != currentNode.Left && mark != currentNode.Right {
			s.Push(currentNode.Left)
		} else if currentNode.Right != nil && mark != currentNode.Right {
			s.Push(currentNode.Right)
		} else {
			fmt.Println(s.Pop().(*Node).Value)
			mark = currentNode
		}
	}
}
func TestPostOrderNonRecursiveTraverse(t *testing.T) {

	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}
	n6 := &Node{Value: 6}
	n7 := &Node{Value: 7}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	postOrderNonRecursiveTraverse(n1)

}

//先序
//借助栈结构实现，中左右，
//按照栈先进后出的特点，应当先放进右节点。
func preNonRecursiveTraverse(head *Node) {
	if head == nil {
		return
	}
	s := mstack.New()
	s.Push(head)
	var node *Node
	for s.Peek() != nil {
		node = s.Pop().(*Node)
		fmt.Println(node.Value)
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}
}
func TestPreNonRecursiveTraverse(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}
	n6 := &Node{Value: 6}
	n7 := &Node{Value: 7}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	preNonRecursiveTraverse(n1)
}
