/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/17 3:58 下午
 * @Desc:
 */

package _table_and_matrix

import (
	"fmt"
	"testing"
)

//小虎区买苹果，商店只提供两种类型的塑料袋，每种类型都有任意数量。
//1）能装下6个苹果
//2）能装下8个苹果
//小虎可以自由的使用两种袋子来装苹果，但是小虎有强迫症，他要求自己使用的袋子必须最少，且使用的每个袋子必须装满。
//给定一个正整数N，返回至少使用多少个袋子，如果N无法让使用的每个袋子必须装满，返回-1

//6，8，12，18，24
func violentApple(n int) (six, eight int) {
	if n < 6 {
		return -1, -1
	}
	eight = n / 8
	eightRest := n % 8
	for eightRest < 24 && eight >= 0 {
		if eightRest == 0 {
			return eight, 0
		}
		if eightRest%6 == 0 {
			return eight, eightRest / 6
		} else {
			eight = eight - 1
			eightRest = n - eight*8
		}
	}
	return -1, -1
}
func TestViolentApple(t *testing.T) {
	for i := 0; i < 100; i++ {
		six, eight := violentApple(i)
		fmt.Printf("%d: %d %d %d\n", i, six, eight, six+eight)
	}
}
