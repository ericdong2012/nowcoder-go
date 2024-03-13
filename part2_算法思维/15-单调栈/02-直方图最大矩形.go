package _5_单调栈

import "fmt"

/*
直方图最大矩形
https://www.nowcoder.com/practice/bfaac4eebe5947af80912d1bcfcf2baa

给定一个数组heights，长度为n，height[i]是在第i点的高度，那么height[i]表示的直方图，能够形成的最大矩形是多少?
1.每个直方图宽度都为1
2.直方图都是相邻的
3.如果不能形成矩形，返回0即可
4.保证返回的结果不会超过231-1

示例1
输入：
[3,4,7,8,1,2]
返回值：
14

示例2
输入：
[1,7,3,2,4,5,8,2,7]
返回值：
16

*/

// 和接雨水，盛水最多的容器 比较看
func largestRectangleArea(height []int) int {
	// 单调栈 单调递增，出栈时计算最大面积
	if len(height) < 1 {
		return 0
	}
	if len(height) == 1 {
		return height[0]
	}
	// 该处非常重要，height 是一个升序的序列，下面的逻辑按照降序执行的，没法进入，最终结果会是0, 所以加入0， 让序列能降序
	height = append(height, 0)
	// 装索引
	stack := make([]int, 0)
	stack = append(stack, -1) // 最边上设置为-1保证不能出栈 保证栈非空
	// 结果
	res := 0
	// [3,4,7,8,1,2]
	for i := 0; i < len(height); i++ {
		// 出栈逻辑, 前一个值 大于当前值, 说明开始降序
		// 出栈时计算最大面积, 其中宽度 = 当前索引 - 栈中倒数第二个值 - 1
		for stack[len(stack)-1] != -1 && height[stack[len(stack)-1]] > height[i] {
			hei := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			// 想清楚这里, 调试看看
			// 当前索引 - ( 前前一个索引 + 1 ),  比如 [2, 4, 0] , i = 2, 计算 0, 4 区间的面积时，wid = 2 - ( 0 + 1 ) =  1 (因为宽度都是1)
			wid := i - (stack[len(stack)-1] + 1)
			res = max(res, hei*wid)
		}
		// 入栈 入下标
		stack = append(stack, i)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	area := largestRectangleArea([]int{3, 4, 7, 8, 1, 2})
	fmt.Println(area)
}
