package validparentheses

import "strings"

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Top(str string) bool {
	if s.IsEmpty() {
		return false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		if strings.Compare(str, element) == 0 {
			return true
		} else {
			return false
		}
	}

}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

func IsValid(s string) bool {
	split_string := strings.Split(s, "")
	var stack Stack

	for _, e := range split_string {
		switch e {
		case "(":
			{
				stack.Push(")")
			}
		case "{":
			{
				stack.Push("}")
			}
		case "[":
			{
				stack.Push("]")
			}
		default:
			{
				if stack.Top(e) {
					stack.Pop()
				} else {
					return false
				}

			}
		}
	}

	return stack.IsEmpty()
}
