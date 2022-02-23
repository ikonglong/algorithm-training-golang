package week1

func reverseString(s []byte) {
	if s == nil || len(s) == 0 {
		return
	}
	// 让两个指针 l, w 分别指向数组最左端、最右端的元素
	// 交换 l, w 指向的元素，直到 l, w 指向同一个元素或者 l > w
	for l, w := 0, len(s)-1; l < w; l, w = l+1, w-1 {
		//swap(s, l, w)
		s[l], s[w] = s[w], s[l]
	}
}

func swap(s []byte, i int, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}
