package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    // 고랭에서는 := 이게 변수를 빠르게 초기화하는 방법이야.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
