/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/10 4:34 下午
 * @Desc:
 */

package _3list

import (
	"fmt"
	"testing"
)

//将单向链表按照某值划分成左边小，中间相等，右边大的形式
//1。把链表放入数组里，在数组上做partition
//2。分成小中大三部分，再把各个部分串起来
//   设置6个指针，小区headPointer, tailPointer, ==:headPointer, tailPointer, > :headPointer, tailPointer
//	 串联六个区，需要考虑各个区域可能为空的情况
func partition(head *Node, partitionValue int) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		smallHeadPointer, smallTailPointer, equalHeadPointer,
		equalTailPointer, bigHeadPointer, bigTailPointer, rangePointer,
		retHead *Node
	)
	rangePointer = head
	for rangePointer != nil {
		if rangePointer.Value < partitionValue {
			if smallHeadPointer == nil || smallTailPointer == nil {
				smallHeadPointer = rangePointer
				smallTailPointer = rangePointer
			} else {
				smallTailPointer.Next = rangePointer
				smallTailPointer = rangePointer
			}
		} else if rangePointer.Value == partitionValue {
			if equalHeadPointer == nil || equalTailPointer == nil {
				equalHeadPointer = rangePointer
				equalTailPointer = rangePointer
			} else {
				equalTailPointer.Next = rangePointer
				equalTailPointer = rangePointer
			}
		} else {
			if bigHeadPointer == nil || bigTailPointer == nil {
				bigHeadPointer = rangePointer
				bigTailPointer = rangePointer
			} else {
				bigTailPointer.Next = rangePointer
				bigTailPointer = rangePointer
			}
		}
		rangePointer = rangePointer.Next
	}

	//处理头尾的问题
	if smallHeadPointer == nil || smallTailPointer == nil { //小区空
		if equalHeadPointer == nil || equalTailPointer == nil { //等于区为空
			if bigHeadPointer == nil || bigTailPointer == nil {
				panic("error occur, nil list")
			} else {
				retHead = bigHeadPointer
				bigTailPointer.Next = nil
			}
		} else {
			//等于区不为空
			retHead = equalHeadPointer
			if bigHeadPointer == nil || bigTailPointer == nil { //大于区为空
				equalTailPointer.Next = nil
			} else {
				equalTailPointer.Next = bigHeadPointer
				bigTailPointer.Next = nil
			}
		}
	} else { //小区有值
		retHead = smallHeadPointer
		if equalHeadPointer == nil || equalTailPointer == nil {
			if bigHeadPointer == nil || bigTailPointer == nil {
				smallTailPointer.Next = nil
			} else {
				smallTailPointer.Next = bigHeadPointer
				bigTailPointer.Next = nil
			}
		} else {
			smallTailPointer.Next = equalHeadPointer
			if bigHeadPointer == nil || bigTailPointer == nil {
				equalTailPointer.Next = nil
			} else {
				equalTailPointer.Next = bigHeadPointer
				bigTailPointer.Next = nil
			}
		}

	}
	return retHead
}
func TestPartition(t *testing.T) {
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
	newHead := partition(n1, 13)
	fmt.Println(newHead)
}
