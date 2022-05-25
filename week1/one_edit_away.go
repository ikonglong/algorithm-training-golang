package week1

import (
	"math"
	"unicode/utf8"
)

// https://leetcode-cn.com/problems/one-away-lcci/
// 这个版本考虑中文
func oneEditAway(first string, second string) bool {
	// 如果两个字符串都为 ""
	if first == "" && second == "" {
		return true
	}

	runeLenDiff := int(math.Abs(float64(utf8.RuneCountInString(first) - utf8.RuneCountInString(second))))
	if runeLenDiff > 1 { // 如果字符串长度差值大于 1，则不可能一次编辑使它们相同
		return false
	} else if runeLenDiff == 0 { // 如果长度相同且只有一个字符不同，那么一次编辑就可使它们相同
		foundDiff := false
		for i, j := 0, 0; i < len(first) && j < len(second); {
			r1, s1 := utf8.DecodeRuneInString(first[i:])
			r2, s2 := utf8.DecodeRuneInString(second[j:])
			if r1 != r2 {
				if foundDiff {
					return false
				} else {
					foundDiff = true
				}
			}
			i += s1
			j += s2
		}
		return true
	} else { // 如果长度差为 1，则较短字符串在某个位置插入一个字符就可使它们相同
		if len(first) < len(second) {
			return oneMoreChar(&first, &second)
		} else {
			return oneMoreChar(&second, &first)
		}
	}
}

func oneMoreChar(shortPtr *string, longPtr *string) bool {
	short, long := *shortPtr, *longPtr
	diffFound := false
	for i, j := 0, 0; i < len(short) && j < len(long); {
		r1, s1 := utf8.DecodeRuneInString(short[i:])
		r2, s2 := utf8.DecodeRuneInString(long[j:])
		if r1 != r2 {
			if diffFound {
				return false
			} else {
				diffFound = true
			}
			j += s2
			r2, s2 = utf8.DecodeRuneInString(long[j:])
			if r1 != r2 {
				return false
			} else {
				i += s1
				j += s2
			}
		} else {
			i += s1
			j += s2
		}
	}
	return true
}

// 这个版本假设所有字符都是英文字符（ASCII）
func oneEditAwayV2(first string, second string) bool {
	// 为了避免传递字符串参数时发生拷贝，内部包装了以字符串指针为参数的同名方法
	return _oneEditAwayV2(&first, &second)
}

func _oneEditAwayV2(firstPtr *string, secondPtr *string) bool {
	first, second := *firstPtr, *secondPtr
	lenF, lenS := len(first), len(second)
	if lenF > lenS {
		// 很巧妙，通过一次递归调用，确保第一个字符串长度总比第二个短
		return _oneEditAwayV2(secondPtr, firstPtr)
	}
	lenDiff := lenS - lenF
	if lenDiff > 1 {
		return false
	} else if lenDiff == 0 {
		diffCnt := 0
		for i := 0; i < lenF; i++ {
			if first[i] != second[i] {
				diffCnt++
				if diffCnt > 1 {
					return false
				}
			}
		}
		return true
	} else {
		// offset 的思路很巧妙呀
		for i, ofs := 0, 0; i < lenF; {
			if first[i] == second[i+ofs] {
				i++
			} else {
				ofs++
				if ofs > 1 {
					return false
				}
			}
		}
		return true
	}
}
