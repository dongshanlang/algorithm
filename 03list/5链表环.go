/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/10 5:57 下午
 * @Desc:
 */

package _3list

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
