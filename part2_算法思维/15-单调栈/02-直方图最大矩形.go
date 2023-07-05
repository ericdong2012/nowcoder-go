package _5_单调栈


/*
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

func largestRectangleArea( heights []int ) int {
	// write code here
	if len(heights) == 0 {
		return 0
	}
	left, right := make([]int, len(heights)), make([]int, len(heights))
	var stack []int
	for i := 0; i < len(heights); i++ {
		right[i] = len(heights)
	}
	for i, v := range heights {
		for len(stack) != 0 && heights[stack[len(stack) - 1]] >= v {
			right[stack[len(stack) - 1]] = i
			stack = stack[: len(stack) - 1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack) - 1]
		}
		stack = append(stack, i)
	}
	var ans int
	for i := 0; i < len(heights); i++ {
		if ans < (right[i]-left[i]-1) * heights[i] {
			ans = (right[i]-left[i]-1) * heights[i]
		}
	}
	return ans
}