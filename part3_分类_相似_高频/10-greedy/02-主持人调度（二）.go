package main

import (
	"fmt"
	"sort"
)

/*
BM96 主持人调度（二）
https://www.nowcoder.com/practice/4edf6e6d01554870a12f218c94e8a299?tpId=295&tqId=1267319&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
有 n 个活动即将举办，每个活动都有开始时间与活动的结束时间，第 i 个活动的开始时间是 starti ,第 i 个活动的结束时间是 endi ,举办某个活动就需要为该活动准备一个活动主持人。
一位活动主持人在同一时间只能参与一个活动。并且活动主持人需要全程参与活动，换句话说，一个主持人参与了第 i 个活动，那么该主持人在 (starti,endi) 这个时间段不能参与其他任何活动。求为了成功举办这 n 个活动，最少需要多少名主持人。

示例1
输入：
[[1,2],[2,3]]
返回值：
1
说明：
只需要一个主持人就能成功举办这两个活动

示例2
输入：
[[1,3],[2,4]]
返回值：
2
说明：
需要两个主持人才能成功举办这两个活动

*/

// 合并区间 + 分糖果
/*
1. 区间能连接   [1，2] [3，4]
2. 区间交叉     [1,3] [2,4]
3. 区间完全重合  [1,3] [1,3]
*/
func minmumNumberOfHost(startEnd [][]int) int {
	// write code here
	starts := make([]int, len(startEnd))
	ends := make([]int, len(startEnd))
	// 将startEnd 中元素的头和尾分别摘出来
	for i := 0; i < len(startEnd); i++ {
		starts[i] = startEnd[i][0]
		ends[i] = startEnd[i][1]
	}
	// 排序
	sort.Ints(starts)
	sort.Ints(ends)
	// 遍历，判断
	count, right := 0, 0
	for left := 0; left < len(startEnd); left++ {
		// 新开始的节目大于上一轮结束的时间，主持人不变
		// [1,3] [1,3] [1,3]
		// [1,3] [2,5] [3,5]
		if starts[left] >= ends[right] {
			right++
		} else {
			count++
		}
	}

	return count
}

func main() {
	host := minmumNumberOfHost([][]int{{1, 3}, {2, 5}, {3, 5}})
	//[2147483646,2147483647],[-2147483648,-2147483647],[2147483646,2147483647],[-2147483648,-2147483647],[2147483646,2147483647],[-2147483648,-2147483647],[2147483646,2147483647],[-2147483648,-2147483647],[2147483646,2147483647],[-2147483648,-2147483647]
	//host := minmumNumberOfHost([][]int{{2147483646, 2147483647}, {-2147483648, -2147483647}, {2147483646, 2147483647}, {-2147483648, -2147483647}, {2147483646, 2147483647}, {-2147483648, -2147483647}})
	fmt.Println(host)
}
