package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
BM39 序列化二叉树
https://www.nowcoder.com/practice/cf7e25aa97c04cc1a68c8f040e71fb84?tpId=295&tqId=23455&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
请实现两个函数，分别用来序列化和反序列化二叉树，不对序列化之后的字符串进行约束，但要求能够根据序列化之后的字符串重新构造出一棵与原二叉树相同的树。

二叉树的序列化(Serialize)是指：把一棵二叉树按照某种遍历方式的结果以某种格式保存为字符串，从而使得内存中建立起来的二叉树可以持久保存。序列化可以基于先序、中序、后序、层序的二叉树等遍历方式来进行修改，序列化的结果是一个字符串，序列化时通过 某种符号表示空节点（#）
二叉树的反序列化(Deserialize)是指：根据某种遍历顺序得到的序列化字符串结果str，重构二叉树。

例如，可以根据层序遍历的方案序列化，如下图:
层序序列化(即用函数Serialize转化)如上的二叉树转为"{1,2,3,#,#,6,7}"，再能够调用反序列化(Deserialize)将"{1,2,3,#,#,6,7}"构造成如上的二叉树。
当然你也可以根据满二叉树结点位置的标号规律来序列化，还可以根据先序遍历和中序遍历的结果来序列化。不对序列化之后的字符串进行约束，所以欢迎各种奇思妙想。


示例1
输入：
{1,2,3,#,#,6,7}
返回值：
{1,2,3,#,#,6,7}


示例2
输入：
{8,6,10,5,7,9,11}
返回值：
{8,6,10,5,7,9,11}

*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func Serialize(root *TreeNode) string {
	// write code here
	var ret []string
	var fn = func(val string) {
		ret = append(ret, val)
	}
	serialize(root, fn)
	return strings.Join(ret, ",")
}

func serialize(root *TreeNode, fn func(val string)) {
	if root == nil {
		fn("#")
		return
	}
	fn(strconv.Itoa(root.Val))
	serialize(root.Left, fn)
	serialize(root.Right, fn)
}

func Deserialize(s string) *TreeNode {
	// write code here
	var str = strings.Split(s, ",")
	var fn func() *TreeNode
	fn = func() *TreeNode {
		if str[0] == "#" {
			str = str[1:]
			return nil
		}
		valInt, _ := strconv.Atoi(str[0])
		str = str[1:]
		return &TreeNode{
			Val:   valInt,
			Left:  fn(),
			Right: fn(),
		}
	}
	return fn()
}

func main16() {
	node := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 5,
		},
	}

	result1 := Serialize(&node)
	fmt.Printf("%+v\n", result1)

	deserialize := Deserialize(result1)
	fmt.Printf("%+v", deserialize)
}
