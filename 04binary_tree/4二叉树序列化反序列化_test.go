/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/14 11:14 上午
 * @Desc:
 */

package _4binary_tree

import (
	"algorithm/utils/mqueue"
	"fmt"
	"go/types"
	"testing"
)

//二叉树的序列化反序列化
//1。可以用先序或者中序或者后序护着按层遍历，来实现二叉树的序列化
//2。用了什么方式序列化，就用什么方式反序列化
func SerializeTree(head *Node) *mqueue.MQueue {
	if head == nil {
		return nil
	}
	l := mqueue.New()
	l.PushBack(head)
	serializeList := mqueue.New()
	serializeList.PushBack(head)
	var node *Node
	for !l.IsEmpty() {
		node = l.PopFront().(*Node)
		if node.Left != nil {
			serializeList.PushBack(node.Left)
			l.PushBack(node.Left)
		} else {
			serializeList.PushBack(nil)
		}
		if node.Right != nil {
			l.PushBack(node.Right)
			serializeList.PushBack(node.Right)
		} else {
			serializeList.PushBack(nil)
		}
		//fmt.Printf("%+v ", node)
	}
	//for !serializeList.IsEmpty() {
	//	item := serializeList.PopFront()
	//	if item == nil {
	//		fmt.Printf("%+v ", item)
	//	} else {
	//		fmt.Printf("%+v ", item.(*Node).Value)
	//	}
	//}
	return serializeList
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
	serialization := SerializeTree(n1)
	head := Deserialization(serialization)
	fmt.Println(head)
}
func Deserialization(list *mqueue.MQueue) *Node {
	if list.IsEmpty() {
		return nil
	}
	var node *Node
	var queue = mqueue.New()
	var head = generateNode(list.PopFront())
	queue.PushBack(head)
	for !queue.IsEmpty() {
		node = queue.PopFront().(*Node)
		node.Left = generateNode(list.PopFront())
		node.Right = generateNode(list.PopFront())
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return head
}
func generateNode(value interface{}) *Node {
	switch value.(type) {
	case *Node:
		node := value.(*Node)
		return node
	case types.Nil:
		return nil
	case interface{}:
		return nil
	case int:
		return &Node{Value: value.(int)}
	default:
		fmt.Println("unknown type")
		return nil
	}
}
