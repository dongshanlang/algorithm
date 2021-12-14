/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/14 7:09 下午
 * @Desc:
 */

package mlist

import "container/list"

type MList struct {
	list *list.List
}

func New() *MList {
	return &MList{list: list.New()}
}
func (l *MList) PushBack(item interface{}) {
	l.list.PushBack(item)
}
func (l *MList) Front() interface{} {
	if l.list.Front() != nil {
		return l.list.Front().Value
	}
	return nil
}
func (l *MList) PopFront() interface{} {
	if l.list.Front() != nil {
		item := l.list.Front()
		l.list.Remove(item)
		return item.Value
	}
	return nil
}
func (l *MList) Len() int {
	return l.list.Len()
}
func (l *MList) IsEmpty() bool {
	return l.list.Len() == 0
}
