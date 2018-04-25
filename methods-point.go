package main

import (
	"math"
	"fmt"
)

type VertextP struct {
	X, Y float64
}

func Abs(v VertextP) float64 {
	return math.Sqrt(v.Y*v.Y + v.X*v.Y)
}

//func (v VertextP) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}

func Scale(v *VertextP, f float64) {
	v.X = v.X * 10
	v.Y = v.Y * 10
}

func main() {

	v := VertextP{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
