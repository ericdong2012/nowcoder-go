package main

import "fmt"

/*
NC237 最大矩形
https://www.nowcoder.com/practice/5720efc1bdff4ca3a7dad37ca012cb60?tpId=117&tqId=39474&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=


描述
给定一个仅包含 0 和 1 ，大小为 n*m 的二维二进制矩阵，找出仅包含 1 的最大矩形并返回面积。
保证输入的矩形中仅含有 0 和 1
例如输入[[1,0,1,0,0],[1,1,1,1,0],[1,1,1,1,1],[1,0,0,1,0]]时，对应的输出为8，所形成的最大矩形如下图所示：

示例1
输入：
[[1]]
返回值：
1

示例2
输入：
[[1,1],[0,1]]
返回值：
2
说明：
第一列的两个 1 组成一个矩形

示例3
输入：
[[0,0],[0,0]]
返回值：
0

示例4
输入：
[[1,0,1,0,0],[1,1,1,1,0],[1,1,1,1,1],[1,0,0,1,0]]
返回值：
8

*/

/*
https://www.nowcoder.com/practice/5720efc1bdff4ca3a7dad37ca012cb60?tpId=117&tqId=39474&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=


def maximalRectangle(self , matrix: List[List[int]]) -> int:
	if not matrix:return 0
	m, n=len(matrix), len(matrix[0])
	# 记录当前位置上方连续“1”的个数
	pre=[0]*(n+1)
	res=0
	for i in range(m):
		for j in range(n):
			# 前缀和
			pre[j]=pre[j]+1 if matrix[i][j]==1 else 0

		# 单调栈
		stack=[-1]
		for k,num in enumerate(pre):
			while stack and pre[stack[-1]]>num:
				index=stack.pop()
				res=max(res,pre[index]*(k-stack[-1]-1))
			stack.append(k)

	return res
*/

/*
// 暴力(下面是用该方法解决的)
本质上是对矩阵中的每行依次求直方图最大矩阵算法。
获得该行的直方图之后，就可以在heights上求解该行的直方图最大矩阵面积，例如题目 [[1,0,1,0,0],[1,1,1,1,0],[1,1,1,1,1],[1,0,0,1,0]] 中的第三行，得到的heights = [3, 2, 3, 2, 1]

有个题解的第一种解法可以看懂


// 单调栈
首先将输入拆分成一系列的直方图，只需要计算每个直方图中的最大面积，就可以计算出数组中的最大矩阵面积。  heights = [3, 2, 3, 2, 1]
而可以优化的地方主要在于第二步的largest_rectangle_area算法，暴力求解的方法时间复杂度过高，可以采用单调栈的思想。

对于输入[3, 2, 3, 2, 1]来说，在观察第一个元素时，由于不知道其右边元素是否比它大，所以可将其先记录下来
然后考察下一个元素，在考察到下一个元素时，发现其值 2 小于第一个元素3。因此，这时可以确定第一个元素的最大矩形面积，用粉红色表示
在确定第一个元素面积后，第一个元素就可以无视了，用白块表示。

*/
//func maximalRectangle(matrix [][]int) (result int) {
//	// len(matrix[0]) 列长度
//	if len(matrix) == 0 || len(matrix[0]) == 0 {
//		return 0
//	}
//	height := make([]int, len(matrix[0]))
//	res := 0
//	/*
//		[1,0,1,0,0]
//		[1,1,1,1,0]
//		[1,1,1,1,1]
//		[1,0,0,1,0]
//
//	最简单的暴力方法就是枚举每一列，每一列的高度就是矩形的高度 h。随后从这根柱子开始向两侧延伸，直到遇到高度小于h的柱子，就确定了矩形的左右边界，左右边界之间的宽度计为w，
//	对应的面积为 w×h。最后取这些列的最大值，就可以得到该行的最大面积，然后遍历所有行，就可以得到这个数组的最大面积
//	*/
//	// 行
//	for i := 0; i < len(matrix); i++ {
//		// 列
//		for j := range height {
//			// 如果数组对应的元素是1，就在之前的最大高度基础上加一，不然就置为0
//			// 之前的最大高度基础上加一 乘以 当前位置的值， 得到当前的高度
//			// 每行i 的 height 数组（从左往右）
//			height[j] = (height[j] + 1) * matrix[i][j]
//			minHeight := height[j]
//			// 回溯，要求height > 0， 计算
//			for k := j; k >= 0 && height[k] > 0; k-- {
//				minHeight = min(minHeight, height[k])
//				// 计算面积
//				s := minHeight * (j - k + 1)
//				// 取较大值
//				if res < s {
//					res = s
//				}
//			}
//		}
//	}
//
//	return res
//}

/*
	[1,0,1,0,0]
	[1,1,1,1,0]
	[1,1,1,1,1]
	[1,0,0,1,0]

	最简单的暴力方法就是枚举每一列，每一列的高度就是矩形的高度 h。随后从这根柱子开始向两侧延伸，直到遇到高度小于h的柱子，就确定了矩形的左右边界，左右边界之间的宽度计为w，
	对应的面积为 w×h。最后取这些列的最大值，就可以得到该行的最大面积，然后遍历所有行，就可以得到这个数组的最大面积
*/
func maximalRectangle(matrix [][]int) int {
	ans := 0
	// 保存某行的列数据
	height := make([]int, len(matrix[0]))
	// 行
	for _, row := range matrix {
		// 索引，每列的值
		for k, v := range row {
			if v == 1 {
				height[k]++
			} else {
				height[k] = 0
			}
		}
		maxArea := getArea(height)
		if maxArea > ans {
			ans = maxArea
		}
	}
	return ans
}

func getArea(height []int) int {
	/*
		[1,0,1,0,0]
		[1,1,1,1,0]
		[1,1,1,1,1]
		[1,0,0,1,0]

		比如第三行： [ 3, 2, 3, 2, 1]
	*/
	ans := 0
	for k, v := range height {
		// 计算以k列为高的矩形面积
		l, r := k, k
		for l-1 >= 0 && height[l-1] >= v {
			l--
		}
		for r+1 < len(height) && height[r+1] >= v {
			r++
		}
		// v是高度，r -l +1 是底
		area := v * (r - l + 1)
		if area > ans {
			ans = area
		}
	}

	return ans
}

func main() {
	rectangle := maximalRectangle([][]int{[]int{1, 0, 1, 0, 0}, []int{1, 1, 1, 1, 0}, []int{1, 1, 1, 1, 1}, []int{1, 0, 0, 1, 0}})
	fmt.Println(rectangle)
}
