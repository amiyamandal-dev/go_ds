package graphds

import (
	"fmt"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	graph := CreateGraph([]interface{}{5, 1, 4, nil, nil, 3, 6})
	fmt.Println(graph)
}
