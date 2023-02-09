package main

import "fmt"

/*
NC67 汉诺塔问题
https://www.nowcoder.com/practice/7d6cab7d435048c4b05251bf44e9f185?tpId=117&tqId=37787&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=591&title=


描述
我们有由底至上为从大到小放置的 n 个圆盘，和三个柱子（分别为左/中/右即left/mid/right），开始时所有圆盘都放在左边的柱子上，按照汉诺塔游戏的要求我们要把所有的圆盘都移到右边的柱子上，要求一次只能移动一个圆盘，而且大的圆盘不可以放到小的上面。

请实现一个函数打印最优移动轨迹。

给定一个 `int n` ，表示有 n 个圆盘。请返回一个 `string` 数组，其中的元素依次为每次移动的描述。描述格式为： `move from [left/mid/right] to [left/mid/right]`。

数据范围：1\le n \le 161≤n≤16
要求：时间复杂度 O(3^n)O(3
n
 ) ， 空间复杂度 O(3^n)O(3
n
 )
示例1
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




汉诺塔问题的解决方案可以分为3步：
1、把n-1个盘子从left 借助 right，搬到mid柱子上
2、把剩下最大的那一个盘子从left搬到right柱子上
3、把n-1个盘子从mid 借助 left，搬到right柱子上。

至于如何把n-1个盘子搬到另一个柱子上，同样参照上面的3步，不过此时柱子扮演的left，mid，right角色已经改变；对于n-2，n-3等等以此类推。

	vector<string> solution;

    void Hanoi(int n, string left, string mid, string right){
        if (n==0) return;
        Hanoi(n-1,left,right,mid);//把n-1个盘子从left借助right搬到mid上去。
        solution.push_back("move from " + left + " to " + right);//把第n个盘子从left搬到right上。
        Hanoi(n-1,mid,left,right);//把n-1个盘子从mid借助left搬到right上去。
    }

    vector<string> getSolution(int n) {
        Hanoi(n, "left", "mid", "right");
        return solution;
    }

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
	//把n-1个盘子从left借助right搬到mid上去。
	hanno(n-1, a, c, b, res)
	//把第n个盘子从left搬到right上。
	*res = append(*res, "move from "+a+" to "+c+",")
	//把n-1个盘子从mid借助left搬到right上去。
	hanno(n-1, b, a, c, res)
}

func main() {
	solution := getSolution(3)
	fmt.Println(solution)
}
