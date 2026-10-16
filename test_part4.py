"""
part4 算法基础 (Python 版) 的验证脚本。

用法：在仓库根目录执行
    python test_part4.py

覆盖范围：
  - 18 个带 `if __name__ == "__main__"` 的文件：直接子进程运行，检查是否正常退出；
  - 16 个仅定义函数、没有入口的文件：importlib 加载后用题目样例调用其函数并断言结果。

全部通过时末尾输出 ALL GOOD。
"""
import glob, importlib.util, subprocess, sys

base = glob.glob('part4_*/')[0]
fails = []
main_count = 0


def load(rel_path, mod_name):
    spec = importlib.util.spec_from_file_location(mod_name, base + rel_path)
    m = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(m)
    return m


def ll_to_list(node):
    out = []
    while node is not None:
        out.append(node.val)
        node = node.next
    return out


def check(label, got, exp):
    if got != exp:
        fails.append(f'{label}: got={got!r} exp={exp!r}')


# === Part A: run all files that have a __main__ block ===
main_files = [
    '01-array/01-移动 0.py', '01-array/02-盛水最多的容器.py', '01-array/03-跳台阶.py',
    '01-array/05-合并两个有序数组.py', '02-linkedlist/01-反转链表.py', '02-linkedlist/02-判断链表中是否有环.py',
    '03-hash/02-异位词分组.py', '03-hash/05-三数之和.py', '04-stack/01-有效括号序列.py',
    '04-stack/03-最长的括号子串.py', '04-stack/04-滑动窗口的最大值.py',
    '05-sort/binary search/01-求平方根.py', '05-sort/binary search/02-二分查找-II.py',
    '05-sort/doublepoint/01-移动 0.py', '05-sort/doublepoint/02-盛水最多的容器.py',
    '06-string/02-最长的括号子串.py', '06-string/08-有效括号序列.py', '06-string/10-最长重复子串.py',
]
for f in main_files:
    r = subprocess.run([sys.executable, base + f], capture_output=True, text=True)
    if r.returncode != 0:
        fails.append(f'RUNTIME ERROR {f}: {r.stderr.strip()}')
    else:
        main_count += 1

# === Part B: logic tests for files without main ===
m = load('01-array/04-加1.py', 'a04')
check('plus_one [1,2,3]', m.plus_one([1, 2, 3]), [1, 2, 4])
check('plus_one [9,9]', m.plus_one([9, 9]), [1, 0, 0])

m = load('02-linkedlist/03-链表内指定区间反转.py', 'l03')
LN = m.ListNode
check('reverseBetween 2,4', ll_to_list(m.reverse_between(LN(1, LN(2, LN(3, LN(4, LN(5))))), 2, 4)), [1, 4, 3, 2, 5])

m = load('02-linkedlist/04-链表中环的入口节点.py', 'l04')
LN = m.ListNode
n1 = LN(1); n2 = LN(2); n3 = LN(3); n4 = LN(4); n5 = LN(5)
n1.next = n2; n2.next = n3; n3.next = n4; n4.next = n5; n5.next = n3
check('entryNodeOfLoop', m.entry_node_of_loop(n1).val, 3)

m = load('03-hash/01-字母异位词.py', 'h01')
check('isAnagram +', m.is_anagram('anagram', 'nagaram'), True)
check('isAnagram -', m.is_anagram('rat', 'car'), False)

m = load('03-hash/04-两数之和.py', 'h04')
check('twoSum', m.two_sum([2, 7, 11, 15], 9), [0, 1])

m = load('04-stack/02-最小栈.py', 's02')
m.push(-1); m.push(2); o = [m.min(), m.top()]; m.pop(); m.push(1); o += [m.top(), m.min()]
check('minStack', o, [-1, 2, 1, -1])

m = load('04-stack/05-接雨水问题(盛水最多的容器类似).py', 's05')
check('maxWater', m.max_water([3, 1, 2, 5, 2, 4]), 5)

m = load('05-sort/04-相对排序(跟合并2个有序数组对比).py', 'so04')
check('relativeSort', m.relative_sort_array([2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19], [2, 1, 4, 3, 9, 6]),
      [2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19])

m = load('05-sort/05-合并区间.py', 'so05')
Iv = m.Interval
res = m.merge([Iv(10, 30), Iv(20, 60), Iv(80, 100), Iv(150, 180)])
check('mergeIntervals', [(x.start, x.end) for x in res], [(10, 60), (80, 100), (150, 180)])

m = load('06-string/01_first-unique-character-in-a-string.py', 'st01')
check('firstUniqChar leetcode', m.first_uniq_char('leetcode'), 0)
check('firstUniqChar loveleetcode', m.first_uniq_char('loveleetcode'), 2)

m = load('06-string/03-反转字符串.py', 'st03')
check('reverse', m.solve('abcd'), 'dcba')

m = load('06-string/04-判断是否为回文字符串.py', 'st04')
check('judge +', m.judge('absba'), True)
check('judge -', m.judge('ranko'), False)

m = load('06-string/05-找到字符串中的所有字母异位词.py', 'st05')
check('findWord', m.find_word('abab', 'ab'), [0, 1, 2])

m = load('06-string/06-最长回文子串.py', 'st06')
check('palindrome ababc', m.get_longest_palindrome1('ababc'), 3)
check('palindrome abbba', m.get_longest_palindrome1('abbba'), 5)

m = load('06-string/07-最长公共子序列.py', 'st07')
check('LCS', m.longest_common_subsequence('abcde', 'ace'), 3)

m = load('06-string/09-最长公共前缀.py', 'st09')
check('LCP', m.longest_common_prefix(['abca', 'abc', 'abca', 'abc', 'abcc']), 'abc')

m = load('06-string/10-最长重复子串.py', 'st10')
check('repeatSubstr ababc=4', m.solve2('ababc'), 4)
check('repeatSubstr abcab=0', m.solve2('abcab'), 0)
check('repeatSubstr single=0', m.solve2('a'), 0)

print(f'main-block files ran OK: {main_count}/{len(main_files)}')
print(f'logic failures: {len(fails)}')
for fl in fails:
    print('  FAIL', fl)
print('ALL GOOD' if not fails else 'HAS FAILURES')
