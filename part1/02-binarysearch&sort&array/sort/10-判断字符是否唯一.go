package main

import "fmt"

/*
NC229 判断字符是否唯一
https://www.nowcoder.com/practice/fb08c63e2c4f4f7dbf41086e46044211?tpId=117&tqId=39454&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590&title=

描述
给定一个字符串，请你判断其中每个字符是否全都不同。
数据范围：字符串长度满足 1 \le n \le 100 \1≤n≤100


示例1
输入：
"nowcoder"
返回值：
false

说明：
 "nowcoder" 中 'o' 出现了两次，因此返回 false


示例2
输入：
"nowcOder"
返回值：
true

说明：
每个字符都只出现了一次
*/

func isUnique(str string) bool {
	// write code here
	hash := make(map[string]int, 0)
	for i := 0; i < len(str); i++ {
		if _, ok := hash[string(str[i])]; ok {
			return false
		}
		hash[string(str[i])] += 1
	}

	return true
}

func main() {
	unique := isUnique("nowcoder")
	fmt.Println(unique)
}
