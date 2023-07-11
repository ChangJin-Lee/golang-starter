

package main

import "fmt"


func canIDrink(age int) bool {
	switch koreaAge := age + 2; koreaAge {
	case 10:
		return false
	case 20:
		return true
	}

	return false
}

func main() {
	fmt.Println(canIDrink(18))
}