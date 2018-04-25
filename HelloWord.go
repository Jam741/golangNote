package main

import (
	"fmt"
	"math/rand"
	"math"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, java, python bool

var i, j int = 1, 2

const Pi = 3.14

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int)int  {
	return x*10 +1
}

func needFloat(x float64)float64  {
	return x*0.1
}

func main() {
	fmt.Println("hello word")
	fmt.Println("My favorite number is ", rand.Intn(10))
	fmt.Printf("New you have %g problems", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(4, 2))
	fmt.Println(swap("a", "b"))
	fmt.Println(split(18))
	fmt.Println(c, java, python)
	fmt.Println(i, j)

	ngla := "尼古拉斯"
	fmt.Println(ngla)

	var x , y int = 3,4
	f := float64(math.Sqrt(float64(x*x + y*y)))
	z := uint(f)
	fmt.Println(f,z)

	v := 24.0
	fmt.Printf("v is of type %T\n",v)
	fmt.Println(Pi)


	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}
