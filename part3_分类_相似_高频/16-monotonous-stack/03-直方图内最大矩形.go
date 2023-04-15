package _6_monotonous_stack


/*
NC171 直方图内最大矩形
https://www.nowcoder.com/practice/bfaac4eebe5947af80912d1bcfcf2baa?tpId=117&tqId=39316&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=1579&title=


描述
给定一个数组heights，长度为n，height[i]是在第i点的高度，那么height[i]表示的直方图，能够形成的最大矩形是多少?
1.每个直方图宽度都为1
2.直方图都是相邻的
3.如果不能形成矩形，返回0即可
4.保证返回的结果不会超过231-1

数据范围:
0 <= heights[i] <= 10^40<=heights[i]<=10
4

0 <= heights.length <=10^50<=heights.length<=10
5


如输入[3,4,7,8,1,2]，那么如下:


示例1
输入：
[3,4,7,8,1,2]
复制
返回值：
14
复制
示例2
输入：
[1,7,3,2,4,5,8,2,7]
复制
返回值：
16


*/
