package week1

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 { // 如果用 0 块目前拼跳水板
		return []int{}
	}
	if shorter == longer { // 如果短板、长板长度相同
		return []int{shorter * k}
	}
	delta := longer - shorter
	rs := make([]int, 0, k+1)
	for i := 0; i <= k; i++ {
		rs[i] = shorter*k + i*delta
	}
	return rs
}
