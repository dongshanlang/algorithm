/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/13 5:15 下午
 * @Desc:
 */

package _4binary_tree

import (
	"container/list"
	"fmt"
	"testing"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func preTraverse(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Value)
	preTraverse(node.Left)
	preTraverse(node.Right)
}
func midTraverse(node *Node) {
	if node == nil {
		return
	}
	midTraverse(node.Left)
	fmt.Printf("%d ", node.Value)
	midTraverse(node.Right)
}
func afterTraverse(node *Node) {
	if node == nil {
		return
	}
	afterTraverse(node.Left)
	afterTraverse(node.Right)
	fmt.Printf("%d ", node.Value)
}
func TestPreTraverse(t *testing.T) {
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
	preTraverse(n1)
	fmt.Println("==========")
	midTraverse(n1)
	fmt.Println("==========")
	afterTraverse(n1)
}
func levelTraversal(head *Node) {
	l := list.New()
	l.PushBack(head)
	for node := l.Front(); node != nil; node = l.Front() {
		fmt.Printf("%d ", node.Value.(*Node).Value)
		if node.Value.(*Node).Left != nil {
			l.PushBack(node.Value.(*Node).Left)
		}
		if node.Value.(*Node).Right != nil {
			l.PushBack(node.Value.(*Node).Right)
		}
		l.Remove(node)
	}

}
func TestLevelTraversal(t *testing.T) {
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
	levelTraversal(n1)
}
