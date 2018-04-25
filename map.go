package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m0 = map[string]Vertex{
	"Bell Labs": {40.3333, -70.283283},
	"Google":    {37.4783, -122.08349384},
}

func main() {
	fmt.Println(m0)

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println(m["Answer"])

	m["Answer"] = 48
	fmt.Println(m["Answer"])

	delete(m, "Answer")
	fmt.Println(m["Answer"])


	//通过双赋值检测某个键是否存在：
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
