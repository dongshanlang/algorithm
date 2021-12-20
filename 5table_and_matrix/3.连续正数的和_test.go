/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/17 5:08 下午
 * @Desc:
 */

package _table_and_matrix

import (
	"fmt"
	"testing"
)

//定义一种数：可以表示成若干（数量>1）连续正数和的数
//比如：
//5=2+3,5就是这样的数
//12=3+4+5,12就是这样的数
//1不是这样的数，因为要求数量大于1个、连续正数和
//2=1+1,2也不是，因为等号右边不是连续正数
//给定一个参数，返回是不是可以表示成若干连续正数和的数
func getContinuousNumber(n int) bool {
	for i := 1; i <= n/2; i++ {
		sum := 0
		j := i
		for ; sum <= n; j++ {
			sum += j
			if sum == n {
				fmt.Println(i, " ", j)
				return true
			}
		}

	}
	return false
}
func TestGetContinuousNumber(t *testing.T) {
	fmt.Println(getContinuousNumber(12))
}
func TestGetContinuousNumberIn100(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(i, " ", getContinuousNumber(i))
	}
}
