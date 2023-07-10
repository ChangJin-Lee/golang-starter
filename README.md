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