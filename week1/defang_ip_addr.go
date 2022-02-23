package week1

// https://leetcode-cn.com/problems/defanging-an-ip-address/
func defangIPaddr(address string) string {
	bytes := make([]byte, len(address)+3*2, len(address)+3*2)
	for i, j := 0, 0; i < len(address); i++ {
		if rune(address[i]) == '.' {
			bytes[j] = '['
			j++
			bytes[j] = '.'
			j++
			bytes[j] = ']'
			j++
		} else {
			bytes[j] = address[i]
			j++
		}
	}
	return string(bytes)
}
