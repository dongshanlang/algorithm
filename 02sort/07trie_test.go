/**
 * @Author: lixiumin
 * @E-Mail: lixiuminmxl@163.com
 * @Date: 2021/12/9 4:16 下午
 * @Desc:
 */

package tsort

import (
	"fmt"
	"testing"
)

//trie 前缀树
//使用前缀数查询字符串出现的次数
type Node struct {
	Pass  int
	End   int
	Paths [26]*Node
}
type Trie struct {
	Root *Node
}

//查询某字符串是否存在
func (t *Trie) Search(str string) bool {
	searchNode := t.Root
	if searchNode == nil {
		return false
	}
	for _, ch := range str {
		if searchNode.Paths[ch-'a'] == nil {
			return false
		}
		searchNode = searchNode.Paths[ch-'a']
	}
	return searchNode.End > 0
}

//查询多少字符串包含指定字串
func (t *Trie) GetCount(str string) int {
	if t.Root == nil {
		return 0
	}
	countNode := t.Root
	for _, ch := range str {
		if countNode.Paths[ch-'a'] == nil {
			return 0
		}
		countNode = countNode.Paths[ch-'a']
	}
	return countNode.Pass
}
func (t *Trie) Push(str string) {
	if t.Root == nil {
		t.Root = &Node{
			Pass:  0,
			End:   0,
			Paths: [26]*Node{},
		}
	}
	insertNode := t.Root
	for _, ch := range str {
		if insertNode.Paths[ch-'a'] == nil {
			insertNode.Paths[ch-'a'] = &Node{
				Pass:  1,
				End:   0,
				Paths: [26]*Node{},
			}
		} else {
			insertNode.Paths[ch-'a'].Pass += 1
		}
		insertNode = insertNode.Paths[ch-'a']
	}
	insertNode.End += 1
}
func (t *Trie) Remove(str string) {
	if !t.Search(str) || t.Root == nil {
		return
	}
	removeNode := t.Root
	for _, ch := range str {
		removeNode.Paths[ch-'a'].Pass--
		removeNode = removeNode.Paths[ch-'a']
	}
	removeNode.End--
}
func TestPushTrie(t *testing.T) {
	var tr Trie
	tr.Push("abc")
	tr.Push("abd")
	tr.Push("abcd")
	tr.Push("abcf")
	tr.Push("abcsd")
	tr.Push("abcd")
	tr.Push("abca")

	fmt.Println("search abc: ", tr.Search("abc"))
	fmt.Println("get count abc", tr.GetCount("abc"))
	tr.Remove("abc")
	fmt.Println("search abc: ", tr.Search("abc"))
	fmt.Println("get count abc", tr.GetCount("abc"))
	fmt.Println(tr)
}
