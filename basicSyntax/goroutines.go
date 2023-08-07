// goroutine은 쓰레드를 구현하는 아주 가벼운 방법입니다.
// 두개의 function call은 비동기적으로 실행되빈다. 둘 중 하나가 끝날 때까지 기다립니다.
// 더 다양하게 만들어 보려면 WaitGroup를 사용하면 됩니다.
package main

import (
	"fmt"
	"time"
)

func f(from string){
	for i := 0; i<3 ; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}