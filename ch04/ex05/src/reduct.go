package src

func Reduct(arr []string) []string {
	var result []string

	for i, a := range arr {
		if i+1 < len(arr) {
			if a != arr[i+1] {
				result = append(result, a)
			}
		} else {
			if result[len(result)-1] != a {
				result = append(result, a)
			}
		}
	}

	return result
}
