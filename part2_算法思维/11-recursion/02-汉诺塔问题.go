package main

import "fmt"

/*
NC67 汉诺塔问题
https://www.nowcoder.com/practice/7d6cab7d435048c4b05251bf44e9f185?tpId=117&tqId=37787&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=591&title=


描述
我们有由底至上为从大到小放置的 n 个圆盘，和三个柱子（分别为左/中/右即left/mid/right），开始时所有圆盘都放在左边的柱子上，按照汉诺塔游戏的要求我们要把所有的圆盘都移到右边的柱子上，要求一次只能移动一个圆盘，而且大的圆盘不可以放到小的上面。

请实现一个函数打印最优移动轨迹。

给定一个 `int n` ，表示有 n 个圆盘。请返回一个 `string` 数组，其中的元素依次为每次移动的描述。描述格式为： `move from [left/mid/right] to [left/mid/right]`。



输入：
2
返回值：
["move from left to mid","move from left to right","move from mid to right"]

*/




/*

1   可以直接移动
2   需要借助mid才能移动
3   借助righ移动了3-1=2到mid上，再将mid到right看成新的2个盘的汉诺塔，借助left即可

仅需要将其当作n-1再利用旁边另一个塔即可实现递归


https://zhuanlan.zhihu.com/p/501020136
汉诺塔问题的解决方案可以分为3步：
1、把n-1个盘子从 left 搬到 mid柱子上, 借助right    从A移动到B
2、把剩下最大的那一个盘子从left搬到right柱子上       从A移动到C
3、把n-1个盘子从 mid  搬到right柱子上, 借助 left   从B移动到C

有点像中序遍历

*/

func getSolution(n int) []string {
	// write code here
	res := make([]string, 0)
	hanno(n, "left", "mid", "right", &res)
	return res
}

func hanno(n int, a, b, c string, res *[]string) {
	if n == 0 {
		// *res = append(*res, "move from " + a + " to " + c)
		return
	}
	// 要按照步骤全局看
	//把n-1个盘子从left 搬到mid上去，借助right
	hanno(n-1, a, c, b, res)
	//把第n个盘子从left搬到right上
	*res = append(*res, "move from "+a+" to "+c+",")
	//把n-1个盘子从mid，搬到right, 借助left上去
	hanno(n-1, b, a, c, res)
}

func main04() {
	solution := getSolution(3)
	fmt.Println(solution)
}
