package main

import (
	"sort"
)

/*
NC37 合并区间
https://www.nowcoder.com/practice/69f4e5b7ad284a478777cb2a17fb5e6a?tpId=117&tqId=37737&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590&title=

描述
给出一组区间，请合并所有重叠的区间。
请保证合并后的区间按区间起点升序排列。

要求：空间复杂度 O(n)，时间复杂度 O(nlogn)
进阶：空间复杂度 O(val)，时间复杂度O(val)

示例1
输入：
[[10,30],[20,60],[80,100],[150,180]]
返回值：
[[10,60],[80,100],[150,180]]

示例2
输入：
[[0,10],[10,20]]
返回值：
[[0,20]]

*/

type Interval struct {
	Start int
	End   int
}

func merge(intervals []*Interval) []*Interval {
	res := []*Interval{}
	if len(intervals) == 0 {
		return res
	}
	// 将所有数据按字典序排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	res = append(res, intervals[0])
	// 判断首位，末尾，考虑是否合并
	for i := 1; i < len(intervals); i++ {
		/*
				[10, 20] [20, 60]
				[10, 30] [20, 40]
			    [10, 30] [20, 25]
		*/
		if intervals[i].Start <= res[len(res)-1].End {
			res[len(res)-1].End = max(res[len(res)-1].End, intervals[i].End)
		} else {
			res = append(res, intervals[i])
		}

	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
