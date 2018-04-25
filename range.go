package main

import "fmt"

func main() {

	//使用range去求一个slice的和。使用数组和这个类似
	nums := []int{2, 3, 4}

	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	//数组上使用range将传入index 和值两个变量。上面那个例子我们不需要使用该元素的序号，所以使用"_"省略。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range 也可以用在map 的键值对上
	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//range也可以用来枚举Unicode字符串，第一个参数是字符的索引，第二个字符 （Unicode的值）本身
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
