package mergetwosortedlists

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := new(ListNode)
	l2 := new(ListNode)
	v1 := []int{1, 2, 3, 4, 5}
	for _, i := range v1 {
		l1.Push(i)
	}
	v2 := []int{1, 2, 3, 4, 5}
	for _, i := range v2 {
		l2.Push(i)
	}
	rez := MergeTwoLists(new(ListNode), new(ListNode))
	for rez != nil {
		fmt.Println(rez.Val)
		rez = rez.Next
	}
}
