package main

/*
NC29 二维数组中的查找
https://www.nowcoder.com/practice/abc3fe2ce8e146608e868a70efebf62e?tpId=117&tqId=37733&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590,5058,578&title=

描述
在一个二维数组array中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
[
[1,2,8,9],
[2,4,9,12],
[4,7,10,13],
[6,8,11,15]
]
给定 target = 7，返回 true。

给定 target = 3，返回 false。

数据范围：矩阵的长宽满足 0 \le n,m \le 5000≤n,m≤500 ， 矩阵中的值满足 0 \le val \le 10^90≤val≤10
9

进阶：空间复杂度 O(1)O(1) ，时间复杂度 O(n+m)O(n+m)
示例1
输入：
7,[[1,2,8,9],[2,4,9,12],[4,7,10,13],[6,8,11,15]]
复制
返回值：
true
复制
说明：
存在7，返回true
示例2
输入：
1,[[2]]
复制
返回值：
false
复制
示例3
输入：
3,[[1,2,8,9],[2,4,9,12],[4,7,10,13],[6,8,11,15]]
复制
返回值：
false
复制
说明：
不存在3，返回false

*/

func Find(target int, array [][]int) bool {
	// write code here
	/*
		class Solution:
		    def Find(self , target: int, array: List[List[int]]) -> bool:
		        # 优先判断特殊
		        if len(array) == 0:
		            return False
		        n = len(array)
		        if len(array[0]) == 0:
		            return False
		        m = len(array[0])
		        i = n-1
		        j = 0
		        # 从最左下角的元素开始往左或往上
		        while i >=0 and j < m:
		            # 元素较大，往上走
		            if array[i][j] > target:
		                i -= 1
		            # 元素较小，往右走
		            elif array[i][j] < target:
		                j += 1
		            else:
		                return True
		        return False
	*/
	if len(array) == 0 || len(array[0]) == 0 {
		return false
	}
	n := len(array)
	m := len(array[0])
	i := n - 1
	j := 0
	//从最左下角的元素开始往左或往上
	for i >= 0 && j < m {
		if array[i][j] > target {
			i--
		} else if array[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}
