package main

import "strconv"

/*
BM74 数字字符串转化成IP地址
https://www.nowcoder.com/practice/ce73540d47374dbe85b3125f57727e1e?tpId=295&tqId=653&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
现在有一个只包含数字的字符串，将该字符串转化成IP地址的形式，返回所有可能的情况。
例如：
给出的字符串为"25525522135",
返回["255.255.22.135", "255.255.221.35"]. (顺序没有关系)

要求：空间复杂度 O(n!),时间复杂度 O(n!)
注意：ip地址是由四段数字组成的数字序列，格式如 "x.x.x.x"，其中 x 的范围应当是 [0,255]。

示例1
输入：
"25525522135"
返回值：
["255.255.22.135","255.255.221.35"]

示例2
输入：
"1111"
返回值：
["1.1.1.1"]

示例3
输入：
"000256"
返回值：
"[]"
*/

/*

递归 + 回溯剪枝

枚举

对于IP字符串，如果只有数字，则相当于需要我们将IP地址的三个点插入字符串中，而第一个点的位置只能在第一个字符、第二个字符、第三个字符之后，而第二个点只能在第一个点后1-3个位置之内，第三个点只能在第二个点后1-3个位置之内，且要要求第三个点后的数字数量不能超过3，因为IP地址每位最多3位数字。

具体做法：
step 1：依次枚举这三个点的位置。
step 2：然后截取出四段数字。
step 3：比较截取出来的数字，不能大于255，且除了0以外不能有前导0，然后才能组装成IP地址加入答案中。

输入：
"25525522135"
返回值：
["255.255.22.135","255.255.221.35"]

class Solution:
    def restoreIpAddresses(self , s: str) -> List[str]:
        res = []
        n = len(s)
        i = 1
        # 遍历IP的点可能的位置（第一个点）
        while i < 4 and i < n - 2:
            j = i + 1
            # 第二个点的位置
            while j < i + 4 and j < n - 1:
                k = j + 1
                # 第三个点的位置
                while k < j + 4 and k < n:
                    # 最后一段剩余数字不能超过3
                    if n - k >= 4:
                        k += 1
                        continue
                    # 从点的位置分段截取
                    a = s[0: i]
                    b = s[i: j]
                    c = s[j: k]
                    d = s[k:]
                    # IP每个数字不大于255
                    if int(a) > 255 or int(b) > 255 or int(c) > 255 or int(d) > 255:
                        k += 1
                        continue

                    # 排除前导0的情况
                    if (len(a) != 1 and a[0] == '0') or (len(b) != 1 and b[0] == '0') or  (len(c) != 1 and c[0] == '0') or (len(d) != 1 and d[0] == '0'):
                        k += 1
                        continue

                    # 组装IP地址
                    temp = a + "." + b + "." + c + "." + d
                    res.append(temp)
                    k += 1
                j += 1
            i += 1
        return res

*/

func restoreIpAddresses(s string) []string {
	// write code here
	res := []string{}
	n := len(s)
	// 从1开始
	i := 1
	// 第二位
	// i < 4 && i < n-2 是因为第一位最长三位，后面三个数字最短也要3位，预留
	for i < 4 && i < n-2 {
		j := i + 1
		// 第三位
		for j < i+4 && j < n-1 {
			k := j + 1
			// 第四位
			for k < j+4 && k < n {
				// 判断最后一位长度， 如果预留给第四位的大于4， 则k要往前走
				if n-k >= 4 {
					k++
					continue
				}
				// 截取
				a := s[0:i]
				b := s[i:j]
				c := s[j:k]
				d := s[k:n]
				// ip中的每个位不大于255
				a1, _ := strconv.Atoi(a)
				b1, _ := strconv.Atoi(b)
				c1, _ := strconv.Atoi(c)
				d1, _ := strconv.Atoi(d)
				if a1 > 255 || b1 > 255 || c1 > 255 || d1 > 255 {
					k++
					continue
				}
				// 排除先导0
				if (len(a) != 1 && a[0] == '0') || (len(b) != 1 && b[0] == '0') || (len(c) != 1 && c[0] == '0') || (len(d) != 1 && d[0] == '0') {
					k ++
					continue
				}
				// 拼接，添加到结果中
				temp := a + "." + b + "." + c + "." + d
				res = append(res, temp)
				k++
			}
			j++
		}
		i++
	}

	return res
}
