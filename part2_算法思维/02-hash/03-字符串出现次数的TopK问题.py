
"""
NC97 字符串出现次数的TopK问题
https://www.nowcoder.com/practice/fd711bdfa0e840b381d7e1b82183b3ee?tpId=117&tqId=37809&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=585&title=

描述
给定一个字符串数组，再给定整数 k ，请返回出现次数前k名的字符串和对应的次数。
返回的答案应该按字符串出现频率由高到低排序。如果不同的字符串有相同出现频率，按字典序排序。
对于两个字符串，大小关系取决于两个字符串从左到右第一个不同字符的 ASCII 值的大小关系。
比如"ah1x"小于"ahb"，"231"<”32“
字符仅包含数字和字母

要求：空间复杂度 O(n)，时间复杂度O(nlog k)

示例1
输入：
["a","b","c","b"], 2
返回值：
[["b","2"],["a","1"]]
说明：
"b"出现了2次，记["b","2"]，"a"与"c"各出现1次，但是a字典序在c前面，记["a","1"]，最后返回[["b","2"],["a","1"]]

示例2
输入：
["123","123","231","32"], 2
返回值：
[["123","2"],["231","1"]]
说明：
 "123"出现了2次，记["123","2"]，"231"与"32"各出现1次，但是"231"字典序在"32"前面，记["231","1"]，最后返回[["123","2"],["231","1"]]

示例3
输入：
["abcd","abcd","abcd","pwb2","abcd","pwb2","p12"], 3
返回值：
[["abcd","4"],["pwb2","2"],["p12","1"]]


"""

# python
from typing import List


class Solution:
    def topKstrings(self, strings: List[str], k: int) -> List[List[str]]:
        # write code here
        temp = defaultdict(int)
        for i in range(len(strings)):
#             if strings[i] not in dict:
#                 dict[strings[i]] = 1
#             else:
#                 dict[strings[i]] += 1

            temp[strings[i]] +=1
        return [sorted(dict.items(), key=lambda x: (-x[1], x[0]))[:k]]


if __name__ == '__main__':
    result = Solution().topKstrings(["a", "b", "c", "b"], 2)
    print(result)


# go
"""

import (
	"fmt"
	"sort"
	"strconv"
)

//type Result []Top
//
//func (x Result) Len() int {
//	return len(x)
//}
//func (x Result) Less(i, j int) bool {
//	if x[i].time > x[i].time {
//		return true
//	} else {
//		return x[i].char < x[j].char
//	}
//}
//func (x Result) Swap(i, j int) {
//	x[i], x[j] = x[j], x[i]
//}

func topKstrings(strings []string, k int) [][]string {
	// write code here
	temp := make(map[string]int, 0)
	for i := 0; i < len(strings); i++ {
		temp[strings[i]]++
	}

	type Top struct {
		char string
		time int
	}
	result := []Top{}
	for k, v := range temp {
		result = append(result, Top{char: k, time: v})
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].time < result[j].time {
			return result[j].time < result[i].time
		} else if result[i].time == result[j].time {
			return result[i].char < result[j].char
		}
		return result[i].time > result[j].time
	})
	//fmt.Println(result)
	rest := [][]string{}
	temp1 := result[:k]
	for i := 0; i < k; i++ {
		rest = append(rest, []string{temp1[i].char, strconv.Itoa((temp1[i].time))})
	}

	return rest
}

func main() {
	//s := topKstrings([]string{"abcd", "abcd", "abcd", "pwb2", "abcd", "pwb2", "p12"}, 3)
	s := topKstrings([]string{"abcd", "abcd", "pwb2", "abcd", "p12", "pwb2", "abcd"}, 3)
	fmt.Println(s)
}



"""
