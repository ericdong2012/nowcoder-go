package main

/*
NC76 用两个栈实现队列
https://www.nowcoder.com/practice/54275ddae22f475981afa2244dd448c6?tpId=117&tqId=37774&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=

描述
用两个栈来实现一个队列，使用n个元素来完成 n 次在队列尾部插入整数(push)和n次在队列头部删除整数(pop)的功能。 队列中的元素为int类型。保证操作合法，即保证pop操作时队列内已有元素。

数据范围： n\le1000n≤1000
要求：存储n个元素的空间复杂度为 O(n)O(n) ，插入与删除的时间复杂度都是 O(1)O(1)
示例1
输入：
["PSH1","PSH2","POP","POP"]
复制
返回值：
1,2
复制
说明：
"PSH1":代表将1插入队列尾部
"PSH2":代表将2插入队列尾部
"POP“:代表删除一个元素，先进先出=>返回1
"POP“:代表删除一个元素，先进先出=>返回2
示例2
输入：
["PSH2","POP","PSH1","POP"]
复制
返回值：
2,1

*/

var stack1 []int
var stack2 []int

func Push1(node int) {
	stack1 = append(stack1, node)
}

func Pop1() int {
	// 如果stack2 存在，从stack2中取
	if len(stack2) != 0 {
		tmp := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		return tmp
	}
	// 如果stack1 存在
	if len(stack1) != 0 {
		tmp := stack1[0]
		// 将stack1中的元素从尾部取出，添加到stack头部
		for i := len(stack1) - 1; i > 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = []int{}
		return tmp
	}
	// 如果stack1 不存在
	return -1
}

/*
class Solution:
    def __init__(self):
        self.A = []
        self.B = []
    def push(self, num):
        self.A.append(num)
    def pop(self):
        if self.B:
            return self.B.pop()
        if not self.A:
            return None
        else:
            while self.A:
                self.B.append(self.A.pop())
            return self.B.pop()
*/
