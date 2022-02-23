package week1

import (
	"fmt"
	"testing"
)

func TestDefangIpAddr(t *testing.T) {
	fmt.Println(defangIPaddr("1.1.1.1"))
	fmt.Println(defangIPaddr("255.100.50.0"))
}
