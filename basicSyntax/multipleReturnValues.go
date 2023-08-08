package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	// 배열을 받아오는거라 스트링 char들을 접근함.
	s := strings.ToUpper(n)
	// " "를 기준으로 Split함. ["I", "AM", "CJ"]
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	// initials의 사이즈를 확인함
	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fn1, sn1 := getInitials("I am CJ!")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("backend possible? yes")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("let's go coding")
	fmt.Println(fn3, sn3)
}
