/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/14 11:14 上午
 * @Desc:
 */

package _4binary_tree

import (
	"algorithm/utils/mlist"
	"fmt"
	"testing"
)

//二叉树的序列化反序列化
//1。可以用先序或者中序或者后序护着按层遍历，来实现二叉树的序列化
//2。用了什么方式序列化，就用什么方式反序列化
func SerializeTree(head *Node) {
	if head == nil {
		return
	}
	l := mlist.New()
	l.PushBack(head)
	serializeList := mlist.New()

	var node *Node
	for !l.IsEmpty() {
		node = l.PopFront().(*Node)
		if node.Left != nil {
			l.PushBack(node.Left)
		} else {

		}
		if node.Right != nil {
			l.PushBack(node.Right)
		} else {

		}
		fmt.Printf("%d ", node.Value)
	}
}
func TestSerialiseTree(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}
	n6 := &Node{Value: 6}
	n7 := &Node{Value: 7}
	n8 := &Node{Value: 8}
	n9 := &Node{Value: 9}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	n5.Left = n8
	n5.Right = n9
	SerializeTree(n1)
}
