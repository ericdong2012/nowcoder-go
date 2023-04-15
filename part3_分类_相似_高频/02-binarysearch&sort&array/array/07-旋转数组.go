package main

/*
NC110 旋转数组
https://www.nowcoder.com/practice/e19927a8fd5d477794dac67096862042?tpId=117&tqId=37834&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590,5058,578&title=


描述
一个数组A中存有 n 个整数，在不允许使用另外数组的前提下，将每个整数循环向右移 M（ M >=0）个位置，

示例1
输入：
6,2,[1,2,3,4,5,6]
返回值：
[5,6,1,2,3,4]

示例2
输入：
4,0,[1,2,3,4]
返回值：
[1,2,3,4]

*/

/*
三次逆转

step 1：因为mmm可能大于nnn，因此需要对nnn取余，因为每次长度为nnn的旋转数组相当于没有变化。
step 2：第一次将整个数组翻转，得到数组的逆序，它已经满足了右移的整体出现在了左边。
step 3：第二次就将左边的mmm个元素单独翻转，因为它虽然移到了左边，但是逆序了。
step 4：第三次就将右边的n−mn-mn−m个元素单独翻转，因此这部分也逆序了


class Solution:
    def solve(self , n: int, m: int, a: List[int]) -> List[int]:
        #取余，因为每次长度为n的旋转数组相当于没有变化
        m = m % n
        #第一次逆转全部数组元素
        a.reverse()
        b = a[:m]
        #第二次只逆转开头m个
        b.reverse()
        c = a[m:]
        #第三次只逆转结尾m个
        c.reverse()
		# 最后拼接
        a[:m] = b
        a[m:] = c
        return a

*/

func solve2(n int, m int, a []int) []int {
	// write code here
	if m == 0 {
		return a
	}
	stack := []int{}
	if m > 0 && m <= n {
		stack = append(stack, a[n-m:]...)
		stack = append(stack, a[:n-m]...)
	}
	if m > n {
		newM := m % n
		stack = append(stack, a[n-newM:]...)
		stack = append(stack, a[:n-newM]...)
	}
	return stack
}
