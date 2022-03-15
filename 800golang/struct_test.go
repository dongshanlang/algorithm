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
