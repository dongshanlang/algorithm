/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/10 4:34 下午
 * @Desc:
 */

package _3list

//将单向链表按照某值划分成左边小，中间相等，右边大的形式
//1。把链表放入数组里，在数组上做partition
//2。分成小中大三部分，再把各个部分串起来
//   设置6个指针，小区headPointer, tailPointer, ==:headPointer, tailPointer, > :headPointer, tailPointer
//	 串联六个区，需要考虑各个区域可能为空的情况
