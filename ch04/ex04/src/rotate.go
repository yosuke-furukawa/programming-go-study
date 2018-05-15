package src

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}

	return a
}

func Rotate(arr []int, diff int) []int {
	for i := 0; i < gcd(diff, len(arr)); i++ {
		for prev, next := i, i+diff; ; prev, next = next, prev+diff {
			if next >= len(arr) {
				next = next - len(arr)
			}
			if next == i {
				break
			}
			arr[next], arr[prev] = arr[prev], arr[next]
		}
	}
	return arr
}
