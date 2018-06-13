package ex11

import (
	"fmt"
	"strings"
	"testing"
)

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
		result, err := Sort(test.data)
		if err != nil {
			t.Errorf("error is thrown %s", err)
		}
		ok, _ := isTopological(test.data, result)
		if !ok {
			t.Errorf("result is not topological %v", result)
		}
	}
}

func TestFailedSort(t *testing.T) {
	tests := []struct {
		data map[string][]string
	}{
		{
			map[string][]string{
				"a": {"b", "c"},
				// cyclic data
				"b": {"a"},
			},
		},
		{
			map[string][]string{
				"algorithms": {"data structures"},
				"calculus":   {"linear algebra"},
				// cyclic data
				"linear algebra": {"calculus"},

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
		items, err := Sort(test.data)
		if len(items) != 0 {
			t.Error("need to throw error")
		}
		if err == nil {
			t.Error("need to throw error")
		}
		if !strings.Contains(fmt.Sprintf("%s", err), "cyclic data is found") {
			t.Error("does not expect error")
		}
	}
}
