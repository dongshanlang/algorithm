/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/3/3 6:11 下午
 * @Desc:
 */

package _golang

import (
	"fmt"
	"testing"
)

type Car struct {
	Owner string
	Speed int
}

func (c Car) ChangeOwner(owner string) {
	c.Owner = owner
}
func (c Car) ChangeSpeed(speed int) {
	c.Speed = speed
}
func (c *Car) ChangeOwnerP(owner string) {
	c.Owner = owner
}
func (c *Car) ChangeSpeedP(speed int) {
	c.Speed = speed
}

func TestChange(t *testing.T) {
	var c = Car{
		Owner: "a",
		Speed: 1,
	}
	fmt.Println(c)
	c.ChangeOwner("b")
	c.ChangeSpeed(2)
	fmt.Println(c)
	c.ChangeOwnerP("c")
	c.ChangeSpeedP(3)
	fmt.Println(c)
}
func TestOddNumber(t *testing.T) {
	arr := []int{1, 2, 3, 4, 3, 2, 1}
	number := findOddNumber(arr)
	t.Log(number)
}
func findOddNumber(arr []int) int {
	if len(arr) < 0 {
		return -1
	}
	result := 0
	for _, v := range arr {
		result = result ^ v
	}
	return result
}

func TestRightOne(t *testing.T) {
	num := 6
	number := getRightOne(num)
	t.Log(number)
	t.Log(^number)
}
func getRightOne(number int) int {
	return number & (^number + 1)
}

func TestGetTwoOddTimesNumber(t *testing.T) {
	arr := []int{1, 2, 3, 4, 3, 2, 1, 5}
	t.Log(getTwoOddTimesNumber(arr))
}
func getTwoOddTimesNumber(arr []int) (int, int) {
	length := len(arr)
	if length < 2 {
		return 0, 0
	}
	exclusiveOrAll := 0
	for i := 0; i < length; i++ {
		exclusiveOrAll = exclusiveOrAll ^ arr[i]
	}
	//get left one
	rightOne := exclusiveOrAll & (^exclusiveOrAll + 1)

	var number1, number2 int
	for i := 0; i < length; i++ {
		if arr[i]&rightOne == 0 {
			number1 = number1 ^ arr[i]
		} else {
			number2 = number2 ^ arr[i]
		}
	}
	return number1, number2
}
