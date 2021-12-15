/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/14 7:09 下午
 * @Desc:
 */

package mqueue

import "container/list"

type MQueue struct {
	list *list.List
}

func New() *MQueue {
	return &MQueue{list: list.New()}
}
func (l *MQueue) PushBack(item interface{}) {
	l.list.PushBack(item)
}
func (l *MQueue) Front() interface{} {
	if l.list.Front() != nil {
		return l.list.Front().Value
	}
	return nil
}
func (l *MQueue) PopFront() interface{} {
	if l.list.Front() != nil {
		item := l.list.Front()
		l.list.Remove(item)
		return item.Value
	}
	return nil
}
func (l *MQueue) Len() int {
	return l.list.Len()
}
func (l *MQueue) IsEmpty() bool {
	return l.list.Len() == 0
}
