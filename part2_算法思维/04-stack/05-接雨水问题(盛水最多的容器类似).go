package main

/*
BM94 接雨水问题
https://www.nowcoder.com/practice/31c1aed01b394f0b8b7734de0324e00f?tpId=295&tqId=1002045&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个整形数组arr，已知其中所有的值都是非负的，将这个数组看作一个柱子高度图，计算按此排列的柱子，下雨之后能接多少雨水。(数组以外的区域高度视为0)

要求：时间复杂度 O(n)

示例1
输入：
[3,1,2,5,2,4]
返回值：
5

说明：
数组 [3,1,2,5,2,4] 表示柱子高度图，在这种情况下，可以接 5个单位的雨水，蓝色的为雨水，如题面图。


示例2
输入：
[4,5,1,3,2]
返回值：
2

*/

/*
// 总结： 第二高减去剩下的， 并累加

方法一： 双指针
	那这种情况下它是不是将一个水桶分割成了两个水桶，而中间的那条边就是两个水桶的边。
	step 1：检查数组是否为空的特殊情况
	step 2：准备双指针，分别指向数组首尾元素，代表最初的两个边界
	step 3：指针往中间遍历，遇到更低柱子就是底，用较短的边界减去底(计算公式)就是这一列的接水量，遇到更高的柱子就是新的边界，更新(累加)边界大小。

class Solution:
    def maxWater(self , arr: List[int]) -> int:
        #排除空数组
        if len(arr) == 0:
            return 0
        res = 0
        #左右双指针
        left = 0
        right = len(arr) - 1
        # 中间区域的边界高度
        maxL = 0
        maxR = 0
        #直到左右指针相遇
        while left < right:
            #每次维护往两边的最大边界
            maxL = max(maxL, arr[left])
            maxR = max(maxR, arr[right])
            # 较短的边界确定该格子的水量
            if maxR > maxL:
                res += maxL - arr[left]
                left += 1
            else:
                res += maxR - arr[right]
                right -= 1
        return res


方法二：单调栈

*/
func maxWater(arr []int) int64 {
	// write code here
	if len(arr) == 0 {
		return 0
	}
	var res int
	left, right := 0, len(arr)-1
	var maxL, maxR int
	for left <= right {
		maxL = max1(maxL, arr[left])
		maxR = max1(maxR, arr[right])
		if arr[left] <= arr[right] {
			res += maxL - arr[left]
			left += 1
		} else {
			res += maxR - arr[right]
			right -= 1
		}
	}

	return int64(res)
}

func max1(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {

}
