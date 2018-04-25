package main

import "fmt"

//指针
func main() {

	i,j := 42,2701
	p := &i
	fmt.Println(*p)
	*p = 21

	fmt.Println(i)
	p = &j
	*p = *p /37
	fmt.Println(j)// see the new value of j

}
