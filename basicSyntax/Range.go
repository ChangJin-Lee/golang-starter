// range는 다양한 데이터의 iterators입니다. 
package main

import "fmt"

func main() {

	nums := []int{2,3,4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i,num := range nums {
		if num == 3{
			fmt.Println("index", i)
		}
	}

	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k,v := range kvs {
		fmt.Println("%s -> %s", k, v)
	}

	for k:= range kvs {
		fmt.Println("key", k)
	}

	for i,c := range "go"{
		fmt.Println(i,c)
	}
}
