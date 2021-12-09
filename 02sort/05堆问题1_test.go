/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/8 6:20 下午
 * @Desc:
 */

package _2sort

import (
	"fmt"
	"testing"
)

//已知一个几乎有序的数组。几乎有序是指，如果把数组排好顺序的话，每个元素移动的距离一定不超过k，并且k相对于数组长度来说是比较小的。
// 请选择一个合适的排序策略，对这个数组进行排序。

//思路，构建一个含有K个数的小根堆，取出最小数，放在排序位置，再加入未排序数，
func TestSortArray(t *testing.T) {
	arr := []int{1, 3, 2, 5, 4, 7, 6, 9, 11, 10}
	k := 3
	Sort05(arr, k)
	fmt.Println(arr)

}
func Sort05(arr []int, biggestMoveDistance int) {
	if len(arr) <= 1 || biggestMoveDistance < 2 {
		return
	}
	var helpArrLength = 0
	if biggestMoveDistance < len(arr) {
		helpArrLength = biggestMoveDistance
	} else {
		helpArrLength = len(arr)
	}
	var helpArr []int
	for i := 0; i < helpArrLength; i++ {
		helpArr = append(helpArr, arr[i])
		InsertHeap05(helpArr, i)
	}
	if len(helpArr) != 0 {
		for sorted := 0; sorted < len(arr)-biggestMoveDistance; sorted++ {
			arr[sorted] = helpArr[0]
			helpArr[0] = arr[sorted+biggestMoveDistance]
			Heapify05(helpArr, 0, len(helpArr))
		}
	}
	fmt.Println(arr)
	if len(helpArr) > 1 {
		for i := 0; i < len(helpArr); i++ {
			helpArr[0], helpArr[len(helpArr)-i-1] = helpArr[len(helpArr)-i-1], helpArr[0]
			Heapify05(helpArr, 0, len(helpArr)-i-1)
		}
		for i := 0; i < len(helpArr); i++ {
			arr[len(arr)-1-i] = helpArr[i]
		}
	}
	fmt.Println(arr)
	fmt.Println(helpArr)
}
func Heapify05(arr []int, index, heapSize int) {
	if heapSize < 1 || len(arr) <= 1 || index >= heapSize {
		return
	}
	left := 2*index + 1
	for left < heapSize {
		smallest := index
		if left+1 < heapSize {
			if arr[left+1] < arr[left] {
				smallest = left + 1
			} else {
				smallest = left
			}
		}
		if arr[index] <= arr[smallest] {
			smallest = index
		}
		if smallest == index {
			break
		}
		arr[index], arr[smallest] = arr[smallest], arr[index]
		index = smallest
		left = 2*index + 1
	}
}

//
func InsertHeap05(arr []int, index int) {
	if index <= 0 {
		return
	}
	parent := (index - 1) / 2
	for parent >= 0 {
		if arr[parent] > arr[index] {
			arr[parent], arr[index] = arr[index], arr[parent]
			index = parent
			parent = (index - 1) / 2
		} else {
			break
		}
	}
}
