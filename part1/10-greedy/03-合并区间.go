package main

import "sort"

/*
BM89 合并区间
https://www.nowcoder.com/practice/69f4e5b7ad284a478777cb2a17fb5e6a?tpId=295&tqId=691&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给出一组区间，请合并所有重叠的区间。
请保证合并后的区间按区间起点升序排列。

数据范围：区间组数 0 \le n \le 2 \times 10^50≤n≤2×10
5
 ，区间内 的值都满足 0 \le val \le 2 \times 10^50≤val≤2×10
5

要求：空间复杂度 O(n)O(n)，时间复杂度 O(nlogn)O(nlogn)
进阶：空间复杂度 O(val)O(val)，时间复杂度O(val)O(val)
示例1
输入：
[[10,30],[20,60],[80,100],[150,180]]
复制
返回值：
[[10,60],[80,100],[150,180]]
复制
示例2
输入：
[[0,10],[10,20]]
复制
返回值：
[[0,20]]

*/

/*
排序 + 贪心  动态规划的一种

什么样的区间能够合并，那肯定是有交叉的区间，即后一个区间的首小于前一个区间的尾，这时候可以将这种交叉区间的尾合并延长区间：
// 区间有重叠，更新结尾
if(intervals[i].start <= res.back().end)
    res.back().end = max(res.back().end, intervals[i].end);


step 1：既然要求重叠后的区间按照起点位置升序排列，我们就将所有区间按照起点位置先进行排序。使用sort函数进行排序
step 2：排序后的第一个区间一定是起点值最小的区间，我们将其计入返回数组res，然后遍历后续区间。
step 3：后续遍历过程中，如果遇到起点值小于res中最后一个区间的末尾值的情况，那一定是重叠，取二者最大末尾值更新res中最后一个区间即可。
step 4：如果遇到起点值大于res中最后一个区间的末尾值的情况，那一定没有重叠，后续也不会有这个末尾的重叠区间了，因为后面的起点只会更大，因此可以将它加入res。
*/

/*
class Solution:
    def merge(self , intervals: List[Interval]) -> List[Interval]:
        res = list()
        #去除特殊情况
        if len(intervals) == 0:
            return res
        #按照区间首排序
        intervals.sort(key=cmp_to_key(lambda a,b:a.start - b.start))
        #放入第一个区间
        res.append(intervals[0])
        #遍历后续区间，查看是否与末尾有重叠
        for i in range(len(intervals)):
            #区间有重叠，更新结尾
            if intervals[i].start <= res[-1].end:
                res[-1].end = max(res[-1].end, intervals[i].end)
            #区间没有重叠，直接加入
            else:
                res.append(intervals[i])
        return res

*/

type Interval struct {
	Start int
	End   int
}

func merge(intervals []*Interval) []*Interval {
	// write code here
	res := []*Interval{}
	if len(intervals) == 0 {
		return res
	}
	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start < intervals[j].Start {
			return true
		} else if intervals[i].Start == intervals[j].Start {
			return intervals[i].End < intervals[j].End
		} else {
			return false
		}
	})

	res = append(res, intervals[0])

	for i := 0; i < len(intervals); i++ {
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
