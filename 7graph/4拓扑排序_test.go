/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2022/1/4 11:15 上午
 * @Desc:
 */

package _graph

import (
	"algorithm/utils/mqueue"
	"fmt"
	"testing"
)

//算法：
//1.在图中找到所有入度为0的点输出
//2.把所有入度为0的点在图中删掉，继续找入度为0的点，周而复始
//3.图的所有点都被删除后，依次输出的顺序就是拓扑排序

//要求: 有向图且其中没有环
//应用: 事件安排、编译顺序

//图的拓扑排序
//思路：
//辅助数据结构1.一个map[*node]int, key:节点， int:入度
//辅助数据结构2. 0入度队列，存放入读为0的点
//过程，依次处理0入度队列的节点，0入度节点的next节点入度-1
func topologicalSort(graph *Graph) []*node {
	if graph.Nodes == nil {
		return nil
	}
	var (
		zeroQueue   = mqueue.New()
		result      []*node
		inDegreeMap = make(map[*node]int)
	)
	//初始化入度map
	for _, n := range graph.Nodes {
		inDegreeMap[n] = n.In
		if n.In == 0 {
			zeroQueue.PushBack(n)
		}
	}
	for !zeroQueue.IsEmpty() {
		curNode := zeroQueue.PopFront().(*node)
		result = append(result, curNode)
		for next := curNode.NextNodes.PopFront(); next != nil; next = curNode.NextNodes.PopFront() {
			nextNode := next.(*node)
			nextNode.In--
			if nextNode.In == 0 {
				zeroQueue.PushBack(nextNode)
			}
			inDegreeMap[nextNode] = nextNode.In
		}
	}
	return result
}

func TestTopologicalSort(t *testing.T) {
	InitGraph()
	result := topologicalSort(g)
	for _, n := range result {
		fmt.Println(n.Value)
	}
}
