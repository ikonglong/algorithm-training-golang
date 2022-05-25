package week1

// https://leetcode-cn.com/problems/master-mind-lcci/
func masterMind(solution string, guess string) []int {
	// 先找出猜中的槽位。对于剩余槽位，再看颜色是否猜中（即伪猜中）
	guessed := 0
	idxNotGuessed := make([]int, 0, 4)
	notGuessed := make(map[rune]int)
	for i := 0; i < len(solution); i++ {
		color := solution[i]
		if color == guess[i] {
			guessed++
		} else {
			idxNotGuessed = append(idxNotGuessed, i)
			notGuessed[rune(color)] = notGuessed[rune(color)] + 1
		}
	}

	pseudoGuessed := 0
	for _, idx := range idxNotGuessed {
		color := rune(guess[idx])
		if cnt, found := notGuessed[color]; found {
			pseudoGuessed++
			cnt--
			if cnt == 0 {
				delete(notGuessed, color)
			} else {
				notGuessed[color] = cnt
			}
		}
	}

	return []int{guessed, pseudoGuessed}
}

func masterMindV2(solution string, guess string) []int {
	cntMap := make(map[byte]int)
	for i := 0; i < len(solution); i++ {
		cntMap[solution[i]]++
	}

	cnt1, cnt2 := 0, 0
	for i := 0; i < len(guess); i++ {
		if solution[i] == guess[i] {
			cnt1++
		}
		if v, found := cntMap[guess[i]]; found && v > 0 {
			cnt2++
			cntMap[guess[i]]--
		}
	}

	return []int{cnt1, cnt2 - cnt1}
}
