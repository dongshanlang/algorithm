/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/10 5:57 下午
 * @Desc:
 */

package tlist

import (
	"fmt"
	"testing"
)

//给定两个可能有环也可能无环的单链表，头节点head1和head2。请实现一个函数，
//如果两个链表相交，返回相交的第一个节点。如果不相交，返回null
//要求：
//如果两个链表长度之和为N，时间复杂度清达到O(N)，额外空间复杂度O(1)

type Node struct {
	Value int
	Next  *Node
}

//找到链表第一个如环节点，如果无环返回null
func getLoopNode(head Node) {
	//1.快指针、满指针同时走
	//2, 相遇时，快指针回到起点，一步一步走，慢指针继续，下次相等时，就是环节点///为啥？
}

//给定链表判定是否有环
func checkListRing(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var slowPointer, fastPointer *Node
	slowPointer = head
	fastPointer = head.Next
	for fastPointer.Next != nil && fastPointer.Next.Next != nil {
		if slowPointer == fastPointer {
			return true
		}
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}
	return false
}
func TestCheckListRing(t *testing.T) {
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
	fmt.Println(checkListRing(n1))
}

//返回环的第一个节点，没环返回空
func getRingNode(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}
	slowPointer := head
	fastPointer := head.Next
	for fastPointer.Next != nil && fastPointer.Next.Next != nil {
		if slowPointer == fastPointer {
			//相交
			slowPointer = slowPointer.Next
			fastPointer = head
			for slowPointer != fastPointer {
				slowPointer = slowPointer.Next
				fastPointer = fastPointer.Next
				fmt.Println(slowPointer.Value, fastPointer.Value)
			}
			return slowPointer
		}
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}
	return nil
}
func TestGetRingNode(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}
	n6 := &Node{Value: 6}
	//n7 := &Node{Value: 2}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n1
	fmt.Println(getRingNode(n1))
}

//给定两个可能有环也可能无环的单链表，头节点head1和head2。请实现一个函数，
//如果两个链表相交，返回相交的第一个节点。如果不相交，返回null
//要求：
//如果两个链表长度之和为N，时间复杂度清达到O(N)，额外空间复杂度O(1)
func getListIntersectNode(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}

	//是否有环
	listRing := checkListRing(head1)
	list2Ring := checkListRing(head2)
	if listRing != list2Ring { //一个有环一个无环则不相交
		return nil
	}

	if listRing { //1. 有环: 不相交、环内相交、环外相交
		var ringNode1, ringNode2 *Node
		ringNode1 = getRingNode(head1)
		ringNode2 = getRingNode(head2)
		tmpRing1 := ringNode1.Next
		ringNode1.Next = nil
		tmpRing2 := ringNode2.Next
		ringNode2.Next = nil
		intersectionNode := getNoneRingIntersectionNode(head1, head2)
		if intersectionNode != nil { //环外相交
			ringNode1.Next = tmpRing1
			ringNode2.Next = tmpRing2
			return intersectionNode
		}
		//恢复
		ringNode1.Next = tmpRing1
		ringNode2.Next = tmpRing2
		//环内相交，不相交
		for iterator := ringNode1; iterator != nil; {
			if iterator == ringNode2 { //相交
				return iterator
			}
			iterator = iterator.Next
			if iterator == ringNode1 { //转了一圈，跳出
				break
			}
		}
		//跳出循环，不相交
		return nil
	} else { //2. 无环：相交、不相交
		return getNoneRingIntersectionNode(head1, head2)
	}
}
func getNoneRingIntersectionNode(head1, head2 *Node) *Node {
	var (
		list1Count, list2Count                           int
		list1Iterator, list2Iterator, list1End, list2End *Node
	)

	for list1Iterator = head1; list1Iterator != nil; list1Iterator = list1Iterator.Next {
		list1Count++
		list1End = list1Iterator
	}
	for list2Iterator = head1; list2Iterator != nil; list2Iterator = list2Iterator.Next {
		list2Count++
		list2End = list2Iterator
	}
	if list1End != list2End { //不相交
		return nil
	}
	//相交
	var startNode1, startNode2 *Node
	if list1Count > list2Count {
		startNode1 = head1
		for i := 0; i < list1Count-list2Count; i++ {
			startNode1 = startNode1.Next
		}
		startNode2 = head2
	} else {
		startNode1 = head2
		for i := 0; i < list2Count-list1Count; i++ {
			startNode1 = startNode1.Next
		}
		startNode2 = head1
	}
	for iterator := startNode1; iterator != nil; iterator = iterator.Next {
		if iterator == startNode2 {
			return iterator
		}
		startNode2 = startNode2.Next
	}
	return nil
}
func TestGetListIntersectNode(t *testing.T) {
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

	n21 := &Node{Value: 21}
	n22 := &Node{Value: 22}
	n23 := &Node{Value: 23}
	n24 := &Node{Value: 24}
	n25 := &Node{Value: 25}
	n26 := &Node{Value: 26}
	n27 := &Node{Value: 27}

	n21.Next = n22
	n22.Next = n23
	n23.Next = n24
	n24.Next = n25
	n25.Next = n26
	n26.Next = n27

	n31 := &Node{Value: 31}
	n32 := &Node{Value: 32}
	n33 := &Node{Value: 33}
	n34 := &Node{Value: 34}
	n35 := &Node{Value: 35}

	n7.Next = n31
	n31.Next = n32
	n32.Next = n33
	n33.Next = n34
	n34.Next = n35
	n35.Next = n33

	n27.Next = n24

	fmt.Println(getListIntersectNode(n1, n21))
}
