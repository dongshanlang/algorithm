/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/7 5:28 下午
 * @Desc:
 */

package _2sort

import (
	"fmt"
	"math/rand"
	"testing"
)

/*
选定3
5 6 3 4 8 2 1 9
p

9 6 3 4 8 2 1 5
p           R

1 6 3 4 8 2 9 5
p         R

1 6 3 4 8 2 9 5
  P       R

1 6 3 4 8 2 9 5
  P       R
1 2 3 4 8 2 9 5
  P       R
*/

//选定一个分区的数n,pointer指向数组第0个数，从第0个数开始, left, right, pointer
//pointer指向的当前数<n: pointer数与left数交换，pointer++
//>n: pointer数与 right交换
//==n: pointer++
func partition(arr []int, startIndex, stopIndex, pivotIndex int) (leftEnd, rightBegin int) {
	if startIndex >= stopIndex || startIndex >= len(arr) || stopIndex >= len(arr) {
		panic("error index")
	}
	leftEnd = startIndex
	rightBegin = stopIndex - 1
	pivotValue := arr[pivotIndex]
	arr[pivotIndex], arr[stopIndex] = arr[stopIndex], arr[pivotIndex]
	pointerIndex := startIndex
	for pointerIndex <= rightBegin {
		if arr[pointerIndex] < pivotValue {
			arr[pointerIndex], arr[leftEnd] = arr[leftEnd], arr[pointerIndex]
			pointerIndex++
			leftEnd++
		} else if arr[pointerIndex] > pivotValue {
			arr[pointerIndex], arr[rightBegin] = arr[rightBegin], arr[pointerIndex]
			rightBegin--
		} else {
			pointerIndex++
		}
	}
	arr[rightBegin+1], arr[stopIndex] = arr[stopIndex], arr[rightBegin+1]
	return leftEnd, rightBegin + 1
}
func quickSort(arr []int, left, right int) {
	if right-left < 1 {
		return
	}
	pivot := left + int(rand.Int31n(int32(right-left)))
	leftEnd, rightBegin := partition(arr, left, right, pivot)
	quickSort(arr, left, leftEnd-1)
	quickSort(arr, rightBegin+1, right)
}
func TestQuickSort(t *testing.T) {
	arr := []int{1, 2, 4, 6, 7, 2, 4, 36, 4, 1}
	fmt.Println(arr)
	quickSort(arr, 0, 9)
	fmt.Println(arr)
}
func TestPartition(t *testing.T) {
	arr := []int{1, 2, 4, 6, 7, 2, 4, 36, 4, 1}
	fmt.Println(arr)
	fmt.Println(partition(arr, 0, len(arr)-1, 2))
	fmt.Println(arr)
}
