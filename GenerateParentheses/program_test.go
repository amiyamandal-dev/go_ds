package generateparentheses

import (
	"fmt"
	"testing"
)

func TestGerateParenthesis(t *testing.T) {
	rez := GenerateParenthesis(3)
	fmt.Println(rez)
}
