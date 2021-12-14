/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/13 5:32 下午
 * @Desc:
 */

package mstack

import "container/list"

type Stack struct {
	list *list.List
}

func New() *Stack {
	return &Stack{list: list.New()}
}
func (s *Stack) Push(item interface{}) {
	if s.list == nil {
		s.list = list.New()
	}
	s.list.PushBack(item)
}
func (s *Stack) Pop() interface{} {
	var item interface{}
	if s.list.Back() != nil {
		item = s.list.Back().Value
	}
	if s.list.Back() != nil {
		s.list.Remove(s.list.Back())
	}
	return item
}
func (s *Stack) Peek() interface{} {
	if s.list.Back() != nil {
		return s.list.Back().Value
	}
	return nil
}
func (s *Stack) Len() int {
	return s.list.Len()
}
func (s *Stack) IsEmpty() bool {
	return s.list.Len() == 0
}
