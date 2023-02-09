package main

/*
NC222 插入区间
https://www.nowcoder.com/practice/1d784b5472ab4dde88ea2331d16ee909?tpId=117&tqId=39447&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590,5058,578&title=

描述
给定一个无重叠的，按照区间起点升序排列的区间列表，在列表中插入一个新区间，如果与原区间有重合，则合并，请返回插入后的区间列表。

示例1
输入:
[[2,5],[6,11]],[5,6]
返回值:
[[2,11]]

示例2
输入:
[],[2,5]
返回值:
[[2,5]]

*/

type Interval struct {
	Start int
	End   int
}

// 删除指定区间的值
func del(Intervals []*Interval, i1, j1 int) []*Interval {
	res := make([]*Interval, 0)
	res = append(res, Intervals[:i1]...)
	res = append(res, Intervals[j1+1:]...)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insertInterval(Intervals []*Interval, newInterval *Interval) []*Interval {
	if len(Intervals) == 0 {
		res := make([]*Interval, 0)
		res = append(res, newInterval)
		return res
	}
	for i, _ := range Intervals {
		if Intervals[i].End < newInterval.Start {
			continue
		} else {
			// 合并新加入的
			if Intervals[i].End < newInterval.End {
				Intervals[i].End = newInterval.End
				if Intervals[i].Start > newInterval.Start {
					Intervals[i].Start = newInterval.Start
				}
			}
			// 合并剩余所有的
			count := i
			for count+1 < len(Intervals) && Intervals[count].End >= Intervals[count+1].Start {
				Intervals[count+1].End = max(Intervals[count].End, Intervals[count+1].End)
				count++
			}
			Intervals[i].End = Intervals[count].End
			return del(Intervals, i+1, count)
		}
	}
	return Intervals
}
