package main

import "fmt"

/*
BM93 盛水最多的容器
https://www.nowcoder.com/practice/3d8d6a8e516e4633a2244d2934e5aa47?tpId=295&tqId=2284579&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定一个数组height，长度为n，每个数代表坐标轴中的一个点的高度，height[i]是在第i点的高度，请问，从中选2个高度与x轴组成的容器最多能容纳多少水
1.你不能倾斜容器
2.当n小于2时，视为不能形成容器，请返回0
3.数据保证能容纳最多的水不会超过整形范围，即不会超过231-1

如输入的height为[1,7,3,2,4,5,8,2,7]，那么如下图:

示例1
输入：
[1,7,3,2,4,5,8,2,7]
返回值：
49

示例2
输入：
[2,2]
返回值：
2

示例3
输入：
[5,4,3,2,1,5]
返回值：
25

*/

/*
这道题利用了水桶的短板原理，较短的一边控制最大水量，因此直接用较短边长乘底部两边距离就可以得到当前情况下的容积。但是要怎么找最大值呢？

可以利用贪心思想：我们都知道容积与最短边长和底边长有关，与长的底边一定以首尾为边，但是首尾不一定够高，中间可能会出现更高但是底边更短的情况，
因此我们可以使用对撞双指针向中间靠，这样底边长会缩短，
因此还想要有更大容积只能是增加最短变长，此时我们每次指针移动就移动较短的一边，因为贪心思想下较长的一边比较短的一边更可能出现更大容积。

//优先舍弃较短的边
if(height[left] < height[right])
    left++;
else
    right--;

具体做法:
step 1：优先排除不能形成容器的特殊情况。
step 2：初始化双指针指向数组首尾，每次利用上述公式(最短边长和底边长)计算当前的容积，维护一个最大容积作为返回值。
step 3：对撞双指针向中间靠，但是依据贪心思想，每次指向较短边的指针向中间靠，另一指针不变。

class Solution:
    def maxArea(self , height: List[int]) -> int:
        #排除不能形成容器的情况
        if len(height) < 2:
            return 0
        res = 0
        # 双指针左右界
        left = 0
        right = len(height) - 1
        # 共同遍历完所有的数组
        while left < right:
            # 计算区域水容量
            capacity = min(height[left], height[right]) * (right - left)
            #维护最大值
            res = max(res, capacity)
            # 优先舍弃较短的边
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1
        return res
*/

// 双指针
func maxArea(height []int) int {
	// write code here
	if len(height) < 2 {
		return 0
	}

	res := 0
	left, right := 0, len(height)-1
	for left <= right {
		if height[left] <= height[right] {
			res = max(res, height[left]*(right-left))
			left++
		} else {
			res = max(res, height[right]*(right-left))
			right--
		}
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
	area := maxArea([]int{1, 7, 3, 2, 4, 5, 8, 2, 7})
	fmt.Println(area)
}
