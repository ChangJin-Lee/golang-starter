

package main

import "fmt"


type person struct {
	name     string
	age      int
	favFoood []string
}

func main() {
	favFoood := []string{"rice", "ramen"}
	// hwany := person{"hwany", 100, favFoood}
	hwany := person{name: "hwany", age: 100, favFoood: favFoood}
	fmt.Println(hwany)
}