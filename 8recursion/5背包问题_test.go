/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/1/12 4:01 下午
 * @Desc:
 */

package trecursion

import (
	"fmt"
	"math"
	"testing"
)

//给定两个长度都为N的数组weights和values
//weights[i]和values[i]分表代表i号货物的重量和价值
//给定一个正数bag，表示一个载重bag的袋子
//你装的物品不能超过这个重量
//返回你能装下最多的价值是多少？
type Goods struct {
	Weights []int
	Values  []int
}
type Bag struct {
	Goods
	Bag int
}

func (b *Bag) Give(g *Goods) {
	b.Goods = *g
}
func (b *Bag) process(index, currentWeight int) int {
	if b.Bag < currentWeight {
		return -1
	}
	if index == len(b.Weights) {
		return 0
	}
	//包含当前index的价值
	includeValue := b.process(index+1, currentWeight+b.Weights[index])
	if includeValue != -1 {
		includeValue = includeValue + b.Values[index]
	}
	//不包含当前index的value
	notIncludeValue := b.process(index+1, currentWeight)
	return int(math.Max(float64(notIncludeValue), float64(includeValue)))
}
func (b *Bag) GetMostValue() int {
	return b.process(0, 0)
}
func TestBag(t *testing.T) {
	b := Bag{
		Goods: Goods{
			Weights: []int{1, 2, 3, 4, 5},
			Values:  []int{1, 2, 4, 4, 5},
		},
		Bag: 3,
	}
	fmt.Println(b.GetMostValue())
}
