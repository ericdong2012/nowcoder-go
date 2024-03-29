package main

import (
	"fmt"
)

/*
NC219 移除K位数字
https://www.nowcoder.com/practice/0fe685c8272d40f1b9785fedd2499c1c?tpId=117&tqId=39423&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=


描述
给定一个以字符串表示的数字 num 和一个数字 k ，从 num 中移除 k 位数字，使得剩下的数字最小。如果可以删除全部数字则剩下0
1.num仅有数字组成
2.num是合法的数字，不含前导0
3.请你保证删除之后的num也不含前导0

示例1
输入：
"1432219",3
返回值：
"1219"

示例2
输入：
"10",1
返回值：
"0"

示例3
输入：
"100999",3
返回值：
"9"

*/

/*
class Solution:
    def removeKnums(self , num: str, k: int) -> str:
        # write code here
        stack = []
        for c in num:
            while stack and k and int(c) < int(stack[-1]):        # 维护单调栈的单调性
                k -= 1
                stack.pop()
            if not stack and c == '0':                            # 如果现在stack开头的全是0就直接跳过
                continue
            stack.append(c)
        stack = stack[:-k] if k else stack
        if not stack:                                             # 如果最后一个数也不剩下就说明应该结果是0
            stack.append('0')
        return ''.join(stack)
*/

// 单调栈
/*
我们需要维护一个单调栈，将数字字符串每一个元素进行入栈处理（在这里我们用列表表示单调栈）
当取到的数字字符大于栈顶的时候，我们让数字正常入栈
当取到的数字字符小于栈顶的时候，我们就将栈顶元素出栈，并且一直比较新的栈顶元素，直到栈顶元素与取得元素相等或更小为止，出栈则说明我们去掉这个数字字符，相应k的值就需要自减
*/
func removeKnums(num string, k int) string {
	// write code here
	stack := make([]byte, 0, len(num))
	for i := 0; i < len(num); i++ {
		// stack不为空，k不为0, 当前值小于栈中最后一个(其实是前一个元素)，则移除栈中最后一个(弹栈); 因为题目要移除后结果最小，所有要找大的数
		// 单调栈
		// "1432219",3
		// stack: [1219]
		for len(stack) > 0 && k > 0 && num[i] < stack[len(stack)-1]  {
			stack = stack[:len(stack)-1]
			k--
		}
		// 剪枝
		// 如果栈中没有元素并且当前值为‘0’, 则继续(比如：某个数字后一堆0)  比如: 100999
		if len(stack) == 0 && num[i] == '0' {
			continue
		}

		stack = append(stack, num[i])
	}

	// k还没有找完，并且stack长度大于k, 需要从stack中截取, 舍弃后面的数据
	if k != 0 && len(stack) >= k {
		stack = stack[:len(stack)-k]
	}
	// 如果 stack 为空，包含k!=0 的情况
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}

func main() {
	knums := removeKnums("1432219", 3)
	//knums := removeKnums("100999", 3)
	//knums := removeKnums("10", 1)
	fmt.Println(knums)
}
