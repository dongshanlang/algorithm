/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/8 6:06 下午
 * @Desc:
 */

package _2sort

import (
	"fmt"
	"testing"
)

//1.先让整个数组变成大根堆结构，建立堆的过程：
//  1）从上到下的方法，时间复杂度O(N*logN)
//  2)从下到上的方法，时间复杂度O(N)
//2.把堆的最大值和堆末尾的值交换，然后减少堆的大小之后，再去调整堆，一直周而复始，时间复杂度为O(N*logN)
//3.堆的大小减为0后，排序完成

//从根节点往下算
//从index位置往下看，到我的孩子都比我小或者没有孩子停
func Heapify(arr []int, index, heapSize int) {
	if index < 0 || index > len(arr)-1 || index >= heapSize {
		return
	}
	left := index*2 + 1
	largest := 0
	for left < heapSize {
		if arr[left] > arr[index] {
			largest = left
		}
		if left+1 < heapSize && arr[left+1] > arr[largest] {
			largest = left + 1
		}
		if largest == index {
			break
		}
		arr[largest], arr[index] = arr[index], arr[largest]
		index = largest
		left = index*2 + 1
	}
}

func TestHeapSort(t *testing.T) {
	arr := []int{3, 2, 5, 7, 9, 1, 4, 8, 6, 5, 3, 7}
	fmt.Println(arr)
	for index := 1; index < len(arr); index++ {
		HeapInsert(arr, index)
	}
	fmt.Println(arr)
	for n := 0; n < len(arr); n++ {
		arr[0], arr[len(arr)-1-n] = arr[len(arr)-1-n], arr[0]
		Heapify(arr, 0, len(arr)-n-1)
		fmt.Println(arr)
	}
	fmt.Println(arr)
}
func TestHeapSort2(t *testing.T) {
	arr := []int{3, 2, 5, 7, 9, 1, 4, 8, 6, 5, 3, 7}
	fmt.Println(arr)
	for index := len(arr) / 2; index >= 0; index-- {
		Heapify(arr, index, len(arr))
	}
	fmt.Println(arr)
	for n := 0; n < len(arr); n++ {
		arr[0], arr[len(arr)-1-n] = arr[len(arr)-1-n], arr[0]
		Heapify(arr, 0, len(arr)-n-1)
		fmt.Println(arr)
	}
	fmt.Println(arr)
}

//0,1,2,3,4
//从最后往上算
//数组构建大根堆、小根堆结构
func HeapInsert(arr []int, index int) {
	if index <= 0 || index > len(arr)-1 {
		return
	}
	var parent = (index - 1) / 2
	for parent >= 0 {
		if arr[index] > arr[parent] { //大于交换，大顶堆
			arr[index], arr[parent] = arr[parent], arr[index]
			index = parent
			parent = (index - 1) / 2
		} else {
			break
		}
	}
}
func TestBuildHeap(t *testing.T) {
	arr := []int{3, 2, 5, 7, 9, 1, 4, 8, 6, 4, 6}
	fmt.Println(arr)
	for index := 1; index < len(arr); index++ {
		HeapInsert(arr, index)
	}
	fmt.Println(arr)
}
