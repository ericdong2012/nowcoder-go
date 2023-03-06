package main

/*
NC175 合法的括号字符串
https://www.nowcoder.com/practice/eceb50e041ec40bd93240b8b3b62d221?tpId=117&tqId=39331&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=


描述
给定一个字符串s，字符串s只包含以下三种字符: (  *  )，请你判断 s是不是一个合法的括号字符串。合法括号字符串有如下规则:
1.左括号'('必须有对应的右括号')'
2.右括号')'必须有对应的左括号'('
3.左括号必须在对应的右括号前面
4.*可以视为单个左括号，也可以视为单个右括号，或者视为一个空字符
5.空字符串也视为合法的括号字符串

示例1
输入：
"()()"
返回值：
true

示例2
输入：
"((*)"
返回值：
true

示例3
输入：
"(*)"
返回值：
true

示例4
输入：
"(((*)"
返回值：
false
*/

/*

我们维护左括号的下标栈left和星号的下标栈star
遍历字符串时，如果遇到左括号则将对应的下标入到left栈，星号栈同理
如果遇到右括号则首先在left栈中匹配，即left栈顶元素出栈；如果left栈空，则顺位到star栈出栈；如果star栈也空了，则说明仍有右括号无法匹配，则返回false结果
如果上一过程顺利完成，则还需要判断是否有多余的左括号无法匹配：则依次对left栈顶进行判断，如果left栈空了说明所有左括号已经匹配完成；如果left栈非空，则判断star栈中栈顶是否存在下标大于left栈顶的元素，因为只有这样才能对左括号进行匹配；否则返回false结果
最终剩余的结果就是true


	stack<int> left;
	stack<int> star;
	for(int i = 0; i < s.size(); i++) {                        // 处理右括号的匹配
		char c = s[i];
		if(c == '(') left.push(i);                             // 左括号下标入栈
		else if(c == '*') star.push(i);                        // 星号下标入栈
		else {
			if(!left.empty()) left.pop();
			else if(!star.empty()) star.pop();
			else return false;                                 // 右括号数量多则直接返回false
		}
	}
	while(!left.empty()) {                    // 处理多的左括号的匹配
		if(!star.empty() && left.top() < star.top()) {
			left.pop();
			star.pop();
		}
		else return false;                    // 左括号数量多则直接返回false
	}
	return true;

*/

func isValidString(s string) bool {
	// write code here
	// 存放对应元素的索引
	left, star := []int{}, []int{}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			left = append(left, i)
		} else if string(s[i]) == "*" {
			star = append(star, i)
		} else {
			if len(left) != 0 {
				left = left[:len(left)-1]
			} else if len(star) != 0 {
				star = star[:len(star)-1]
			} else {
				// 左括号并且 * 都为0， 此时右括号多
				return false
			}
		}
	}
	// 处理特例，left还有多个， 则要求 * 有多个，此时 * 当做右括号
	// (((*) => ((*
	for len(left) != 0 {
		// left 在 * 左侧， left 索引小于 * 索引
		if len(star) != 0 && left[len(left)-1] < star[len(star)-1] {
			left = left[:len(left)-1]
			star = star[:len(star)-1]
		} else {
			return false
		}
	}

	return true
}
