package strings

func Join(sep string, strs ...string) string {
	result := ""
	for i, str := range strs {
		if i < len(strs)-1 {
			result += (str + sep)
		} else {
			result += str
		}
	}
	return result
}
