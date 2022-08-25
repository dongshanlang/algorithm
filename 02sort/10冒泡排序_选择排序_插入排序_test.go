/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/8/25 11:03 AM
 * @Desc:
 */

package tsort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 9, 7, 4, 5, 6, 3, 2, 8}
	fmt.Println(arr)
	bubbleSort(arr)
	fmt.Println(arr)

}
func bubbleSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				continue
			}
		}
	}
	return
}

func TestSelectSort(t *testing.T) {
	arr := []int{1, 9, 7, 4, 5, 6, 3, 2, 8}
	fmt.Println(arr)
	selectSort(arr)
	fmt.Println(arr)

}
func selectSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	length := len(arr)
	var minIndex int
	for i := 0; i < length; i++ {
		minIndex = i
		for j := i; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func TestInsertSort(t *testing.T) {
	arr := []int{1, 9, 7, 4, 5, 6, 3, 2, 8}
	fmt.Println(arr)
	insertSort(arr)
	fmt.Println(arr)
}
func insertSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	for i := 1; i < length; i++ {
		for j := i; j >= 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				continue
			} else {
				break
			}
		}
	}
	return
}
