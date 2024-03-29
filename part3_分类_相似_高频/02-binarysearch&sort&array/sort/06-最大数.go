package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
NC111 最大数
https://www.nowcoder.com/practice/fc897457408f4bbe9d3f87588f497729?tpId=117&tqId=37835&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590&title=

描述
给定一个长度为n的数组nums，数组由一些非负整数组成，现需要将他们进行排列并拼接，每个数不可拆分，使得最后的结果最大，返回值需要是string类型，否则可能会溢出。
进阶：时间复杂度O(nlogn) ，空间复杂度：O(n)

示例1
输入：
[30,1]
返回值：
"301"

示例2
输入：
[2,20,23,4,8]
返回值：
"8423220"

*/

func solve(nums []int) string {
	// write code here
	// 本质上还是排序，但是是将每个元素的首位拿出来比较，如果首位相同，则拿剩下的位跟之前的比，直到比出大小
	/*
		class cmp(str):
		    def __lt__(x, y):
		        return x+y > y+x
		class Solution:
		    def solve(self , nums ):
		        # write code here
		        #将整型的数字转化为字符串
		        strs = [str(i) for i in nums]
		        #排序
		        strs = sorted(strs,key = cmp)
		        #这个地方需要注意如果第一个字符串已经是0了，那么直接输出0
		        if strs[0] == "0": return "0"
		        return "".join(strs)
	*/

	sort.Slice(nums, func(i, j int) bool {
		// 此时拿到的是字符串
		var left, right = strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		//fmt.Println("left:", left)
		//fmt.Println("right:", right)
		// 要求返回较大值，所以降序排   220  202
		return left+right > right+left
	})
	// 看看2, 20 的排序
	//fmt.Println("nums:", nums)
	var res string
	for _, num := range nums {
		// 如果有一个0则添加进去， 下方的代码会添加进去
		// 如果是多个0挨着，当前项是0， 之前添加过0， 则跳过
		if num == 0 && res == "0" {
			continue
		}
		res += strconv.Itoa(num)
	}

	return res
}

func main06() {
	s := solve([]int{2, 20, 23, 4, 8})
	fmt.Println(s)
}