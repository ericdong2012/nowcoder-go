package main

import (
	"fmt"
)

/*
NC124 字典树的实现
https://www.nowcoder.com/practice/a55a584bc0ca4a83a272680174be113b?tpId=117&tqId=37818&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=579&title=


描述
字典树又称为前缀树或者Trie树，是处理字符串常用的数据结构。

假设组成所有单词的字符仅是‘a’～‘z’，请实现字典树的结构，并包含以下四个主要的功能。

1. void insert(String word)：添加word，可重复添加；
2. void delete(String word)：删除word，如果word添加过多次，仅删除一次；
3. boolean search(String word)：查询word是否在字典树中出现过(完整的出现过，前缀式不算)；
4. int prefixNumber(String pre)：返回以字符串pre作为前缀的单词数量。

现在给定一个m，表示有m次操作，每次操作都为以上四种操作之一。每次操作会给定一个整数op和一个字符串word，op代表一个操作码，如果op为1，则代表添加word，op为2则代表删除word，op为3则代表查询word是否在字典树中，op为4代表返回以word为前缀的单词数量（数据保证不会删除不存在的word）。

对于每次操作，如果op为3时，如果word在字典树中，请输出“YES”，否则输出“NO”；如果op为4时，请输出返回以word为前缀的单词数量，其它情况不输出。
数据范围：操作数满足 0\le m \le 10^50≤m≤10
5
 ，字符串长度都满足 0 \le n \le 200≤n≤20
进阶：所有操作的时间复杂度都满足 O(n)

示例1
输入：
[["1","qwer"],["1","qwe"],["3","qwer"],["4","q"],["2","qwer"],["3","qwer"],["4","q"]]
返回值：
["YES","2","NO","1"]

*/

type TrieNode struct {
	Val byte                    //字符
	End bool                    //是否是某个单词的结尾
	Count int                   //插入字符串时，节点被访问的次数
	Child [26]*TrieNode         //子节点
}

func New(c byte) *TrieNode {
	return &TrieNode{Val:c}
}

//插入
func (t *TrieNode) Insert(word string) {
	curr := t
	n := len(word)
	var c byte = 'a'
	for i:=0;i<n;i++ {
		c = word[i]
		if curr.Child[c-'a'] == nil {
			curr.Child[c-'a'] = New(c)
		}
		curr = curr.Child[c-'a']
		curr.Count++
	}
	curr.End = true
}

//删除
func (t *TrieNode) Delete(word string) {
	curr := t
	n := len(word)
	var c byte = 'a'
	for i:=0;i<n;i++{
		c = word[i]
		curr = curr.Child[c-'a']
		curr.Count--
	}
	if curr.Count == 0 {
		curr.End = false
	}
}

func (t *TrieNode) Search(word string) bool {
	curr := t
	n := len(word)
	var c byte = 'a'
	for i:=0;i<n;i++ {
		c = word[i]
		if curr.Child[c-'a'] == nil {
			return false
		}
		curr = curr.Child[c-'a']
	}
	return curr.End
}

func (t *TrieNode) PrefixNumber(pre string) int {
	curr := t
	n := len(pre)
	var c byte = 'a'
	for i:=0;i<n;i++ {
		c = pre[i]
		if curr.Child[c-'a'] == nil {
			return 0
		}
		curr = curr.Child[c-'a']
	}
	return curr.Count
}


func trieU( operators [][]string ) []string {
	root := New('A')
	n := len(operators)
	res := make([]string,0)
	for i:=0;i<n;i++{
		if operators[i][0] == "1" {
			root.Insert(operators[i][1])
		}
		if operators[i][0] == "2" {
			root.Delete(operators[i][1])
		}
		if operators[i][0] == "3" {
			if root.Search(operators[i][1]) {
				res = append(res,"YES")
			} else {
				res = append(res,"NO")
			}
		}
		if operators[i][0] == "4" {
			count := root.PrefixNumber(operators[i][1])
			res = append(res, fmt.Sprintf("%d",count))
		}
	}
	return res
}