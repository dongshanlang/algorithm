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
	"net/http"
	"sync"
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
func partition(arr []int, pivot int) (leftEnd, rightBegin int) {
	pivotValue := arr[pivot]
	fmt.Println(pivot, pivotValue)
	arr[pivot], arr[len(arr)-1] = arr[len(arr)-1], arr[pivot]
	fmt.Println(arr)
	left := 0
	right := len(arr) - 2
	pointer := 0
	for pointer <= right {
		fmt.Println(left, right, pointer)
		if arr[pointer] < pivotValue {
			arr[pointer], arr[left] = arr[left], arr[pointer]
			left++
			pointer++
		} else if arr[pointer] > pivotValue {
			arr[pointer], arr[right] = arr[right], arr[pointer]
			right--
		} else {
			pointer++
		}
		fmt.Println(arr)
	}

	arr[right+1], arr[len(arr)-1] = arr[len(arr)-1], arr[right+1]
	return left, right + 1

}
func quickSort(arr []int, left, right int) {
	if right-left < 1 {
		return
	}
	pivot := left + int(rand.Int31n(int32(right-left)))
	leftEnd, rightBegin := partition(arr, pivot)
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
	fmt.Println(partition(arr, 2))
	fmt.Println(arr)
}

func TestList(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			for {
				_, err := http.Get("http://dt1.8tupian.com/4614a104b880.pg0?8_t=13")
				if err != nil {
					t.Log(err)
				}
			}
		}()
	}
	wg.Wait()
	resp, err := http.Get("http://dt1.8tupian.com/4614a104b880.pg0?8_t=13")
	if err != nil {
		t.Log(err)
	}
	t.Log(resp.Status)
}
