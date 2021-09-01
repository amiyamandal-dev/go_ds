package mergetwosortedlists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Push(Data int) {
	new_node := new(ListNode)
	new_node.Val = Data
	new_node.Next = nil
	for n.Next != nil {
		n = n.Next
	}
	n.Next = new_node
}

func MergeTwoListsOld(l1 *ListNode, l2 *ListNode) *ListNode {
	sorted_list := new(ListNode)

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			sorted_list.Push(l2.Val)
			l2 = l2.Next
		} else {
			sorted_list.Push(l1.Val)
			l1 = l1.Next
		}
	}
	for l1 != nil {
		sorted_list.Push(l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		sorted_list.Push(l2.Val)
		l2 = l2.Next
	}
	return sorted_list
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		if l2 == nil {
			return nil
		} else {
			return l2
		}
	}
	if l2 == nil {
		if l1 == nil {
			return nil
		} else {
			return l1
		}
	} else {

		new_node := new(ListNode)
		if l1.Val > l2.Val {
			new_node.Val = l2.Val
			new_node.Next = MergeTwoLists(l1, l2.Next)
		} else {
			new_node.Val = l1.Val
			new_node.Next = MergeTwoLists(l1.Next, l2)
		}
		return new_node
	}
}
