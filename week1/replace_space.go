package week1

func replaceSpace(s string) string {
	runes := make([]rune, 0, len(s) + len(s)/5)
	for _, r := range s {
		if r == ' ' {
			runes = append(runes, '%', '2', '0')
		} else {
			runes = append(runes, r)
		}
	}
	return string(runes)
}
