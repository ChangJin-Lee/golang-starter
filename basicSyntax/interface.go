package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type conractor struct {
	name			string
	hourlyPay		int
	hoursPerYear	int
}

func (c conractor) getName() string {
	return c.name
}

func (c conractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name	string
	salary	int
}

func (ft fullTime) getSalary() int{
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func test(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("================================================================")
}

func main() {
	test(fullTime{
		name:	"Jack",
		salary: 50000,
	})
	test(conractor{
		name:			"Bob",
		hourlyPay:		100,
		hoursPerYear:	73,
	})
	test(conractor{
		name:			"Jill",
		hourlyPay:		872,
		hoursPerYear:	982,
	})
}
