package ex10

import (
	"testing"
)

func isTopological(tree map[string][]string, result []string) bool {

	nodes := make(map[string]int)

	for i, course := range result {
		nodes[course] = i
	}

	for parent, v := range tree {
		for _, child := range v {
			// 親のが子よりも順序が先に来てたらNG
			if nodes[parent] < nodes[child] {
				return false
			}
		}
	}
	return true

}

func TestSort(t *testing.T) {
	tests := []struct {
		data map[string][]string
	}{
		{
			map[string][]string{
				"data structures": {"discrete math"},
				"discrete math":   {"intro to programming"},
				"algorithms":      {"data structures"},
			},
		},
		{
			map[string][]string{
				"algorithms":      {"data structures"},
				"discrete math":   {"intro to programming"},
				"data structures": {"discrete math"},
			},
		},
		{
			map[string][]string{
				"algorithms": {"data structures"},
				"calculus":   {"linear algebra"},

				"compilers": {
					"data structures",
					"formal languages",
					"computer organization",
				},

				"data structures":       {"discrete math"},
				"databases":             {"data structures"},
				"discrete math":         {"intro to programming"},
				"formal languages":      {"discrete math"},
				"networks":              {"operating systems"},
				"operating systems":     {"data structures", "computer organization"},
				"programming languages": {"data structures", "computer organization"},
			},
		},
	}

	for _, test := range tests {
		result := Sort(test.data)
		if !isTopological(test.data, result) {
			t.Errorf("result is not topological %v", result)
		}
	}
}
