package main

import "fmt"

/*
NC52 有效括号序列
https://www.nowcoder.com/practice/37548e94a270412c8b9fb85643c8ccc2?tpId=117&tqId=37749&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=


描述
给出一个仅包含字符'(',')','{','}','['和']',的字符串，判断给出的字符串是否是合法的括号序列
括号必须以正确的顺序关闭，"()"和"()[]{}"都是合法的括号序列，但"(]"和"([)]"不合法。

要求：空间复杂度 O(n)，时间复杂度 O(n)
示例1
输入：
"()[]{}"
返回值：
true

示例2
输入：
"[]"
返回值：
true

示例3
输入：
"([)]"
返回值：
false

*/
func isValid(s string) bool {
	col := make([]rune, 0)

	for _, c := range s {
		switch c {
		// 以下case 只压栈
		case '(':
			col = append(col, ')')
		case '[':
			col = append(col, ']')
		case '{':
			col = append(col, '}')
		default:
			// s 是个“]” 等特殊情况
			if len(col) == 0 {
				return false
			}
			// 当前字符跟col中最后一个做比较，相等则弹栈， 否则false
			if rune(c) == col[len(col)-1] {
				col = col[:len(col)-1]
			} else {
				return false
			}
		}
	}

	// 全部弹栈完成，如果为空，说明是符合要求的
	return len(col) == 0
}

func main() {
	//valid := isValid("()[]{}")
	//valid := isValid("[()]{}")
	valid := isValid("[([[)]{}")
	//valid := isValid("[{{}}]")
	fmt.Printf("%+v", valid)
}
