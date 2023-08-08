# golang-starter
고랭을 나시고랭

#### 모든 내용은 공식 문서를 참고했습니다.

https://go.dev/doc/


1. 고랭 초기 설정.

        $ go mod init example/hello

2. hello.go 파일을 만들고 다음 내용 복사 붙여넣기하기

        package main

        import "fmt"

        func main() {
            fmt.Println("Hello,         World!")
        }

3. 마법처럼 Hello, World! 출력하기

        $ go run .



 * This function takes a name parameter whose type is string. The function also returns a string. In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name

- 고랭에서는 콜론 안쓰고 바로 리턴 타입을 명시해.

![Alt text](https://go.dev/doc/tutorial/images/function-syntax.png)

4. golang은 c언어와 비슷하게 class, object가 없고, method나 constructor들은 모두 struct로 만든다.
- Go 언어는 객체지향 프로그래밍(Object Oriented Programming, OOP)을 고유의 방식으로 지원한다. 즉, Go에는 전통적인 OOP 언어가 가지는 클래스, 객체, 상속 개념이 없다. 전통적인 OOP의 클래스(class)는 Go 언어에서 Custom 타입을 정의하는 struct로 표현되는데, 전통적인 OOP의 클래스가 필드와 메서드를 함께 갖는 것과 달리 Go 언어의 struct는 필드만을 가지며, 메서드는 별도로 분리하여 정의한다


5. 메서드 (Methods)

go 언어는 객체지향 프로그래밍 (OOP)를 고유의 방식으로 지원해요.
타 언어의 OOP 클래스가 필드와 메서드를 함께 갖는 것과 달리 Go 언어에서는 struct가 필드만을 가지며, 메서드는 별도로 분리되어 저장돼요.
Go 메서드는 특별한 형태의 func 함수에요. 

메서드는 func 키워드와 함수명 사이에 "그 함수가 어떤 struct를 위한 메서드인지"를 표시하게 돼요. 
흔히 receiver로 불리우는 이 부분은 메서드가 속한 struct 타입과 struct 변수명을 지정하는데, 
struct 변수명은 함수 내에서 마치 입력 파라미터처럼 사용돼요. 
예를 들어, 아래 예제는 Rect라는 struct를 정의하고 area() 라는 메서드를 정의하고 있어요.
func와 area() 사이에 Rect 타입의 r 이 정의되고 이를 함수 본문에서 사용하고 있어요.
메서드가 선언된 이후에는 Rect 구조체의 객체는 rect.area() 문장처럼 area() 메소드를 struct 객체로부터 직접 호출할 수 있어요.

```go
package main
 
//Rect - struct 정의
type Rect struct {
    width, height int
}
 
//Rect의 area() 메소드
func (r Rect) area() int {
    return r.width * r.height   
}
 
func main() {
    rect := Rect{10, 20}
    area := rect.area() //메서드 호출
    println(area)
}
```

- pointer receiver 사용법

```go
// 포인터 Receiver
func (r *Rect) area2() int {
    r.width++
    return r.width * r.height
}
 
func main() {
    rect := Rect{10, 20}
    area := rect.area2() //메서드 호출
    println(rect.width, area) // 11 220 출력
}
```

6. Multiple Return Values
- 여러 개의 return value가 가능함.
- 파이썬과 유사하고 인덱스 접근도 비슷하게 가능

```go
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
```
