/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/31 2:24 下午
 * @Desc:
 */

package _graph

import (
	"algorithm/utils/mqueue"
)

type Graph struct {
	Nodes map[int]*node
	Edges []*edge
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: map[int]*node{},
		Edges: nil,
	}
}

type node struct {
	ID        int
	Value     int
	In, Out   int
	NextNodes *mqueue.MQueue
	Edges     *mqueue.MQueue
}

func NewNode(value int) *node {
	return &node{
		ID:        0,
		Value:     value,
		In:        0,
		Out:       0,
		NextNodes: mqueue.New(),
		Edges:     mqueue.New(),
	}
}

type edge struct {
	Value    int
	From, To *node
}

func NewEdge(value int, from, to *node) *edge {
	return &edge{
		Value: value,
		From:  from,
		To:    to,
	}
}
