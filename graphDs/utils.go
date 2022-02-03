package graphds

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func NewGraph(Val interface{}) *TreeNode {
	return &TreeNode{
		Val:   Val,
		Left:  nil,
		Right: nil,
	}
}

func generate_util(Arr []interface{}, root *TreeNode, index_root int) {
	left_child := (2 * index_root) + 1
	right_child := (2 * index_root) + 2
	if len(Arr) > left_child {
		if Arr[left_child] != nil {
			temp_node := NewGraph(Arr[left_child])
			root.Left = temp_node
			fmt.Printf("%d -> %d\n", index_root, left_child)
			generate_util(Arr, temp_node, left_child)
		}

	}
	if len(Arr) > right_child {
		if Arr[right_child] != nil {
			temp_node := NewGraph(Arr[right_child])
			root.Right = temp_node
			fmt.Printf("%d -> %d\n", index_root, right_child)
			generate_util(Arr, temp_node, right_child)
		}
	}

}

func (t *TreeNode) TreeGrenerate(Arr []interface{}) {
	generate_util(Arr, t, 0)
}

func CreateGraph(Arr []interface{}) *TreeNode {
	if len(Arr) > 0 {
		t := NewGraph(Arr[0])
		t.TreeGrenerate(Arr)
		return t
	}
	return nil
}
