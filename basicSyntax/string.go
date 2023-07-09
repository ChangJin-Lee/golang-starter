package main

import "fmt"

var someName = "hello"

func main() {
	// strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	// default value is ""
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	//  Short Assignment Statement
	// 선언과 동시에 초기화 가능해.
	// function 내에서만 사용 가능해. 함수 밖에서는 꼭 var 써야해.
	namefour := "yoshi"

	fmt.Println(nameOne, nameTwo, nameThree, namefour)

}
