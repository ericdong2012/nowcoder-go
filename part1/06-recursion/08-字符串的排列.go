package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
NC121 字符串的排列
https://www.nowcoder.com/practice/fe6b651b66ae47d7acce78ffdd9a96c7?tpId=117&tqId=37778&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=591&title=


描述
输入一个长度为 n 字符串，打印出该字符串中字符的所有排列，你可以以任意顺序返回这个字符串数组。
例如输入字符串ABC,则输出由字符A,B,C所能排列出来的所有字符串ABC,ACB,BAC,BCA,CBA和CAB。

数据范围：n < 10
要求：空间复杂度 O(n!)，时间复杂度 O(n!)
输入描述：
输入一个字符串,长度不超过10,字符只包括大小写字母。

示例1
输入：
"ab"
返回值：
["ab","ba"]

说明：
返回["ba","ab"]也是正确的


示例2
输入：
"aab"
返回值：
["aab","aba","baa"]

示例3
输入：
"abc"
返回值：
["abc","acb","bac","bca","cab","cba"]

示例4
输入：
""
返回值：
[]

*/

/*
使用临时变量去组装一个排列的情况：每当我们选取一个字符以后，就确定了其位置，相当于对字符串中剩下的元素进行全排列添加在该元素后面，给剩余部分进行全排列就是一个子问题，因此可以使用递归。

终止条件： 临时字符串中选取了n个元素，已经形成了一种排列情况了，可以将其加入输出数组中。
返回值： 每一层给上一层返回的就是本层级在临时字符串中添加的元素，递归到末尾的时候就能添加全部元素。
本级任务： 每一级都需要选择一个元素加入到临时字符串末尾（遍历原字符串选择）。
递归过程也需要回溯，比如说对于字符串“abbc”，如果事先在临时字符串中加入了a，后续子问题只能是"bbc"的全排列接在a后面，对于b开头的分支达不到，因此也需要回溯：将临时字符串刚刚加入的字符去掉，同时vis修改为没有加入，这样才能正常进入别的分支。


官方解法：
对于这个排列，我们是固定A不动，然后交换B与C，从而得到"ABC" 和 "ACB"
同理，对于"BAC"、"BCA" 、"CAB"和"CBA"是同样道理

递归三部曲：
	递归函数的功能：dfs(int pos, string s), 表示固定字符串s的pos下标的字符s[pos]
	递归终止条件：当pos+1 == s.length()的时候，终止，表示对最后一个字符进行固定，也就说明，完成了一次全排列
	下一次递归：dfs(pos+1, s), 很显然，下一次递归就是对字符串的下一个下标进行固定
	但是，对于"ABB"来说，就会有重复，如图

所以，我们用set可以进行去重，并且可以达到按字母顺序排序。
*/

func Permutation(str string) []string {
	if str == "" {
		return nil
	}
	// 构建slice并排序
	strList := make([]string, 0, len(str))
	for _, v := range str{
		strList = append(strList, string(v))
	}
	sort.Strings(strList)
	// dfs
	result := []string{}
	dfs(strList, 0, &result)

	return result
}

func dfs(strs []string, index int, result *[]string) {
	// 排列到最后一位，无需再处理，直接生成字符序列
	if index == len(strs)-1 {
		newStr := strings.Join(strs, "")
		*result = append(*result, newStr)
		return
	}

	for i := index; i < len(strs); i++ {
		//// 剪枝, 下一次循环的，字符相同
		if i != index && strs[i] == strs[index] {
			continue
		}
		// 递归处理下一层
		strs[i], strs[index] = strs[index], strs[i]
		dfs(strs, index+1, result)
	}
	for i := len(strs) - 1; i > index; i-- {
		strs[i], strs[index] = strs[index], strs[i]
	}
}

func main08() {
	permutation := Permutation("abc")
	fmt.Println(permutation)
}
