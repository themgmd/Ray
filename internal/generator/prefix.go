package generator

func prefix(s string) string {
	return string([]rune(s)[:1])
}
