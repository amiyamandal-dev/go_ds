package graphds

import "testing"

func TestIsValidBST(t *testing.T) {
	input_arr := []interface{}{2, 1, 3}
	tree := CreateGraph(input_arr)
	result := IsValidBST(tree)
	if result != true {
		t.Error("result dont match")
	}
	input_arr = []interface{}{5, 1, 4, nil, nil, 3, 6}
	tree = CreateGraph(input_arr)
	result = IsValidBST(tree)
	if result != false {
		t.Error("result dont match")
	}
}
