/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/3/22 11:51 上午
 * @Desc:
 */

package _0binary_search

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	var arr = []int{-1, 0, 3, 4, 6, 10, 13, 14, 15}
	fmt.Println(search(arr, 1))
}
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	mid := len(nums) / 2
	start := 0
	end := len(nums) - 1
	for {
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
		mid = start + (end-start)/2
		if mid < 0 || mid > len(nums)-1 || end < start {
			break
		}
	}
	return -1
}
