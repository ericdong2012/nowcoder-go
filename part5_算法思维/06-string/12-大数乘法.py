"""
NC10 大数乘法
https://www.nowcoder.com/practice/c4c488d4d40d4c4e9824c3650f7d5571

描述
以字符串的形式读入两个数字，编写一个函数计算它们的乘积，以字符串形式返回。
要求：空间复杂度 O(n)，时间复杂度 O(n^2)

示例1
输入："11","99"
返回值："1089"
说明：11*99=1089

示例2
输入："1", "0"
返回值："0"
"""

# 遍历法，一个索引从头开始，一个索引从后开始，将乘法转换成加法
# 库函数

"""
参考代码1
class Solution {
public:
    string solve(string s, string t) {
        if(s.empty() == true || t.size()== true)
            return "0";
        int len1 = s.size();
        int len2 = t.size();
        string res(len1 + len2, '0');           // 设置初始变量 "0000"
        for(int i = len1 - 1; i >= 0; i--) {    // 从最后一位开始乘
            for(int j = len2 - 1; j >= 0; j--){
                int temp = (s[i] - '0') * (t[j] - '0') + (res[i+j+1] - '0'); // 计算两个数乘之后加进位的运算
                res[i+j+1] = temp % 10 + '0';   // 对 10 取余
                res[i+j] = res[i+j] + temp/10;  // 进位
            }
        }
        for(int i = 0; i < res.size(); i++){
            if(res[i] != '0')
                return res.substr(i);
        }
        return res;
    }
};
"""

"""
参考代码2：
import java.util.*;
public class Solution {
    public String solve (String s, String t) {
        int lenS = s.length(), lenT = t.length();
        int[] arrS = new int[lenS];
        int[] arrT = new int[lenT];
        for(int i = 0; i < lenS; i++) arrS[i] = s.charAt(i) - '0';
        for(int i = 0; i < lenT; i++) arrT[i] = t.charAt(i) - '0';
        int[] res = new int[lenS + lenT];
        for(int i = 0; i < lenS; i++) {
            for(int j = 0; j < lenT; j++) {
                res[i + j + 1] += arrS[i] * arrT[j];
            }
        }
        int carry = 0;
        for(int i = lenS + lenT - 1; i >= 0; i--) {
            res[i] += carry;
            carry = res[i] / 10;
            res[i] = res[i] % 10;
        }
        StringBuilder ans = new StringBuilder();
        int cur = 0;
        while(cur < lenS + lenT && res[cur] == 0) cur++;
        for(int i = cur; i < res.length; i++) ans.append(res[i]);
        return ans.length() == 0 ? "0" : ans.toString();
    }
}
"""


# 解法1：
def solve3(s, t):
    """
    将两个字符串转化为整型数组形式
    逐位相乘，中间 i * j 会有重复出现(索引 1,0 和 0,1)，所以 res = [0, 8*1, 8*2+9*1, 9*2]
    从最低位，也就是数组 res 最右边，从后往前，开始处理进位
    处理先导 0 等边界情况
    最后收集结果
    """
    num_s = [ord(c) - ord('0') for c in s]
    num_t = [ord(c) - ord('0') for c in t]
    temp_res = [0] * (len(s) + len(t))
    for i in range(len(s)):
        for j in range(len(t)):
            # 主要是考虑到进位，在头部加了一位
            temp_res[i + j + 1] += num_s[i] * num_t[j]
    # 从后往前，处理进位
    for i in range(len(temp_res) - 1, 0, -1):
        if temp_res[i] >= 10:
            # 前一位加上进位，也就是整除的数
            temp_res[i - 1] += temp_res[i] // 10
            # 当前位保留余数
            temp_res[i] %= 10
    # 找到第一个是 0 的索引，用 k 记录，目的是去除前导 0
    k = 0
    # for k in range(len(temp_res)):
    #     if temp_res[k] != 0:
    #         break
    for v in temp_res:
        if v != 0:
            break
        k += 1
    # 全是 0
    if k == len(temp_res):
        return "0"
    # k 之后是不为 0 的地方
    temp_res = temp_res[k:]
    # 构建字符串
    res = ""
    for n in temp_res:
        res += str(n)

    return res


if __name__ == "__main__":
    s = solve3("1", "1")
    print(s)
