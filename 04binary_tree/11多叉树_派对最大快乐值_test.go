/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/16 5:56 下午
 * @Desc:
 */

package ttree

import (
	"fmt"
	"math"
	"testing"
)

//求派对的最大快乐值
//员工的信息定义如下
type Employee struct {
	Happy        int
	SubOrdinates []*Employee
}

//这个公司要开Party，你可以决定哪些员工来，哪些员工不来，规则：
//1。如果某个员工来了，那么这个员工所有的直接下级都不能来
//2。派对的整体快乐值是所有到场员工的快乐值累加
//3。你的目标是让派对的整体快乐值尽量大
//给定一棵多叉树的头节点boss，请返回派对的最大快乐值
type HappyInfo struct {
	IncludeCurrentNodeHappyValue int
	ExcludeCurrentNodeHappyValue int
}

func GetHappyValue(e *Employee) HappyInfo {
	if e == nil {
		return HappyInfo{0, 0}
	}
	if len(e.SubOrdinates) == 0 {
		return HappyInfo{
			IncludeCurrentNodeHappyValue: e.Happy,
			ExcludeCurrentNodeHappyValue: 0,
		}
	}
	var includeHappyValue, excludeHappyValue int
	//包含当前节点
	//不包含
	for _, eNode := range e.SubOrdinates {
		info := GetHappyValue(eNode)
		includeHappyValue += GetHappyValue(eNode).ExcludeCurrentNodeHappyValue
		excludeHappyValue = excludeHappyValue + int(math.Max(float64(info.IncludeCurrentNodeHappyValue),
			float64(info.ExcludeCurrentNodeHappyValue)))
	}
	includeHappyValue += e.Happy
	return HappyInfo{
		IncludeCurrentNodeHappyValue: includeHappyValue,
		ExcludeCurrentNodeHappyValue: excludeHappyValue,
	}
}
func TestGetHappyValue(t *testing.T) {
	n1 := &Employee{Happy: 1}
	n2 := &Employee{Happy: 2}
	n3 := &Employee{Happy: 3}
	n4 := &Employee{Happy: 4}
	n5 := &Employee{Happy: 5}
	n6 := &Employee{Happy: 6}
	n7 := &Employee{Happy: 7}
	n8 := &Employee{Happy: 8}
	n9 := &Employee{Happy: 9}
	n1.SubOrdinates = append(n1.SubOrdinates, n2)
	n1.SubOrdinates = append(n1.SubOrdinates, n3)
	n1.SubOrdinates = append(n1.SubOrdinates, n4)
	n2.SubOrdinates = append(n2.SubOrdinates, n5)
	n2.SubOrdinates = append(n2.SubOrdinates, n6)

	n4.SubOrdinates = append(n4.SubOrdinates, n7)
	n4.SubOrdinates = append(n4.SubOrdinates, n8)
	n4.SubOrdinates = append(n4.SubOrdinates, n9)

	info := GetHappyValue(n1)
	fmt.Println(info)

}
