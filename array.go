package main

import "fmt"

//数组
func main() {
	//var a [2]string
	//a[0] = "Hello"
	//a[1] = "World"
	//fmt.Println(a[0],a[1])
	//fmt.Println(a)
	//
	//primse := [6]int{2,3,5,7,11,13}
	//fmt.Println(primse)
	//
	//var s = primse[1:4]
	//fmt.Println(s)

	//names := [4]string{
	//	"John",
	//	"Paul",
	//	"George",
	//	"Ringo",
	//}
	//fmt.Println(names)
	//
	//a := names[0:2]
	//b := names[1:3]
	//fmt.Println(a,b)
	//
	//b[0] = "XXX"
	//fmt.Println(b)
	//fmt.Println(names)

	//q := []int{2, 3, 5, 7, 11, 13}
	//fmt.Println(q)
	//
	//r := []bool{true, false, true, true, false, true}
	//fmt.Println(r)
	//
	//s := []struct{
	//	i int
	//	b bool
	//}{
	//	{2,true},
	//	{3,false},
	//	{5,true},
	//	{7,true},
	//	{11,false},
	//	{13,true},
	//}
	//
	//fmt.Println(s)

	//切片左闭右开，默认左边界值为0，右边界值为数组长度
	//一下切片是等价的
	//a[0:10]
	//a[:10]
	//a[0:]
	//a[:]
	//s :=[]int{2,3,5,7,11,13}
	//
	//s = s[1:4]
	//fmt.Println(s)
	//
	//s = s[:2]
	//fmt.Println(s)
	//
	//s=s[1:]
	//fmt.Println(s)

	////切片拥有长度与容量
	//s := []int{2,3,5,7,11,13}
	//printSlice(s)
	//
	//s = s[:0]
	//printSlice(s)
	//
	//s = s[:4]
	//printSlice(s)
	//
	//s = s[2:]
	//printSlice(s)

	//a := make([]int,5)
	//printSlice("a",a)
	//
	//b := make([]int,0,5)
	//printSlice("b",b)
	//
	//c := b[:2]
	//printSlice("c",c)
	//
	//d := c[2:5]
	//printSlice("d",d)


	//遍历数组
	//pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	//
	//for i, v := range pow {
	//	fmt.Printf("2**%d = %d\n", i, v)
	//}
}

//func printSlice(s []int) {
//	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
//}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
