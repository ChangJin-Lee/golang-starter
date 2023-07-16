// go 언어는 객체지향 프로그래밍 (OOP)를 고유의 방식으로 지원해요.
// 타 언어의 OOP 클래스가 필드와 메서드를 함께 갖는 것과 달리 Go 언어에서는 struct가 필드만을 가지며, 메서드는 별도로 분리되어 저장돼요.
// Go 메서드는 특별한 형태의 func 함수에요. 

// 메서드는 func 키워드와 함수명 사이에 "그 함수가 어떤 struct를 위한 메서드인지"를 표시하게 돼요. 
// 흔히 receiver로 불리우는 이 부분은 메서드가 속한 struct 타입과 struct 변수명을 지정하는데, 
// struct 변수명은 함수 내에서 마치 입력 파라미터처럼 사용돼요. 
// 예를 들어, 아래 예제는 Rect라는 struct를 정의하고 area() 라는 메서드를 정의하고 있어요.
// func와 area() 사이에 Rect 타입의 r 이 정의되고 이를 함수 본문에서 사용하고 있어요.
// 메서드가 선언된 이후에는 Rect 구조체의 객체는 rect.area() 문장처럼 area() 메소드를 struct 객체로부터 직접 호출할 수 있어요.

package main

import "fmt"

//Rect - struct definition
type Rect struct {
	width, height int
}

//Rect area() method
func (r Rect) area() int {
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area() // call method
	fmt.Println(area)
}