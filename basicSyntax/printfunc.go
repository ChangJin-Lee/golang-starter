package main

import "fmt"

func main() {

	age := 27
	name := "CJ"

	//print
	fmt.Print("Hello")
	fmt.Print("World!")
	fmt.Print("This is new line")

	// println
	fmt.Println("Hello, ninjas!")
	fmt.Println("goodbye ninjas!")
	fmt.Println("my age is", age, "and my name is", name)

	// printf (formatted strings)
	// %_ = 이건 format specifier 라고 해.
	// variable의 약자로 v를 사용해.
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f points! \n", 225.55)
	fmt.Printf("you scored %0.1f points! \n", 225.55)

	//Sprintf (save formatted strings)
	// print를 저장해 놓고 사용할 수 있어.
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is : ", str)
}
