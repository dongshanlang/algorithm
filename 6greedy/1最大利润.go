/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/30 10:34 上午
 * @Desc:
 */

package _greedy

import "math"

//贪心算法
//题目
//输入： 正数数组costs、正数数组profits、正数K、正数M
//costs[i]表示i号项目的花费，profits[i]表示i号项目在扣除花费之后还能挣到的钱
//K表示你只能串行的最多做K个项目
//M表示你初始的资金
//说明：每做完一个项目，马上获得的收益，可以支持你去做下一个项目。不能并行的做项目。
//输出：你最后获得的最大钱数

type ProjectConstruction struct {
	Cost   int
	Profit int
}

var (
	initialFound int = 1
	canDo            = 2
)

//分析：1。根据花费构建小顶堆，
//2。将花费小于初始资金的项目取出，根据收益构建大顶堆
//3。取出大顶堆的堆顶项目，第一个做，做完之后初始资金增加
//4。循环1-3，结束条件是无项目可做或者

func ComponentProjectConstruction(costs, profits []int) []*ProjectConstruction {
	var pc []*ProjectConstruction
	for index, v := range costs {
		pc = append(pc, &ProjectConstruction{
			Cost:   v,
			Profit: profits[index],
		})
	}
	return pc
}

func smallHeapify(datas []*ProjectConstruction, index, end int) {
	if len(datas) <= 1 {
		return
	}
	var leftIndex int
	var smallIndex = index
	for index >= 0 && index < end {
		smallIndex = index
		leftIndex = 2*index + 1
		if leftIndex <= end {
			if datas[leftIndex].Cost < datas[index].Cost {
				datas[leftIndex], datas[index] = datas[index], datas[leftIndex]
				smallIndex = leftIndex
			}
		}
		rightIndex := 2*index + 2
		if rightIndex <= end {
			if datas[rightIndex].Cost < datas[index].Cost {
				datas[rightIndex], datas[index] = datas[index], datas[rightIndex]
				smallIndex = rightIndex
			}
		}
		if index == smallIndex || index*2+1 > len(datas) {
			break
		}
		index = smallIndex
	}
}

type ProfitsCal struct {
	ProjectCanDo []*ProjectConstruction
}

func (p *ProfitsCal) PutProject(project *ProjectConstruction) {
	p.ProjectCanDo = append(p.ProjectCanDo, project)
	p.heapInsert(len(p.ProjectCanDo) - 1)
}
func (p *ProfitsCal) GetProject() *ProjectConstruction {
	if len(p.ProjectCanDo) <= 0 {
		return nil
	}
	mostProfitProject := p.ProjectCanDo[0]
	p.ProjectCanDo[0], p.ProjectCanDo[len(p.ProjectCanDo)-1] = p.ProjectCanDo[len(p.ProjectCanDo)-1], p.ProjectCanDo[0]
	if len(p.ProjectCanDo) > 1 {
		p.ProjectCanDo = p.ProjectCanDo[:len(p.ProjectCanDo)-1]
		p.heapify(0, len(p.ProjectCanDo)-1)
	} else {
		p.ProjectCanDo = nil
	}
	return mostProfitProject
}
func (p *ProfitsCal) heapify(index, end int) {
	if index < 0 || index > end {
		return
	}
	left := 2*index + 1
	biggestProfitIndex := index
	for left <= end {
		if p.ProjectCanDo[left].Profit > p.ProjectCanDo[biggestProfitIndex].Profit {
			biggestProfitIndex = left
		}
		right := left + 1
		if right <= end && p.ProjectCanDo[right].Profit > p.ProjectCanDo[biggestProfitIndex].Profit {
			biggestProfitIndex = right
		}
		if biggestProfitIndex == index {
			break
		}
		p.ProjectCanDo[biggestProfitIndex], p.ProjectCanDo[index] = p.ProjectCanDo[index], p.ProjectCanDo[biggestProfitIndex]
		index = biggestProfitIndex
		if index > math.MaxInt64/2 { //越界
			break
		}
		left = 2*index + 1
	}
}
func (p *ProfitsCal) heapInsert(index int) {
	if index <= 0 {
		return
	}
	parent := index / 2
	for parent >= 0 {
		if p.ProjectCanDo[parent].Profit < p.ProjectCanDo[index].Profit {
			p.ProjectCanDo[parent], p.ProjectCanDo[index] = p.ProjectCanDo[index], p.ProjectCanDo[parent]
			index = parent
		} else {
			break
		}
		parent = index / 2
	}
}

type CostCal struct {
	CostSmallHeap []*ProjectConstruction
}

func (c *CostCal) InitProject(costs, profits []int) {
	datas := ComponentProjectConstruction(costs, profits)
	end := len(datas)
	for i := len(datas) / 2; i >= 0; i-- {
		smallHeapify(datas, i, end-1)
	}
	c.CostSmallHeap = datas
}
func (c *CostCal) GetLeastCostProject() *ProjectConstruction {
	if len(c.CostSmallHeap) == 0 {
		return nil
	}
	retPC := c.CostSmallHeap[0]
	c.CostSmallHeap[0], c.CostSmallHeap[len(c.CostSmallHeap)-1] = c.CostSmallHeap[len(c.CostSmallHeap)-1], c.CostSmallHeap[0]
	c.CostSmallHeap = c.CostSmallHeap[0 : len(c.CostSmallHeap)-1]
	smallHeapify(c.CostSmallHeap, 0, len(c.CostSmallHeap)-1)
	return retPC
}
