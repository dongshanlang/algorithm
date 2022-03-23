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

//描述
//请实现无重复数字的升序数组的二分查找
//给定一个 元素升序的、无重复数字的整型数组 nums 和一个目标值 target ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标（下标从 0 开始），否则返回 -1
//数据范围：0≤len(nums)≤2×10^5
//数组中任意值满足 ∣val∣≤10^9
//进阶：时间复杂度 O(logn) ，空间复杂度 O(1)

//示例1
//输入：
//[-1,0,3,4,6,10,13,14],13
//返回值：
//6
//13 出现在nums中并且下标为 6
//示例2
//输入：
//[],3

//返回值：
//-1

//说明：
//nums为空，返回-1
//示例3
//输入：
//[-1,0,3,4,6,10,13,14],2

//返回值：
//-1

//说明：
//2 不存在nums中因此返回 -1
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
