package week1

import "testing"

// "" 非法
// nil 非法
// "123.92.2.34" 合法
// " 123 .92 . 2 . 34   " 合法
// "123.9  2.2.34" 非法
// "123.92,2.34" 非法
// "12b.92.2.34" 非法
// "-1.25.2.34"  非法
// "257.1.2.34"  非法
// "123.92.2.34  x  " 非法
// " 1  123.92.2.34 " 非法

func TestCheckIp(t *testing.T) {
	if checkIp("") == true {
		t.Errorf(`checkIp(""): expected: false, actual: true`)
	}
	if checkIp("123.92.2.34") == false {
		t.Errorf(`checkIp("123.92.2.34"): expected: true, actual: false`)
	}
	if checkIp(" 123 .92 . 2 . 34   ") == false {
		t.Errorf(`checkIp(" 123 .92 . 2 . 34   "): expected: true, actual: false`)
	}
	if checkIp("123.9  2.2.34") == true {
		t.Errorf(`checkIp("123.9  2.2.34"): expected: false, actual: true`)
	}
	if checkIp("123.92,2.34") == true {
		t.Errorf(`checkIp("123.92,2.34"): expected: false, actual: true`)
	}
	if checkIp("12b.92.2.34") == true {
		t.Errorf(`checkIp("12b.92.2.34"): expected: false, actual: true`)
	}
	if checkIp("-1.25.2.34") == true {
		t.Errorf(`checkIp("-1.25.2.34"): expected: false, actual: true`)
	}
	if checkIp("257.1.2.34") == true {
		t.Errorf(`checkIp("257.1.2.34"): expected: false, actual: true`)
	}
	if checkIp("123.92.2.34  x  ") == true {
		t.Errorf(`checkIp("123.92.2.34  x  "): expected: false, actual: true`)
	}
	if checkIp(" 1  123.92.2.34 ") == true {
		t.Errorf(`checkIp(" 1  123.92.2.34 "): expected: false, actual: true`)
	}
}