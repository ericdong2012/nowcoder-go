package main

import (
	"fmt"
	"strconv"
)

/*
NC10 大数乘法
https://www.nowcoder.com/practice/c4c488d4d40d4c4e9824c3650f7d5571?tpId=117&tqId=37843&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=579&title=

描述
以字符串的形式读入两个数字，编写一个函数计算它们的乘积，以字符串形式返回。
要求：空间复杂度 O(n)，时间复杂度 O(n^2)

示例1
输入：
"11","99"
返回值：
"1089"
说明：
11*99=1089

示例2
输入：
"1", "0"
返回值：
"0"

*/

/*

遍历法， 一个索引从头开始，一个索引从后开始， 将乘法转换成加法
库函数

*/

/*
参考代码1
class Solution {
public:
	string solve(string s, string t) {
	// write code here
	if(s.empty() == true || t.size()== true)
		return "0";
	int len1 = s.size();
	int len2 = t.size();
	string res(len1 + len2,'0');           //设置初始变量"0000"
	for(int i = len1 - 1; i >= 0; i--) {   //从最后一位开始乘
		for(int j = len2 - 1; j >= 0; j--){
			int temp = (s[i] - '0') * (t[j] - '0') + (res[i+j+1] - '0');//计算两个数乘之后加进位的运算
			res[i+j+1] = temp % 10 + '0';   //对10取余
			res[i+j] = res[i+j] + temp/10;  //进位
			}
	}
	for(int i = 0; i < res.size();i++){
		if(res[i] != '0')
			return res.substr(i);//从指定位置开始，并具有指定的长度。
	}
    return res;
	}
};


*/

/*
参考代码2：
import java.util.*;
public class Solution {
    public String solve (String s, String t) {
        //将字符串转化为数组形式
        int lenS = s.length(), lenT = t.length();
        int[] arrS = new int[lenS];
        int[] arrT = new int[lenT];
        for(int i = 0; i < lenS; i++) {
            arrS[i] = s.charAt(i) - '0';
        }
        for(int i = 0; i < lenT; i++) {
            arrT[i] = t.charAt(i) - '0';
        }
        // 存放结果
        int[] res = new int[lenS + lenT];
        // 中间 i * j 会有重复出现，如89 * 12, 会有res = [0, 8*1, 8*2+9*1, 9*2]
        for(int i = 0; i < lenS; i++) {
            for(int j = 0; j < lenT; j++) {
                // i + j + 1 是为了给进位留一个位置
                res[i + j + 1] += arrS[i] * arrT[j];
            }
        }
        // 进位的值
        int carry = 0;
        // 从最低位，也就是数组res最右边元素开始处理进位
        for(int i = lenS + lenT - 1; i >= 0; i--) {
            res[i] += carry;
            carry = res[i] / 10;
            res[i] = res[i] % 10;
        }
        StringBuilder ans = new StringBuilder();
        // 表示当前位在数组的索引，目的是找出高位的0的位置，前置的0需要忽略掉
        int cur = 0;
        while(cur < lenS + lenT && res[cur] == 0) {
            cur++;
        }
        // 收集结果
        for(int i = cur; i < res.length; i++) {
            ans.append(res[i]);
        }
        return ans.length() == 0 ? "0" : ans.toString();
    }
}

*/

// 解法1：
func solve3(s string, t string) string {
	/*
		将两个字符串转化为整型数组形式
		逐位相乘, 中间 i * j 会有重复出现(索引1，0 和 0， 1),  所以res = [0, 8*1, 8*2+9*1, 9*2]
		从最低位, 也就是数组res最右边， 从后往前， 开始处理进位
	    处理先导0等边界情况
		最后收集结果
	*/
	numS, numT := []byte(s), []byte(t)
	tempRes := make([]int, len(s)+len(t))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			// 主要是考虑到进位，在头部加了一位
			tempRes[i+j+1] += int(numS[i]-'0') * int(numT[j]-'0')
		}
	}
	// 从后往前，处理进位
	for i := len(tempRes) - 1; i > 0; i-- {
		if tempRes[i] >= 10 {
			// 前一位加上进位，也就是整除的数
			tempRes[i-1] += tempRes[i] / 10
			// 当前位保留余数
			tempRes[i] %= 10
		}
	}
	// 找到第一个不是0的索引，用k记录，目的是去除前导0
	k := 0
	//for ; k < len(tempRes); k++ {
	//	if tempRes[k] != 0 {
	//		break
	//	}
	//}
	for _, v := range tempRes {
		if v !=0 {
			break
		}
		k++
	}
	// 全是0
	if k == len(tempRes) {
		return "0"
	}
	// k 之后是不为0的地方
	tempRes = tempRes[k:]
	// 构建字符串， rune 和 string 转换相对简单
	//res := make([]rune, len(tempRes))
	res := ""
	for _, n := range tempRes {
		//res[i] = rune(n + '0')
		res += strconv.Itoa(n)
	}

	return res
}


func main() {
	s := solve3("1", "1")
	fmt.Println(s)
}
