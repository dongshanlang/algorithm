/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/1/13 5:11 下午
 * @Desc:
 */

package tscale

import (
	"fmt"
	"math"
	"testing"
)

//给定一个整形数组arr，代表数值不同的纸牌排成一条线
//玩家A和玩家B依次拿走每张纸牌
//规定玩家A先拿走，玩家B后拿走，
//但是每个玩家每次只能拿走最左或者左右的纸牌，
//玩家A和玩家B都绝顶聪明。请返回最后获胜者的分数
type CardScore struct {
	arr []int
}

func (c *CardScore) first(left, right int) int {
	if left == right {
		return c.arr[left]
	}
	leftScore := c.arr[left] + c.second(left+1, right)
	rightScore := c.arr[right] + c.second(left, right-1)
	return int(math.Max(float64(leftScore), float64(rightScore)))
}
func (c *CardScore) second(left, right int) int {
	if left == right {
		return 0
	}
	//先选
	leftScore := c.arr[left] + c.second(left+1, right)
	rightScore := c.arr[right] + c.second(left, right-1)
	if leftScore > rightScore { //先手选left
		left++
	} else { //先手选right
		right--
	}
	return c.first(left, right)
}
func TestCardScore(t *testing.T) {
	sc := CardScore{
		arr: []int{1, 2, 4, 1, 9},
	}
	fmt.Println(sc.first(0, len(sc.arr)-1))
	fmt.Println(sc.second(0, len(sc.arr)-1))
}
