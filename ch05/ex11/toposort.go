package ex11

import "fmt"

func isTopological(tree map[string][]string, result []string) (bool, error) {

	nodes := make(map[string]int)

	for i, course := range result {
		nodes[course] = i
	}

	for parent, v := range tree {
		for _, child := range v {
			// 親のが子よりも順序が先に来てたらNG
			if nodes[parent] < nodes[child] {
				return false, fmt.Errorf("cyclic data is found, %s", child)
			}
		}
	}
	return true, nil

}

func Sort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	for key := range m {
		visitAll([]string{key})
	}
	ok, err := isTopological(m, order)
	if !ok {
		return nil, err
	}
	return order, nil
}
