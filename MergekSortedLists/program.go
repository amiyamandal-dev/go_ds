package mergeksortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}


func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else {
		new_node := new(ListNode)
		smallest := 99999
		main_index := -1
		for index, i := range lists {
			if i != nil {
				if smallest > i.Val {
					smallest = i.Val
					main_index = index
				}
			}
		}
		if main_index == -1 {
			return nil
		}
		new_node.Val = smallest
		lists[main_index] = lists[main_index].Next
		new_node.Next = MergeKLists(lists)
		return new_node
	}
}
