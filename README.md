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