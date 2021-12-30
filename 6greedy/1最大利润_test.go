/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/22 1:42 下午
 * @Desc:
 */

package _greedy

import (
	"fmt"
	"testing"
)

//贪心算法
//题目
//输入： 正数数组costs、正数数组profits、正数K、正数M
//costs[i]表示i号项目的花费，profits[i]表示i号项目在扣除花费之后还能挣到的钱
//K表示你只能串行的最多做K个项目
//M表示你初始的资金
//说明：每做完一个项目，马上获得的收益，可以支持你去做下一个项目。不能并行的做项目。
//输出：你最后获得的最大钱数
var (
	K       int = 2 //K表示你只能串行的最多做K个项目
	M       int = 1 //表示你初始的资金
	costs       = []int{6, 7, 9, 1, 1, 4, 5}
	profits     = []int{2, 3, 5, 6, 3, 4, 6}
)

func TestBestChoice(t *testing.T) {
	var (
		c            CostCal
		p            ProfitsCal
		downProjects int
		gain         = M
	)
	c.InitProject(costs, profits)

	for project := c.GetLeastCostProject(); project != nil; project = c.GetLeastCostProject() {
		if downProjects >= K {
			break
		}
		for project.Cost <= gain {
			p.PutProject(project)
			project = c.GetLeastCostProject()
			continue
		}

		for todo := p.GetProject(); todo != nil; {
			gain += todo.Profit
			downProjects += 1
			if project.Cost <= gain {
				p.PutProject(project)
				break
			} else if downProjects < K {
				todo = p.GetProject()
			} else {
				break
			}
		}
	}
	fmt.Println("profit: ", gain, ", K: ", K)
}
