/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/10 4:49 下午
 * @Desc:
 */

package _3链表

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
func TestSNode(t *testing.T) {
	sn1 := SpecialNode{}
	fmt.Println(sn1)
}
