package series

func All(n int, s string) []string {
	out := []string{}
	for i := 0; i <= len(s)-n; i++ {
		out = append(out, s[i:i+n])
	}
	return out
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
