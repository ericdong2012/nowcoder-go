"""
part5 算法思维 (Python 版) 的验收脚本。

用法：在仓库根目录执行
    python test_part5.py

三道检查：
  Part A：import 全部 117 个 .py 模块，捕获定义期 / 导入期错误。
  Part B：把每个带 `if __name__ == "__main__"` 守卫的文件作为子进程跑一遍，捕获运行期错误。
  Part C：对每个目录挑代表性文件做逻辑断言（含子 agent 修复过 bug 的文件）。
  Part D：给原先既没有 __main__ 演示、也没被 Part C 覆盖的文件补样例断言
          （15-单调栈/01-接雨水问题.py 是空占位文件，无可测逻辑，不在其中）。
  Part E：part5 中与 part4 同名/同题、但自身从未被任何输入执行过的 14 个副本文件。

全部通过时末尾输出 ALL GOOD。
"""
import glob, importlib.util, subprocess, sys, os

BASE = glob.glob('part5_*/')[0]
ALL_FILES = sorted(glob.glob('part5_*/**/*.py', recursive=True))
fails = []


def check(label, got, exp):
    if got != exp:
        fails.append(f'{label}: got={got!r} exp={exp!r}')


def load(rel_path, mod_name):
    spec = importlib.util.spec_from_file_location(mod_name, BASE + rel_path)
    m = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(m)
    return m


def ll_to_list(node):
    out = []
    while node is not None:
        out.append(node.val)
        node = node.next
    return out


# ============================================================
# Part A: import every module
# ============================================================
import_errors = []
for idx, f in enumerate(ALL_FILES):
    try:
        spec = importlib.util.spec_from_file_location(f'm{idx}', f)
        m = importlib.util.module_from_spec(spec)
        spec.loader.exec_module(m)
    except Exception as e:
        import_errors.append((os.path.relpath(f, BASE), repr(e)))
print(f'Part A: imported {len(ALL_FILES) - len(import_errors)}/{len(ALL_FILES)} modules OK')
for path, err in import_errors:
    print('  IMPORT FAIL', path, err)
    fails.append(f'IMPORT {path}: {err}')


# ============================================================
# Part B: run every file that has a __main__ guard
# ============================================================
main_targets = []
for f in ALL_FILES:
    txt = open(f, encoding='utf-8').read()
    if '__name__' in txt and '__main__' in txt:
        main_targets.append(f)
main_ok = 0
for f in main_targets:
    r = subprocess.run([sys.executable, f], capture_output=True, text=True, timeout=60)
    if r.returncode != 0:
        fails.append(f'RUNTIME {os.path.relpath(f, BASE)}: {r.stderr.strip()[:200]}')
    else:
        main_ok += 1
print(f'Part B: ran {main_ok}/{len(main_targets)} __main__ files OK')


# ============================================================
# Part C: curated logic assertions
# ============================================================
# 01-binarysearch
m = load('01-binarysearch/03-丑数.py', 'b03')
check('ugly(7)', m.get_ugly_number(7), 8)
check('ugly(10)', m.get_ugly_number(10), 12)
m = load('01-binarysearch/05-调整数组顺序使奇数位于偶数前面(一)(e).py', 'b05')
check('reOrder [1,2,3,4]', m.re_order_array([1, 2, 3, 4]), [1, 3, 2, 4])

# 02-hash
m = load('02-hash/10-和为K的连续子数组.py', 'h10')
check('maxlenEqualK', m.maxlen_equal_k([1, -2, 1, 1, 1], 0), 3)
m = load('02-hash/11-最长不含重复字符的子字符串.py', 'h11')
check('longestSubstring abcabcbb', m.length_of_longest_substring('abcabcbb'), 3)
m = load('02-hash/08-缺失的第一个正整数.py', 'h08')
check('firstMissingPositive', m.min_number_disappeared([-2, 3, 4, 1, 5]), 2)
m = load('02-hash/09-设计LRU缓存结构.py', 'h09')
sol = m.Solution(2)
sol.set(1, 1); sol.set(2, 2)
check('LRU get1', sol.get(1), 1)
sol.set(3, 3)  # evicts 2
check('LRU get2 evicted', sol.get(2), -1)
sol.set(4, 4)  # evicts 1
check('LRU get1 evicted', sol.get(1), -1)
check('LRU get3', sol.get(3), 3)
check('LRU get4', sol.get(4), 4)

# 03-doublepoint
m = load('03-doublepoint/05-长度最小的连续子数组.py', 'd05')
check('minSubarray', m.min_subarray([1, 2, 4, 4, 1, 1, 1], 9), 3)
m = load('03-doublepoint/06-比较版本号.py', 'd06')
check('compare 1.1/2.1', m.compare('1.1', '2.1'), -1)
check('compare 1.1/1.01', m.compare('1.1', '1.01'), 0)
check('compare 1.1/1.1.1', m.compare('1.1', '1.1.1'), -1)

# 04-stack
m = load('04-stack/06-素数.py', 's06')
check('isPrime(7)', m.is_prime(7), True)
check('isPrime(9)', m.is_prime(9), False)
check('findPrimes(11)', m.find_primes(11), [2, 3, 5, 7, 11])

# 05-前缀和
m = load('05-前缀和/02-除自身以外数组的乘积.py', 'p02')
check('timesExceptSelf', m.times_except_self([1, 2, 3, 4]), [24, 12, 8, 6])
m = load('05-前缀和/03-和大于等于k的最短子数组.py', 'p03')
check('shortestSubarray', m.shortest_subarray([2, 1, 2, 3], 5), 2)
m = load('05-前缀和/01-长度最小的连续子数组.py', 'p01')
check('minSubarray ok', m.min_subarray([1, 2, 4, 4, 1, 1, 1], 9), 3)
check('minSubarray none->0', m.min_subarray([1, 1, 1], 100), 0)  # agent-fixed: 找不到合法子数组时返回 0
check('minSubarray single', m.min_subarray([1, 4, 4], 4), 1)

# 06-string
m = load('06-string/11-大数加法.py', 'st11')
check('bigAdd 1+99', m.solve1('1', '99'), '100')
m = load('06-string/12-大数乘法.py', 'st12')
check('bigMul 11*99', m.solve3('11', '99'), '1089')
check('bigMul 1*0', m.solve3('1', '0'), '0')
m = load('06-string/13-把字符串转换成整数(atoi).py', 'st13')
check('atoi "  -42abc"', m.str_to_int('   -42abc'), -42)
check('atoi "4193 with words"', m.str_to_int('4193 with words'), 4193)
check('atoi overflow', m.str_to_int('91283472332'), 2**31 - 1)

# 08-sort (agent-fixed: 最小的k个数 guard && -> ||)
m = load('08-sort/05-最小的k个数.py', 'so05')
fns = [n for n in dir(m) if not n.startswith('_') and callable(getattr(m, n))]
# call the first list-returning function with sample input
fn = getattr(m, fns[0]) if fns else None
if fn:
    try:
        out = fn([4, 5, 1, 6, 2, 7, 3, 8], 4)
        check('leastK(4)', sorted(out), [1, 2, 3, 4])
    except Exception as e:
        fails.append(f'leastK call: {e!r}')

# 10-dp
m = load('10-dp/01-斐波那契数列.py', 'dp01')
fns = [n for n in dir(m) if not n.startswith('_') and callable(getattr(m, n))]
fn = getattr(m, fns[0]) if fns else None
if fn:
    try:
        check('fib(5)', fn(5), 5)
    except Exception as e:
        fails.append(f'fib call: {e!r}')

# 10-dp / 19-把数字翻译成字符串 (agent-fixed: 末位无合法译码时返回 0，原 KeyError)
m = load('10-dp/19-把数字翻译成字符串.py', 'dp19')
check('decode "12"', m.solve('12'), 2)
check('decode "226"', m.solve('226'), 3)
check('decode "30"->0', m.solve('30'), 0)
check('decode big', m.solve('31717126241541717'), 192)

# 16-位运算 (agent-fixed: 只出现一次的数字(二) int32 overflow)
m = load('16-位运算/03-只出现一次的数字(二).py', 'bi03')
fns = [n for n in dir(m) if not n.startswith('_') and callable(getattr(m, n))]
fn = getattr(m, fns[0]) if fns else None
if fn:
    try:
        # [2,2,3,2] => 3 ; 其余元素各出现 3 次，-4 出现一次 => -4
        check('singleNumberII(2,2,3,2)', fn([2, 2, 3, 2]), 3)
        check('singleNumberII(-4)', fn([-2, -2, -2, 1, 1, 1, 4, 4, 4, -4]), -4)
    except Exception as e:
        fails.append(f'singleNumberII call: {e!r}')

# 18-linkedlist
m = load('18-linkedlist/01-反转链表.py', 'l01')
LN = m.ListNode
head = LN(1, LN(2, LN(3, LN(4, LN(5)))))
check('reverseList', ll_to_list(m.reverse_list(head) if hasattr(m, 'reverse_list') else m.ReverseList(head)), [5, 4, 3, 2, 1])

# ============================================================
# Part D: 给原先没有任何执行记录的文件补样例断言
# ============================================================
# 01-binarysearch/04-在两个长度相等的排序数组中找到上中位数
m = load('01-binarysearch/04-在两个长度相等的排序数组中找到上中位数.py', 'd04')
check('upperMedian sample', m.find_median_in_two_sorted_array([1, 2, 3, 4], [3, 4, 5, 6]), 3)
check('upperMedian [1,2]/[3,4]', m.find_median_in_two_sorted_array([1, 2], [3, 4]), 2)

# 04-stack/05-接雨水问题（part5 副本的函数此前未被断言过）
m = load('04-stack/05-接雨水问题(盛水最多的容器类似).py', 'd05r')
check('maxWater [3,1,2,5,2,4]', m.max_water([3, 1, 2, 5, 2, 4]), 5)
check('maxWater [4,5,1,3,2]', m.max_water([4, 5, 1, 3, 2]), 2)

# 04-stack/07-用两个栈实现队列（模块级栈，样例 ["PSH1","PSH2","POP","POP"] -> 1,2）
m = load('04-stack/07-用两个栈实现队列.py', 'd07')
m.push1(1); m.push1(2)
check('twoStackQueue pop', [m.pop1(), m.pop1()], [1, 2])

# 09-tree 辅助函数
def _pre(t):
    return [] if t is None else [t.val] + _pre(t.left) + _pre(t.right)

def _in(t):
    return [] if t is None else _in(t.left) + [t.val] + _in(t.right)

# 09-tree/07-重建二叉树：重建后前序+中序应与输入一致
m = load('09-dfs&bfs&记忆化搜索&树/07-重建二叉树.py', 'd07t')
pre, vin = [1, 2, 4, 7, 3, 5, 6, 8], [4, 7, 2, 1, 5, 3, 8, 6]
t = m.re_construct_binary_tree(pre, vin)
check('rebuildTree preorder', _pre(t), pre)
check('rebuildTree inorder', _in(t), vin)

# 09-tree/09-二叉搜索树的第k个节点
m = load('09-dfs&bfs&记忆化搜索&树/09-二叉搜索树的第k个节点.py', 'd09')
TN = m.TreeNode
bst = TN(5, TN(3, TN(2), TN(4)), TN(7, TN(6), TN(8)))
check('kthNode k=3', m.kth_node(bst, 3), 4)
check('kthNode k=1', m.kth_node(bst, 1), 2)
check('kthNode k>n -> -1', m.kth_node(bst, 8), -1)
check('kthNode empty -> -1', m.kth_node(None, 1), -1)

# 09-tree/12-将升序数组转化为平衡二叉搜索树：中序有序 + 高度平衡
m = load('09-dfs&bfs&记忆化搜索&树/12-将升序数组转化为平衡二叉搜索树.py', 'd12')
t = m.sorted_array_to_bst([-1, 0, 1, 2])
check('arrToBST inorder', _in(t), [-1, 0, 1, 2])

def _hb(t):  # (是否平衡, 高度)
    if t is None:
        return True, 0
    lok, lh = _hb(t.left)
    rok, rh = _hb(t.right)
    return lok and rok and abs(lh - rh) <= 1, max(lh, rh) + 1

check('arrToBST balanced', _hb(t)[0], True)

# 09-tree/13-在二叉树中找到两个节点的最近公共祖先
m = load('09-dfs&bfs&记忆化搜索&树/13-在二叉树中找到两个节点的最近公共祖先.py', 'd13')
TN = m.TreeNode
bt = TN(3, TN(5, TN(6), TN(2, TN(7), TN(4))), TN(1, TN(0), TN(8)))
check('LCA(5,1)=3', m.lowest_common_ancestor2(bt, 5, 1), 3)
check('LCA(2,7)=2', m.lowest_common_ancestor2(bt, 2, 7), 2)

# 10-dp/15-正则表达式匹配
m = load('10-dp/15-正则表达式匹配.py', 'd15')
check('regex aaa/a*a', m.match2('aaa', 'a*a'), True)
check('regex aad/c*a*d', m.match2('aad', 'c*a*d'), True)
check('regex a/.*', m.match2('a', '.*'), True)
check('regex aaa/ab*a', m.match2('aaa', 'ab*a'), False)
check('regex aaa/aa.a', m.match2('aaa', 'aa.a'), False)

# 18-linkedlist/08-给单链表加一
m = load('18-linkedlist/08-给单链表加一.py', 'd08')
LN = m.ListNode

def _ll(vals):
    head = None
    for v in reversed(vals):
        head = LN(v, head)
    return head

check('plusOne 123', ll_to_list(m.plus_one(_ll([1, 2, 3]))), [1, 2, 4])
check('plusOne 120', ll_to_list(m.plus_one(_ll([1, 2, 0]))), [1, 2, 1])
check('plusOne 9', ll_to_list(m.plus_one(_ll([9]))), [1, 0])
check('plusOne 999', ll_to_list(m.plus_one(_ll([9, 9, 9]))), [1, 0, 0, 0])

# 06-string/10-最长重复子串（含修复后的边界：长度<2 返回 0）
m = load('06-string/10-最长重复子串.py', 'd10')
check('repeatSubstr ababc=4', m.solve2('ababc'), 4)
check('repeatSubstr abcab=0', m.solve2('abcab'), 0)
check('repeatSubstr single=0', m.solve2('a'), 0)

# ============================================================
# Part E: part5 中与 part4 同名/同题、但自身从未被任何输入执行过的副本文件
# ============================================================
m = load('02-hash/01-字母异位词.py', 'e01')
check('isAnagram +', m.is_anagram('anagram', 'nagaram'), True)
check('isAnagram -', m.is_anagram('rat', 'car'), False)

m = load('02-hash/02-异位词分组.py', 'e02')
check('groupAnagrams', m.group_anagrams(["eat", "tea", "tan", "ate", "nat", "bat"]),
      [['eat', 'tea', 'ate'], ['tan', 'nat'], ['bat']])

m = load('02-hash/04-两数之和.py', 'e04')
check('twoSum', m.two_sum([2, 7, 11, 15], 9), [0, 1])

m = load('04-stack/02-最小栈.py', 'e05')
m.push(-1); m.push(2); o = [m.min(), m.top()]; m.pop(); m.push(1); o += [m.top(), m.min()]
check('minStack', o, [-1, 2, 1, -1])

m = load('06-string/01_first-unique-character-in-a-string.py', 'e06')
check('firstUniqChar leetcode', m.first_uniq_char('leetcode'), 0)
check('firstUniqChar loveleetcode', m.first_uniq_char('loveleetcode'), 2)

m = load('06-string/03-反转字符串.py', 'e07')
check('reverse str', m.solve('abcd'), 'dcba')

m = load('06-string/04-判断是否为回文字符串.py', 'e08')
check('judge +', m.judge('absba'), True)
check('judge -', m.judge('ranko'), False)

m = load('06-string/05-找到字符串中的所有字母异位词.py', 'e09')
check('findWord', m.find_word('abab', 'ab'), [0, 1, 2])

m = load('06-string/06-最长回文子串.py', 'e10')
check('palindrome ababc', m.get_longest_palindrome1('ababc'), 3)
check('palindrome abbba', m.get_longest_palindrome1('abbba'), 5)

m = load('06-string/07-最长公共子序列.py', 'e11')
check('LCS', m.longest_common_subsequence('abcde', 'ace'), 3)

m = load('06-string/09-最长公共前缀.py', 'e12')
check('LCP', m.longest_common_prefix(['abca', 'abc', 'abca', 'abc', 'abcc']), 'abc')

m = load('10-dp/07-最长公共子序列.py', 'e13')
check('LCS dp', m.longestCommonSubsequence('abcde', 'ace'), 3)

m = load('18-linkedlist/03-链表内指定区间反转.py', 'e14')
LN = m.ListNode
check('reverseBetween 2,4', ll_to_list(m.reverse_between(LN(1, LN(2, LN(3, LN(4, LN(5))))), 2, 4)), [1, 4, 3, 2, 5])

m = load('18-linkedlist/04-链表中环的入口节点.py', 'e15')
LN = m.ListNode
n1 = LN(1); n2 = LN(2); n3 = LN(3); n4 = LN(4); n5 = LN(5)
n1.next = n2; n2.next = n3; n3.next = n4; n4.next = n5; n5.next = n3
check('entryNodeOfLoop', m.entry_node_of_loop(n1).val, 3)

# ============================================================
print(f'\nPart C/D/E: curated logic failures: {len(fails)}')
for fl in fails:
    print('  FAIL', fl)
print('ALL GOOD' if not fails else 'HAS FAILURES')
