package week3

type SortedStack struct {
	s []int
}

func NewSortedStack() SortedStack {
	return SortedStack{
		s: make([]int, 0, 8),
	}
}

func (ss *SortedStack) Push(val int) {
	ss.s = append(ss.s, val)
	ss.s = sort(ss.s)
}

func (ss *SortedStack) Pop() {
	ss.s = ss.s[:len(ss.s)-1]
}

func (ss *SortedStack) Peek() int {
	if ss.IsEmpty() {
		return -1
	}
	return ss.s[len(ss.s)-1]
}

func (ss *SortedStack) IsEmpty() bool {
	return len(ss.s) == 0
}

func sort(s []int) []int {
	minTop := make([]int, 0, len(s))
	for len(s) != 0 {
		// pop from s
		tmp := s[len(s)-1]
		s = s[:len(s)-1]
		for len(minTop) != 0 && minTop[len(minTop)-1] > tmp { // peek minTop
			// pop from minTop
			top := minTop[len(minTop)-1]
			minTop = minTop[:len(minTop)-1]
			// push into s
			s = append(s, top)
		}
		// push tmp into minTop
		minTop = append(minTop, tmp)
	}
	return minTop
}

/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */
