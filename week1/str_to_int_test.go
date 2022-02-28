package week1

import (
	"fmt"
	"testing"
)

func TestStrToInt(t *testing.T) {
	//fmt.Println(strToInt("  +42  "))
	//fmt.Println(strToInt("  -42 abc"))
	//fmt.Println(strToInt("  42 12"))
	//fmt.Println(strToInt(" 0000  "))
	//fmt.Println(strToInt(" 0005  "))
	//fmt.Println(strToInt("-91283472332"))
	fmt.Println(strToInt("9223372036854775808"))

	//fmt.Println(strToInt(""))
	//fmt.Println(strToInt("    "))
	//fmt.Println(strToInt("  000 abc"))
	//fmt.Println(strToInt("  123 4"))
	//fmt.Println(strToInt("  0005   "))
	//fmt.Println(strToInt("  hell042"))
	//fmt.Println(strToInt("  42 hello"))
	//fmt.Println(strToInt("  42 hello"))
}
