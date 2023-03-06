package main

import "fmt"

/*
NC157 单调栈
https://www.nowcoder.com/practice/ae25fb47d34144a08a0f8ff67e8e7fb5?tpId=117&tqId=37867&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=

描述
给定一个长度为 n 的可能含有重复值的数组 arr ， 找到每一个 i 位置左边和右边离 i 位置最近且值比 arri 小的位置。
请设计算法，返回一个二维数组，表示所有位置相应的信息。位置信息包括：两个数字 l 和 r，如果不存在，则值为 -1，下标从 0 开始。

进阶：空间复杂度 O(n) , 时间复杂度 O(n) .

示例1
输入：
[3,4,1,5,6,2,7]
返回值：
[[-1,2],[0,2],[-1,-1],[2,5],[3,5],[2,-1],[5,-1]]

示例2
输入：
[1,1,1,1]
返回值：
[[-1,-1],[-1,-1],[-1,-1],[-1,-1]]

*/

/*
单调栈：单调栈就是使栈内元素单调递增或者单调递减的栈，单调栈也只能在栈顶操作。

while(栈顶元素比入栈元素小且栈不空){
	栈顶元素弹出
}
元素入栈


让我们模拟一遍单调栈的运行过程，这里的单调栈是单调递增栈。
我们现在有6个数 4,1,7,9,3,6。
我们当前栈内没有元素，将4加入。现在栈内元素应该是4。
当前元素为1，我们发现加入之后不能单调，于是4出栈，加入1。当前栈内元素为2。
接下来是7,9。我们发现加入不会破坏单调性，于是直接加入，栈内元素1,7,9。
遇到3，只能吐出栈内7，9，加入3。
最后加入6。结束。

对于本题来说需要返回左右的最小值，很显然只需要进行两次遍历，第一次遍历啊超出右边的最小值，然后第二次遍历找出最左边最小值。
*/

func foundMonotoneStack(nums []int) [][]int {
	// write code here
	n := len(nums)
	// 单调栈， 存放索引，没有 -1
	stack := []int{}
	// 构造数据结构
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, 2)
	}
	for i := 0; i < n; i++ {
		for len(stack) != 0 && nums[i] < nums[stack[len(stack)-1]] {
			stackValue := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			leftIndex := -1
			if len(stack) != 0 {
				leftIndex = stack[len(stack)-1]
			}
			ans[stackValue][0] = leftIndex
			ans[stackValue][1] = i
		}
		stack = append(stack, i)
	}

	for len(stack) != 0 {
		stackValue := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		leftIndex := -1
		if len(stack) != 0 {
			leftIndex = stack[len(stack)-1]
		}
		ans[stackValue][0] = leftIndex
		ans[stackValue][1] = -1
	}

	// 从左往右，单调递增
	//for i := 0; i < n; i++ {
	//	//保证栈顶元素是小于待入栈的
	//	for len(stack) != 0 && nums[i] <= nums[stack[len(stack)-1]] {
	//		stackValue := stack[len(stack)-1]
	//		stack = stack[:len(stack)-1]
	//		//栈为空，说明该位置元素左侧没有比它小的数
	//		if len(stack) == 0 {
	//			ans[stackValue][0] = -1
	//		} else {
	//			//如果栈不为空,则左侧比其值小的元素位置距离最近
	//			index := stack[len(stack)-1]
	//			//stack = stack[:len(stack)-1]
	//			ans[stackValue][0] = index
	//		}
	//	}
	//	stack = append(stack, i)
	//}
	//
	//// 清空栈
	//for len(stack) != 0 {
	//	stack = stack[:len(stack)-1]
	//}
	//
	//// 从右往左，单调递减
	//for i := n - 1; i >= 0; i-- {
	//	for len(stack) != 0 && nums[i] <= nums[stack[len(stack)-1]] {
	//		stackValue := stack[len(stack)-1]
	//		stack = stack[:len(stack)-1]
	//		if len(stack) == 0 {
	//			ans[stackValue][1] = -1
	//		} else {
	//			index := stack[len(stack)-1]
	//			//stack = stack[:len(stack)-1]
	//			ans[stackValue][1] = index
	//		}
	//	}
	//	stack = append(stack, i)
	//}

	return ans
}

func main() {
	stack := foundMonotoneStack([]int{0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(stack)
}
