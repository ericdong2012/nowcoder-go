package _5_单调栈

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

func largestRectangleArea(height []int) int {
	// 单调栈 单调递增，出栈时计算最大面积
	if len(height) < 1 {
		return 0
	}
	if len(height) == 1 {
		return height[0]
	}
	height = append(height, 0)
	stack := make([]int, 0)
	stack = append(stack, -1) // 最边上设置为-1保证不能出栈 保证栈非空
	res := 0
	// [3,4,7,8,1,2]
	for i := 0; i < len(height); i++ {
		// 出栈逻辑,  前一个值 大于当前值,  开始降序
		// 出栈时计算最大面积  其中宽度 当前索引 - 栈中倒数第二个值 - 1
		for stack[len(stack)-1] != -1 && height[stack[len(stack)-1]] > height[i] {
			//hei := height[stack[len(stack)-1]]
			//wid := i - stack[len(stack)-2] - 1
			//res = max(res, hei*wid)
			//stack = stack[:len(stack)-1]

			hei := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			wid := i - stack[len(stack)-1] - 1
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
