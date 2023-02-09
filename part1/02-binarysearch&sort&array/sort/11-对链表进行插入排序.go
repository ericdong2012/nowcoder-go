package main

/*
NC244 对链表进行插入排序
https://www.nowcoder.com/practice/cc6c61215dfb446f8eccea3663e3d8db?tpId=117&tqId=39589&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590&title=

描述
对链表进行插入排序，
从链表第一个元素开始可以视为部分已排序，每次操作从链表中移除一个元素，然后原地将移除的元素插入到已排好序的部分。
数据范围：链表长度满足 1 \le n \le 1000\1≤n≤1000  ，链表中每个元素满足 |val| \le 10000 \∣val∣≤10000
例如输入{2,4,1}时，对应的返回值为{1,2,4}，转换过程如下图所示：

示例1
输入：
{1,2,3}
返回值：
{1,2,3}


示例2
输入：
{2,4,1}
返回值：
{1,2,4}
*/

type TreeNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	// write code here
	/*

		ListNode* insertionSortList(ListNode* head) {
		        // write code here
		        if(!head || !head->next)
		            return head;
		        typedef ListNode* list;

		        list node = (list)malloc(sizeof(ListNode));
		        node->next = head;
		        node->val = -99999;

		        list phead = head;
		        list p = node;//给链表增加一个头节点，以便于插入操作
		        list loc = nullptr;

		        while(phead){
		            p = node;
		            loc = phead->next;//当前节点的下一个节点，也是可能的排序元素
		            if(loc){
		                if(phead->val > loc->val){
		                    //若当前元素大于后面一个元素，则将后一个元素插入到前面的有序序列中
		                    list loc_min = nullptr;//插入点的前一个节点
		                    while(p->val < loc->val){
		                        loc_min = p;
		                        p = p->next;
		                    }
		                    phead->next = loc->next;
		                    loc->next = loc_min->next;
		                    loc_min->next = loc;
		                    continue;
		                }
		            }
		            else//后一个节点为空，则退出循环
		                break;
		            phead = phead->next;
		        }
		        return node->next;//应该是返回队头节点，而不是头节点
		    }

	*/

	/*
		插入排序思想,当前值 <= val 时继续迭代,直到遇到较大值,由于链表结构,需要重置 next 节点,若是数组要进行移位操作
	*/

	//if head == nil {
	//	return nil
	//}
	//dummy := &ListNode{-10001, head}
	//tail := head //已排序链表的尾节点
	//pre := head
	//head = head.Next
	//for head != nil {
	//	if head.Val >= tail.Val {
	//		//直接挂在已排序链表的尾节点后面即可
	//		tail.Next = head
	//		tail = tail.Next
	//	} else {
	//		//从原链表中将head断开
	//		next := head.Next
	//		pre.Next = next
	//
	//		//从dummy.next后面开始找到第一个节点值大于head的值的节点，然后将head插在它的前面
	//		prev := dummy
	//		start := dummy.Next
	//		for start.Val <= head.Val {
	//			prev = start
	//			start = start.Next
	//		}
	//		prev.Next = head
	//		head.Next = start
	//		head = pre
	//	}
	//	pre = head
	//	head = head.Next
	//}
	//return dummy.Next

	//if head == nil || head.Next == nil {
	//	return head
	//}
	//dummy := &ListNode{Val: -1}
	//cur := head
	//for cur != nil {
	//	pre := dummy
	//	for pre.Next != nil && pre.Next.Val < cur.Val {
	//		pre = pre.Next
	//	}
	//	temp := cur.Next
	//	cur.Next = pre.Next
	//	pre.Next = cur
	//	cur = temp
	//}
	//return dummy.Next

	//dumny := &ListNode{Val: -10001}
	//p := dumny
	//for head != nil {
	//	for p.Next != nil && p.Next.Val < head.Val {
	//		p = p.Next
	//	}
	//
	//	node := p.Next
	//	p.Next = head
	//	head = head.Next
	//
	//	p.Next.Next = node
	//
	//	p = dumny
	//}
	//
	//return dumny.Next

	//dummy := new(ListNode)
	//for head != nil {
	//	cur := dummy
	//	for cur.Next != nil && cur.Next.Val < head.Val {
	//		cur = cur.Next
	//	}
	//
	//	cur.Next, head.Next, head = head, cur.Next, head.Next
	//}
	//
	//return dummy.Next

	//cur := head
	//for cur.Next != nil {
	//	if cur.Val <= cur.Next.Val {
	//		cur = cur.Next
	//	} else {
	//		// 2.4,1
	//		tmp := cur.Next
	//		cur.Next = cur.Next.Next
	//		if tmp.Val <= head.Val {
	//			tmp.Next = head
	//			head = tmp
	//		} else {
	//			// 2,6,5
	//			p := head
	//			for p.Next != nil {
	//				// 两头小，中间大
	//				if p.Val < tmp.Val && p.Next.Val >= tmp.Val {
	//					tmp.Next = p.Next
	//					p.Next = tmp
	//					break
	//				}
	//				// ???
	//				p = p.Next
	//			}
	//		}
	//	}
	//}
	//return head

	//var (
	//	pHead = &ListNode{}
	//)
	//for head != nil {
	//	p := pHead
	//	for p.Next != nil {
	//		if p.Next.Val < head.Val {
	//			p = p.Next
	//		} else {
	//			break
	//		}
	//	}
	//	tmp := head.Next
	//	head.Next = p.Next
	//	p.Next = head
	//	head = tmp
	//}
	//return pHead.Next

	dummy := &ListNode{Val: -10000}
	cur := head
	for cur != nil {
		pre := dummy
		for pre.Next != nil && pre.Next.Val < cur.Val {
			pre = pre.Next
		}
		tmp := pre.Next
		pre.Next = cur
		cur = cur.Next
		pre.Next.Next = tmp
	}
	return dummy.Next


}
