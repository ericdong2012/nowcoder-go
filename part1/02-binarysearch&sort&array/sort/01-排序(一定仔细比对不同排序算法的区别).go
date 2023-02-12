package main

import "fmt"

/*
NC140 排序
https://www.nowcoder.com/practice/2baf799ea0594abd974d37139de27896?tpId=117&tqId=37851&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590&title=

描述
给定一个长度为 n 的数组，请你编写一个函数，返回该数组按升序排序后的结果。

注：本题数据范围允许绝大部分排序算法，请尝试多种排序算法的实现。


示例1
输入：
[5,2,3,1,4]
返回值：
[1,2,3,4,5]

示例2
输入：
[5,1,6,2,5]
返回值：
[1,2,5,5,6]

*/

/*
冒泡： 相邻的两个元素进行比较，然后把较大的元素放到后面（正向排序），在一轮比较完后最大的元素就放在了最后一个位置，因为这一点像鱼儿在水中吐的气泡在上升的过程中不断变大，所以得名冒泡排序
class Solution:
    def MySort(self , arr: List[int]) -> List[int]:
        # write code here
        for i in range(len(arr)):
            for j in range(i+1,len(arr)):
                if arr[j]<arr[i]:
                    arr[i],arr[j]=arr[j],arr[i]
        return arr

选择： 第一轮的时候，所有的元素都和第一个元素进行比较，如果比第一个元素小，就和第一个元素进行交换，在这轮比较完后，就找到了最小的元素；第二轮的时候所有的元素都和第二个元素进行比较找出第二个位置的元素，以此类推
class Solution:
    def MySort(self , arr: List[int]) -> List[int]:
        # write code here
        #选择排序
        for i in range(len(arr)):
            min_index=i
            for j in range(i+1,len(arr)):
                if arr[j]<arr[min_index]:
                    min_index=j
            arr[i],arr[min_index]=arr[min_index],arr[i]
        return arr

快排：
时间复杂度：O（nlogn）空间复杂度：O（1）
不稳定

1.首先data列表传进来
2.递归结束条件,递归传进来的data为空,结束递归.意思是当取最后一个元素当做基准值之后,left=[],right=[],也就是base左边和右边都没有元素了
3.选取基准值
4.基准值左边排序
5.基准值右边排序
6.返回left+[base]+right

选基准
分区
递归

class Solution:
    def MySort(self , arr ):
        # write code here
        self.quick_sort(0, len(arr)-1, arr)
        return arr

    def quick_sort(self, start, end, arr):
        if start >= end:
            return

        left, right = start, end
		// 定基准
        pivot = arr[(start + end)//2]

		// 分区
        while left <= right:
			// 顺序找左边比标准小，右边比标准大，不符合条件跳出
            while left <= right and arr[left] < pivot:
                left += 1
            while left <= right and arr[right] > pivot:
                right -= 1

            // 找到第一个需要交换的元素，左边比标准大， 右边比标准小，交换
            if left <= right:
                arr[left], arr[right] = arr[right], arr[left]
                left += 1
                right -= 1

		// 递归
        self.quick_sort(start, right, arr)
        self.quick_sort(left, end, arr)


插入：
class Solution:
    def MySort(self , arr: List[int]) -> List[int]:
        # write code here
        #插入排序
        for i in range(1,len(arr)):
            pre_index=i-1
            cur=arr[i]
            while pre_index>=0 and arr[pre_index]>cur:
                arr[pre_index + 1] = arr[pre_index]
                pre_index-=1
            arr[pre_index + 1] = cur;
        return arr

希尔：
class Solution:
    def MySort(self , arr: List[int]) -> List[int]:
        # write code here
        #希尔排序
        import math
        gap=1
        while(gap < len(arr)/3):
            gap = gap*3+1
        while gap > 0:
            for i in range(gap,len(arr)):
                temp = arr[i]
                j = i-gap
                while j >=0 and arr[j] > temp:
                    arr[j+gap]=arr[j]
                    j-=gap
                arr[j+gap] = temp
            gap = math.floor(gap/3)
        return arr

归并：
class Solution:
    def MySort(self , arr: List[int]) -> List[int]:
        # write code here
        #归并排序
        import math
        if len(arr)<2:
            return arr
        mid=math.floor(len(arr)/2)
        left,right=arr[0:mid],arr[mid:]
        return merge(self.MySort(left),self.MySort(right))

def merge(left,right):
    result=[]
    while left and right:
        if left[0]<right[0]:
            result.append(left.pop(0))
        else:
            result.append(right.pop(0))
    while left:
        result.append(left.pop(0))
    while right:
        result.append(right.pop(0))
    return result

堆：

*/

func MySort(arr []int) []int {
	// write code here
	// 冒泡
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			// 后面的比前面小， 升序
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr

	// 选择
	//for i := 0; i < len(arr); i++ {
	//	min_index := i
	//	for j := i + 1; j < len(arr); j++ {
	//		// 找最小元素的索引
	//		if arr[j] < arr[min_index] {
	//			min_index = j
	//		}
	//	}
	//	arr[i], arr[min_index] = arr[min_index], arr[i]
	//}
	//return arr

	// 快排
	//quick_sort(0, len(arr)-1, arr)
	//return arr

}

//func quick_sort(start, end int, arr []int) {
//	if start >= end {
//		return
//	}
//
//	left, right := start, end
//	pivot := arr[(start+end)/2]
//	for left <= right {
//		for left <= right && arr[left] <= pivot {
//			left++
//		}
//
//		for left <= right && arr[right] >= pivot {
//			right--
//		}
//
//		//arr[left], arr[right] = arr[right], arr[left]
//		if left <= right {
//			arr[left], arr[right] = arr[right], arr[left]
//			//left++
//			//right--
//		}
//
//	}
//	// 写法上没有问题，但是实际上 left, end 值已经发生了变化
//	quick_sort(start, right, arr)
//	quick_sort(left, end, arr)
//}

func main() {
	sort := MySort([]int{5, 2, 3, 1, 4})
	fmt.Println(sort)
}
