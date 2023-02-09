package main

import "fmt"

/*
NC220 重复的DNA序列
https://www.nowcoder.com/practice/fe9099e5308042a8af2f7aabdb3719fe?tpId=117&tqId=39424&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5074&title=

描述
所有的 DNA 序列都是由 'A' , ‘C’ , 'G' , 'T' 字符串组成的，例如 'ACTGGGC'
请你实现一个函数找出所有的目标子串，目标连续子串的定义是，长度等于 10 ，且在 DNA 序列中出现次数超过 1 次的连续子串（允许两个连续子串有重合的部分，如下面的示例2所示）。
（注：返回的所有目标子串的顺序必须与原DNA序列的顺序一致，如下面的示例1所示）

示例1
输入：
"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
返回值：
["AAAAACCCCC","CCCCCAAAAA"]
说明：
"AAAAACCCCC"和"CCCCCAAAAA"长度等于 10 且在DNA序列中分别出现了 2 次。
不能返回["CCCCCAAAAA","AAAAACCCCC"]，因为在原DNA序列中，"AAAAACCCCC"要比"CCCCCAAAAA"先出现。

示例2
输入：
"AAAAAAAAAAA"
返回值：
["AAAAAAAAAA"]

*/

func repeatedDNA(DNA string) []string {
	// write code here
	// 长度是10
	/*
		用哈希表tmp储存子串和位置信息，遍历一遍字符串，统计每个子串出现的次数
		然后为了保证重复子串的顺序，再遍历一遍字符串，对于出现次数大于1的子串保存到res中，同时在哈希表中删除这一子串的次数信息，防止重复
	*/
	//储存子串和次数
	temp := make(map[string]int, len(DNA))
	for i := 0; i <= len(DNA)-10; i++ {
		// 长度为10
		str := DNA[i : i+10]
		temp[str]++
	}
	res := []string{}
	for i := 0; i <= len(DNA)-10; i++ {
		// 长度为10
		str := DNA[i : i+10]
		if temp[str] > 1 {
			res = append(res, str)
			delete(temp, str)
		}
	}
	return res
}

func main() {
	dna := repeatedDNA("AAAAACCCCCCCCCCCAAAAACCCCC")
	fmt.Println(dna)
}
